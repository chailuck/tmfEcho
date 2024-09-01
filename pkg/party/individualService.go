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

	"github.com/labstack/echo"
)

type IndividualData struct {
	Id         string `json:"id,omitempty"`
	Type       string `json:"@type"`
	BaseType   string `json:"@baseType"`
	GivenName  string `json:"givenName,omitempty"`
	FamilyName string `json:"familyName,omitempty"`
	Name       string `json:"name,omitempty"`
}

func SaveIndividualService(s *PartyHandler, c echo.Context, lt log.LogTracing) error {
	var data IndividualData
	if err := c.Bind(&data); err != nil {
		lg := log.GenErrLog("Wrong Request payload", lt, log.E201434, err)
		log.AppTraceLog.Error(lg)
		omErr := util.NewOMError(lg)
		return c.JSON(http.StatusBadRequest, omErr.ErrorReponsTMFJSON())
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
		"VALUES ($1, $2, $3,'02',20,'A','01','-', 'T',0,1,'1','1','0','1',current_timestamp,'ADMIN',current_timestamp,'ADMIN')"
	_, sqlErr := s.DB.Exec(sqlStmt, data.Id, data.GivenName, data.FamilyName)
	if sqlErr != nil {
		lg := log.GenErrLog("SQL:"+sqlStmt, lt, log.E000000, sqlErr)
		log.AppTraceLog.Error(lg)
		omErr := util.NewOMError(lg)
		return c.JSON(http.StatusInternalServerError, omErr.ErrorReponsTMFJSON())
	}

	return c.JSON(http.StatusOK, data)

}

func GetIndividualService(s *PartyHandler, c echo.Context, lt log.LogTracing) error {
	cond := make(map[string]interface{})
	sqlOrder := " ORDER BY CUST_NUMB "
	return get(s, c, sqlOrder, cond, lt)
}

func GetIndividualByIdService(s *PartyHandler, c echo.Context, id string, lt log.LogTracing) error {
	cond := make(map[string]interface{})
	if !util.IsNotEmptyString(id) {
		lg := log.GenErrLog("ID is empty", lt, log.E000000, nil)
		log.AppTraceLog.Error(lg)
		omErr := util.NewOMError(lg)
		return c.JSON(http.StatusBadRequest, omErr.ErrorReponsTMFJSON())
	}
	lt.CcbUser = id
	cond["cust_numb"] = id
	sqlOrder := " ORDER BY CUST_NUMB "
	return get(s, c, sqlOrder, cond, lt)
}

func get(s *PartyHandler, c echo.Context, sqlOrder string, cond map[string]interface{}, lt log.LogTracing) error {
	var values []interface{}
	var where []string
	i := 1
	for k, v := range cond {
		values = append(values, v)
		w := fmt.Sprintf("%s = %s%v", k, database.DB_CONST_TERM_VAR_PREFIX, i)
		where = append(where, w)
	}
	sqlStmt := "SELECT cust_numb,frst_name, last_name FROM cs_cust"

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
		rowErr := rows.Scan(&data.Id, &data.GivenName, &data.FamilyName)
		if rowErr != nil {
			lg := log.GenErrLog("SQL:"+sqlStmt, lt, log.E000000, rowErr)
			log.AppTraceLog.Error(lg)
			omErr := util.NewOMError(lg)
			return c.JSON(http.StatusInternalServerError, omErr.ErrorReponsTMFJSON())
		}
		data.Name = data.GivenName + " " + data.FamilyName
		apihelper.JSONOmitFilteredData(s.fields, &data)
		data.BaseType = "Party"
		data.Type = "Individual"
		dataSet = append(dataSet, data)
	}
	//sqlErr := s.DB.QueryRowx(sqlStmt, values...).Scan(&data.Id, &data.GivenName, &data.FamilyName)
	if len(dataSet) == 1 {
		return c.JSON(http.StatusOK, dataSet[0])
	}
	if len(dataSet) == 0 {
		return c.NoContent(http.StatusOK)
	}
	return c.JSON(http.StatusOK, dataSet)
}
