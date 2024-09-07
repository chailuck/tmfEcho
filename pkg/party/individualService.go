package party

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"tmfEcho/internal/database"
	"tmfEcho/internal/log"
	"tmfEcho/internal/util"

	"github.com/labstack/echo"
)

type IndividualData struct {
	Id                       string                         `json:"id" db:"cust_numb" dbTable:"CS_CUST"`
	Type                     string                         `json:"@type"  db:"" dbTable:""`
	BaseType                 string                         `json:"@baseType"  db:"" dbTable:"CS_CUST"`
	GivenName                string                         `json:"givenName"  db:"frst_name"  validate:"required" dbTable:"CS_CUST"`
	FamilyName               string                         `json:"familyName,omitempty"  db:"last_name" dbTable:"CS_CUST"`
	Name                     string                         `json:"name,omitempty"  db:"" dbTable:"CS_CUST"`
	Title                    string                         `json:"title,omitempty" db:"titl" dbTable:"CS_CUST"`
	IndividualIdentification []individualIdentificationData `json:"individualIdentification,omitempty" db:""  maxArray:"1"`
	ContactMedium            []contactMediumData            `json:"contactMedium,omitempty" db:""`
}

type individualIdentificationData struct {
	Type               string `json:"@type"  db:"" dbTable:""`
	IdentificationId   string `json:"identificationId,omitempty" db:"id_type"  dbTable:"CS_CUST"`
	IdentificationType string `json:"identificiationType,omitempty" db:"id_numb" validate:"required"  dbTable:"CS_CUST"`
}

type contactMediumData struct {
	ContactType     string `json:"contactType" db:""`
	Preferred       bool   `json:"" db:""`
	City            string `json:"city" db:"AMPR_DESC" dbTable:"CS_PSCD"`
	Country         string `json:"country" db:"CNTRY_CODE" dbTable:"CS_CNTY"`
	PostCode        string `json:"postCode" db:"POST_CODE" dbTable:"CS_CSAD"`
	StateOrProvince string `json:"stateOrProvince" db:"PNVC_DESC" dbTable:"CS_PVNC"`
	Street1         string `json:"street1" db:"ADR1" dbTable:"CS_CSSAD"`
	Street2         string `json:"street2" db:"ADR2" dbTable:"CS_CSSAD"`
	PostCodeID      string `json:"postCodeID" db:"POST_CODE_SEQN" dbTable:"CS_CSAD"`
	EmailAddress    string `json:"emailAddress" db:"EMAL_ADDR" dbTable:"CS_CUST"`
	PhoneNumber     string `json:"phoneNumber" db:"HOME_TELP_NUMB" dbTable:"CS_CUST"`
	FaxNumber       string `json:"faxNumber" db:"HOME_FAX_NUMB" dbTable:"CS_CUST"`
	SocialNetworkId string `json:"socialNetworkId" db:""`
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
		return c.JSON(http.StatusInternalServerError, util.NewOMError(lg).WriteLog().ErrorReponsTMFJSON())
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

	requestMap := make(map[string]interface{})

	log.AppTraceLog.Debug(log.GenAppLog(fmt.Sprintf("BEGIN JSON bind to STRUCT: %v | %v\n", c, data), lt))
	if sErr := c.Bind(&data); sErr != nil {
		lg := log.GenErrLog("Wrong Request payload (Binding STRUCT)", lt, log.E201434, sErr)
		return c.JSON(http.StatusBadRequest, util.NewOMError(lg).WriteLog().ErrorReponsTMFJSON())
	}

	log.AppTraceLog.Debug(log.GenAppLog(fmt.Sprintf("CONTEXT MAP: %v\n", requestMap), lt))
	if !util.IsNotEmptyString(id) {
		lg := log.GenErrLog("ID is empty", lt, log.E206247, nil)
		return c.JSON(http.StatusBadRequest, util.NewOMError(lg).WriteLog().ErrorReponsTMFJSON())
	}

	var sqlWhereMap map[string][]string
	var sqlFieldMap map[string][]string
	var sqlValueMap map[string][]interface{}
	sqlStmt := ""

	sqlFieldMap, sqlValueMap, sqlWhereMap, OMerr := util.StructGenerateSQLUpdate(data, lt)
	if OMerr.Err != nil {
		lg := log.GenErrLog("Wrong Request payload (Binding SQL Update)", lt, log.E201434, OMerr.Err)
		return c.JSON(http.StatusBadRequest, util.NewOMError(lg).WriteLog().ErrorReponsTMFJSON())
	}
	log.AppTraceLog.Debug(log.GenAppLog(fmt.Sprintf("SQL WHERE MAP: %v\n", sqlWhereMap), lt))
	// BEGIN TRANSACTIONS
	ctx := c.Request().Context()
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		lg := log.GenErrLog("Begin DB transaction: ", lt, log.E000000, err)
		return c.JSON(http.StatusInternalServerError, util.NewOMError(lg).WriteLog().ErrorReponsTMFJSON())
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	for k, v := range sqlFieldMap {
		sqlStmt = ""
		cond := []string{}
		for _, valueTerm := range v {
			m := fmt.Sprintf("%v=%v%v", valueTerm, database.DB_CONST_TERM_VAR_PREFIX, strconv.Itoa(len(cond)+1))
			cond = append(cond, m)
		}
		switch k {
		case "CS_CUST":
			sqlStmt = "update " + k + " SET " + strings.Join(cond, ",")
			sqlStmt += " WHERE cust_numb = " + id

		default:
		}
		log.AppTraceLog.Debug(log.GenAppLog("SQL STMT:"+sqlStmt, lt))

		if sqlStmt != "" {
			stmt1, err := tx.Prepare(sqlStmt)
			if err != nil {
				lg := log.GenErrLog("SQL : "+sqlStmt, lt, log.E000000, err)
				return c.JSON(http.StatusInternalServerError, util.NewOMError(lg).WriteLog().ErrorReponsTMFJSON())
			}
			defer stmt1.Close()
			log.AppTraceLog.Debug(log.GenAppLog("Execute SQL: "+sqlStmt, lt))

			sqlValueSlice, ok := sqlValueMap[k]
			if !ok {
				sqlValueMap[k] = []interface{}{}
				log.AppTraceLog.Debug(log.GenAppLog(fmt.Sprintf("-->INITIAL SQLSLICE:%v", k), lt))
			}
			_, err = stmt1.Exec(sqlValueSlice...)
			if err != nil {
				lg := log.GenErrLog("SQL : "+sqlStmt, lt, log.E000000, err)
				return c.JSON(http.StatusInternalServerError, util.NewOMError(lg).WriteLog().ErrorReponsTMFJSON())
			}
			log.AppTraceLog.Debug(log.GenAppLog("SQL EXECUTE DONE:"+sqlStmt, lt))
		}
	}

	if err := tx.Commit(); err != nil {
		lg := log.GenErrLog("Commit transaction", lt, log.E000000, err)
		return c.JSON(http.StatusInternalServerError, util.NewOMError(lg).WriteLog().ErrorReponsTMFJSON())
	} else {
		log.AppTraceLog.Debug(log.GenAppLog("COMMITTED UpdateIndividualService", lt))
	}
	return getIndividual(s, c, sqlOrder, cond, lt)
	//return c.JSON(http.StatusOK, data)
}

