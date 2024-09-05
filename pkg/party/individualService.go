package party

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"tmfEcho/internal/api/apihelper"
	"tmfEcho/internal/database"
	"tmfEcho/internal/log"
	"tmfEcho/internal/util"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

type IndividualData struct {
	Id                       string                         `json:"id,omitempty" db:"cust_numb"`
	Type                     string                         `json:"@type"  db:""`
	BaseType                 string                         `json:"@baseType"  db:""`
	GivenName                string                         `json:"givenName,omitempty"  db:"frst_name"  validate:"required"`
	FamilyName               string                         `json:"familyName,omitempty"  db:"last_name"`
	Name                     string                         `json:"name,omitempty"  db:""`
	IndividualIdentification []individualIdentificationData `json:"individualIdentification,omitempty" db:""  validate:"required"`
}

type individualIdentificationData struct {
	IdentificationId   string `json:"identificationId,omitempty" db:"id_type"`
	IdentificationType string `json:"identificiationType,omitempty" db:"id_numb" validate:"required"`
}

func DeleteIndividualService(s *PartyHandler, c echo.Context, id string, lt log.LogTracing) error {
	rowCnt := 0
	sqlStmt := "select count(*) cnt from cs_cust where cust_numb = " + database.DB_CONST_TERM_VAR_PREFIX + "1"

	sqlErr := s.DB.QueryRowx(sqlStmt, id).Scan(&rowCnt)
	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			lg := log.GenErrLog("SQL:"+sqlStmt, lt, log.E100017, sqlErr)
			log.AppTraceLog.Error(lg)
			omErr := util.NewOMError(lg)
			return c.JSON(http.StatusNotFound, omErr.ErrorReponsTMFJSON())
		}
		lg := log.GenErrLog("SQL:"+sqlStmt, lt, log.E000000, sqlErr)
		log.AppTraceLog.Error(lg)
		omErr := util.NewOMError(lg)
		return c.JSON(http.StatusInternalServerError, omErr.ErrorReponsTMFJSON())
	}
	if rowCnt <= 0 {
		return c.NoContent(http.StatusNotFound)
	}

	ctx := c.Request().Context()
	sqlStmt = "delete from cs_cust where cust_numb = " + database.DB_CONST_TERM_VAR_PREFIX + "1"
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		lg := log.GenErrLog("Begin DB transaction: ", lt, log.E000000, err)
		log.AppTraceLog.Error(lg)
		omErr := util.NewOMError(lg)
		return c.JSON(http.StatusInternalServerError, omErr.ErrorReponsTMFJSON())
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()
	stmt1, err := tx.Prepare(sqlStmt)
	if err != nil {
		lg := log.GenErrLog("SQL : "+sqlStmt, lt, log.E000000, err)
		log.AppTraceLog.Error(lg)
		omErr := util.NewOMError(lg)
		return c.JSON(http.StatusInternalServerError, omErr.ErrorReponsTMFJSON())
	}
	defer stmt1.Close()

	log.AppTraceLog.Debug(log.GenAppLog("Execute SQL: "+sqlStmt, lt))
	_, err = stmt1.Exec(id)
	if err != nil {
		lg := log.GenErrLog("SQL : "+sqlStmt, lt, log.E000000, err)
		log.AppTraceLog.Error(lg)
		omErr := util.NewOMError(lg)
		return c.JSON(http.StatusInternalServerError, omErr.ErrorReponsTMFJSON())
	}
	if err = tx.Commit(); err != nil {
		lg := log.GenErrLog("Commit transaction", lt, log.E000000, err)
		log.AppTraceLog.Error(lg)
		omErr := util.NewOMError(lg)
		return c.JSON(http.StatusInternalServerError, omErr.ErrorReponsTMFJSON())
	}

	return c.NoContent(http.StatusNoContent)

}

