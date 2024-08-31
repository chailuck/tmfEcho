package party

import (
	"database/sql"
	"net/http"
	"tmfEcho/internal/api/apihelper"
	"tmfEcho/internal/database"
	"tmfEcho/internal/log"
	"tmfEcho/internal/util"

	"github.com/labstack/echo"
)

type IndividualData struct {
	Id         string `json:"id,omitempty"`
	GivenName  string `json:"givenName,omitempty"`
	FamilyName string `json:"familyName,omitempty"`
	Name       string `json:"name,omitempty"`
	Age        int    `json:"age,omitempty"`
}

// func omitFilteredData(fieldFilter map[string]bool, data *IndividualData) {

func GetIndividualService(s *PartyHandler, c echo.Context, lt log.LogTracing) error {
	//a := IndividualData{Id: "1", GivenName: "John", FamilyName: "Doe", Name: "John Doe"}
	data := IndividualData{Age: 40}
	sqlStmt := "SELECT cust_numb, frst_name, last_name " +
		"FROM cs_cust " +
		"ORDER BY CUST_NUMB "
	sqlStmt = database.AddLimitOffset(sqlStmt, s.limit, s.offset)

	log.AppTraceLog.Debug(log.GenAppLog("Execute SQL: "+sqlStmt, lt))
	sqlErr := s.DB.QueryRowx(sqlStmt).Scan(&data.Id, &data.GivenName, &data.FamilyName)
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
	data.Name = data.GivenName + " " + data.FamilyName

	apihelper.JSONOmitFilteredData(s.fields, &data)
	return c.JSON(http.StatusOK, data)
}

func GetIndividualByIdService(s *PartyHandler, c echo.Context, id string, lt log.LogTracing) error {
	data := IndividualData{Id: id, Age: 40}
	sqlStmt := "SELECT frst_name, last_name " +
		"FROM cs_cust" +
		"WHERE cust_numb = " + database.DB_CONST_TERM_VAR_PREFIX + "1" +
		"ORDER BY CUST_NUMB "
	sqlStmt = database.AddLimitOffset(sqlStmt, s.limit, s.offset)

	log.AppTraceLog.Debug(log.GenAppLog("Execute SQL: "+sqlStmt, lt))
	sqlErr := s.DB.QueryRowx(sqlStmt, id).Scan(&data.GivenName, &data.FamilyName)
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
	data.Name = data.GivenName + " " + data.FamilyName

	//a := IndividualData{Id: id, GivenName: "John", FamilyName: "Doe", Name: "John Doe", Age: 30}
	apihelper.JSONOmitFilteredData(s.fields, &data)
	return c.JSON(http.StatusOK, data)
}