func SaveIndividualService(s *PartyHandler, c echo.Context, lt log.LogTracing) error {
	var data IndividualData
	if err := c.Bind(&data); err != nil {
		lg := log.GenErrLog("Wrong Request payload", lt, log.E201434, err)
		return c.JSON(http.StatusBadRequest, util.NewOMError(lg).WriteLog().ErrorReponsTMFJSON())
	}
	//validate required fields
	omErr := util.ValidateStruct(&data, lt)

	if omErr.Err != nil {
		lg := log.GenErrLog("Wrong Request payload", lt, log.E201434, omErr.Err)
		return c.JSON(http.StatusBadRequest, util.NewOMError(lg).WriteLog().ErrorReponsTMFJSON())
	}

	sqlStmt := "select max(cust_numb) from cs_cust"
	var custNumb string
	err := s.DB.QueryRow(sqlStmt).Scan(&custNumb)
	if err != nil {
		if err == sql.ErrNoRows {
			custNumb = "100000"
		} else {
			lg := log.GenErrLog("SQL:"+sqlStmt, lt, log.E000000, err)
			return c.JSON(http.StatusInternalServerError, util.NewOMError(lg).WriteLog().ErrorReponsTMFJSON())
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
		return c.JSON(http.StatusInternalServerError, util.NewOMError(lg).WriteLog().ErrorReponsTMFJSON())
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
		id.IdentificationId = strings.Trim(id.IdentificationId, " ")
		id.IdentificationType = strings.Trim(id.IdentificationType, " ")
		id.Type = "IndividualIdentification"
		data.IndividualIdentification = append(data.IndividualIdentification, id)
		util.JSONOmitFilteredData(s.fields, &data)
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