func UpdateIndividualService(s *PartyHandler, c echo.Context, id string, lt log.LogTracing) error {
	var data IndividualData
	requestMap := make(map[string]interface{})

	if bErr := c.Bind(&requestMap); bErr != nil {
		lg := log.GenErrLog("Wrong Request payload (Binding MAP)", lt, log.E201434, bErr)
		log.AppTraceLog.Error(lg)
		omErr := util.NewOMError(lg)
		return c.JSON(http.StatusBadRequest, omErr.ErrorReponsTMFJSON())
	}
	log.AppTraceLog.Debug(log.GenAppLog(fmt.Sprintf("CONTEXT MAP: %v\n", requestMap), lt))
	if !util.IsNotEmptyString(id) {
		lg := log.GenErrLog("ID is empty", lt, log.E206247, nil)
		log.AppTraceLog.Error(lg)
		omErr := util.NewOMError(lg)
		return c.JSON(http.StatusBadRequest, omErr.ErrorReponsTMFJSON())
	}
	sqlUpdate, OMerr := apihelper.JSONconverToUpdateValue(requestMap, &data, lt)
	if OMerr.Err != nil {
		lg := log.GenErrLog("Wrong Request payload (Binding SQL Update)", lt, log.E201434, OMerr.Err)
		log.AppTraceLog.Error(lg)
		omErr := util.NewOMError(lg)
		return c.JSON(http.StatusBadRequest, omErr.ErrorReponsTMFJSON())
	}
	sqlStmt := "update CS_CUST "
	if len(sqlUpdate) > 0 {
		sqlStmt += " SET " + strings.Join(sqlUpdate, " , ")
	}
	sqlStmt += " WHERE cust_numb = " + id
	data.Id = id
	log.AppTraceLog.Debug(log.GenAppLog("SQL STMT:"+sqlStmt, lt))

	// BEGIN TRANSACTIONS
	ctx := c.Request().Context()
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		lg := log.GenErrLog("Begin DB transaction: ", lt, log.E000000, err)
		log.AppTraceLog.Error(lg)
		omErr := util.NewOMError(lg)
		return c.JSON(http.StatusInternalServerError, omErr.ErrorReponsTMFJSON())
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()
	stmt1, err := tx.Prepare(sqlStmt)
	if err != nil {
		lg := log.GenErrLog("SQL : "+sqlStmt, lt, log.E000000, err)
		log.AppTraceLog.Error(lg)
		omErr := util.NewOMError(lg)
		return c.JSON(http.StatusInternalServerError, omErr.ErrorReponsTMFJSON())
	}
	defer stmt1.Close()

	log.AppTraceLog.Debug(log.GenAppLog("Execute SQL: "+sqlStmt, lt))
	_, err = stmt1.Exec()
	if err != nil {
		lg := log.GenErrLog("SQL : "+sqlStmt, lt, log.E000000, err)
		log.AppTraceLog.Error(lg)
		omErr := util.NewOMError(lg)
		return c.JSON(http.StatusInternalServerError, omErr.ErrorReponsTMFJSON())
	}
	if err = tx.Commit(); err != nil {
		lg := log.GenErrLog("Commit transaction", lt, log.E000000, err)
		log.AppTraceLog.Error(lg)
		omErr := util.NewOMError(lg)
		return c.JSON(http.StatusInternalServerError, omErr.ErrorReponsTMFJSON())
	}

	return c.JSON(http.StatusOK, data)
}

func SaveIndividualService(s *PartyHandler, c echo.Context, lt log.LogTracing) error {
	var data IndividualData
	if err := c.Bind(&data); err != nil {
		lg := log.GenErrLog("Wrong Request payload", lt, log.E201434, err)
		log.AppTraceLog.Error(lg)
		omErr := util.NewOMError(lg)
		return c.JSON(http.StatusBadRequest, omErr.ErrorReponsTMFJSON())
	}
	//validate required fields
	validate := validator.New()
	vErr := validate.Struct(data)
	if vErr != nil {
		lg := log.GenErrLog("Required fields", lt, log.E100009, vErr)
		log.AppTraceLog.Error(lg)
		omErr := util.NewOMError(lg)
		return c.JSON(http.StatusInternalServerError, omErr.ErrorReponsTMFJSON())
	}
	if len(data.IndividualIdentification) != 1 {
		lg := log.GenErrLog("Array of Indentification shall be 1", lt, log.E100009, vErr)
		log.AppTraceLog.Error(lg)
		omErr := util.NewOMError(lg)
		return c.JSON(http.StatusInternalServerError, omErr.ErrorReponsTMFJSON())
	}
	vErr = validate.Struct(data.IndividualIdentification[0])
	if vErr != nil {
		lg := log.GenErrLog("Required fields", lt, log.E100009, vErr)
		log.AppTraceLog.Error(lg)
		omErr := util.NewOMError(lg)
		return c.JSON(http.StatusInternalServerError, omErr.ErrorReponsTMFJSON())
	}

	sqlStmt := "select max(cust_numb) from cs_cust"
	var custNumb string
	err := s.DB.QueryRow(sqlStmt).Scan(&custNumb)
	if err != nil {
		if err == sql.ErrNoRows {
			custNumb = "100000"
		} else {
			lg := log.GenErrLog("SQL:"+sqlStmt, lt, log.E000000, err)
			log.AppTraceLog.Error(lg)
			omErr := util.NewOMError(lg)
			return c.JSON(http.StatusInternalServerError, omErr.ErrorReponsTMFJSON())
		}
	}
	custNumbInt, _ := strconv.Atoi(custNumb)
	custNumbInt++
	custNumb = strconv.Itoa(custNumbInt)
	data.Id = custNumb
	data.BaseType = "Party"
	data.Type = "Individual"
	sqlStmt = "INSERT INTO cs_cust " +
		"(cust_numb, frst_name, last_name, blpd_code, comp_code, cust_stts, id_type, id_numb,lang, grup_code, grup_levl, rprt_levl_flag, pmnt_levl_flag, grup_subr_indc, docm_addr_type, crtd_dttm, crtd_by,last_chng_dttm, last_chng_by) " +
		"VALUES ($1, $2, $3,'02',20,'A',$4,$5, 'T',0,1,'1','1','0','1',current_timestamp,'ADMIN',current_timestamp,'ADMIN')"
	_, sqlErr := s.DB.Exec(sqlStmt, data.Id, data.GivenName, data.FamilyName, data.IndividualIdentification[0].IdentificationType, data.IndividualIdentification[0].IdentificationId)
	if sqlErr != nil {
		lg := log.GenErrLog("SQL:"+sqlStmt, lt, log.E000000, sqlErr)
		log.AppTraceLog.Error(lg)
		omErr := util.NewOMError(lg)
		return c.JSON(http.StatusInternalServerError, omErr.ErrorReponsTMFJSON())
	}

	return c.JSON(http.StatusCreated, data)

}

func GetIndividualService(s *PartyHandler, c echo.Context, lt log.LogTracing) error {
	cond := make(map[string]interface{})
	sqlOrder := " ORDER BY CUST_NUMB "
	return getIndividual(s, c, sqlOrder, cond, lt)
}

func GetIndividualByIdService(s *PartyHandler, c echo.Context, id string, lt log.LogTracing) error {
	cond := make(map[string]interface{})
	if !util.IsNotEmptyString(id) {
		lg := log.GenErrLog("ID is empty", lt, log.E206247, nil)
		log.AppTraceLog.Error(lg)
		omErr := util.NewOMError(lg)
		return c.JSON(http.StatusBadRequest, omErr.ErrorReponsTMFJSON())
	}
	lt.CcbUser = id
	cond["cust_numb"] = id
	sqlOrder := " ORDER BY CUST_NUMB "
	return getIndividual(s, c, sqlOrder, cond, lt)
}

func getIndividual(s *PartyHandler, c echo.Context, sqlOrder string, cond map[string]interface{}, lt log.LogTracing) error {
	var values []interface{}
	var where []string
	i := 1
	for k, v := range cond {
		values = append(values, v)
		w := fmt.Sprintf("%s = %s%v", k, database.DB_CONST_TERM_VAR_PREFIX, i)
		where = append(where, w)
	}
	sqlStmt := "SELECT cust_numb,frst_name, last_name, id_type, id_numb FROM cs_cust"

	if len(where) > 0 {
		sqlStmt += " WHERE " + strings.Join(where, " AND ")
	}
	sqlStmt += sqlOrder
	sqlStmt = database.AddLimitOffset(sqlStmt, s.limit, s.offset)

	log.AppTraceLog.Debug(log.GenAppLog("Execute SQL: "+sqlStmt, lt))

	rows, sqlErr := s.DB.Queryx(sqlStmt, values...)
	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			lg := log.GenErrLog("SQL:"+sqlStmt, lt, log.E100017, sqlErr)
			log.AppTraceLog.Error(lg)
			omErr := util.NewOMError(lg)
			return c.JSON(http.StatusNotFound, omErr.ErrorReponsTMFJSON())
		}
		lg := log.GenErrLog("SQL:"+sqlStmt, lt, log.E000000, sqlErr)
		log.AppTraceLog.Error(lg)
		omErr := util.NewOMError(lg)
		return c.JSON(http.StatusInternalServerError, omErr.ErrorReponsTMFJSON())
	}

	dataSet := []IndividualData{}

	for rows.Next() {
		var data IndividualData
		var id individualIdentificationData
		rowErr := rows.Scan(&data.Id, &data.GivenName, &data.FamilyName, &id.IdentificationId, &id.IdentificationType)
		if rowErr != nil {
			lg := log.GenErrLog("SQL:"+sqlStmt, lt, log.E000000, rowErr)
			log.AppTraceLog.Error(lg)
			omErr := util.NewOMError(lg)
			return c.JSON(http.StatusInternalServerError, omErr.ErrorReponsTMFJSON())
		}
		data.Name = data.GivenName + " " + data.FamilyName

		data.IndividualIdentification = append(data.IndividualIdentification, id)
		apihelper.JSONOmitFilteredData(s.fields, &data)
		data.BaseType = "Party"
		data.Type = "Individual"
		dataSet = append(dataSet, data)
	}
	if len(dataSet) == 1 {
		return c.JSON(http.StatusOK, dataSet[0])
	}
	if len(dataSet) == 0 {
		return c.NoContent(http.StatusOK)
	}
	return c.JSON(http.StatusOK, dataSet)
}
