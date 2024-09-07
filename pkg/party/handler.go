package party

import (
	"tmfEcho/internal/log"
	"tmfEcho/internal/util"

	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

const TMFno = "TMF632"
const maxLimit = 50

type PartyHandler struct {
	DB     *sqlx.DB
	fields map[string]bool
	limit  int
	offset int
}

func (h *PartyHandler) Initialize(db *sqlx.DB) {

	//db.AutoMigrate(&Customer{})
	h.DB = db

}

func (h *PartyHandler) callAPIStart(c echo.Context, lt log.LogTracing) {
	logMessage := log.GenAppLog("START CALL:"+lt.ApiName, lt)
	logMessage.SetStartTime()
	log.AppTraceLog.Info(logMessage)

	filterField := c.QueryParam("fields")
	//fmt.Println("filterField: ", filterField)
	h.fields = util.ConvertCommaStringToMap(filterField)
	limitStr := c.QueryParam("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		logmessage := log.GenAppLog("limit is not int "+err.Error(), lt)
		log.AppTraceLog.Debug(logmessage)
		limit = maxLimit
	}
	h.limit = limit
	offsetStr := c.QueryParam("offset")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		logmessage := log.GenAppLog("offset is not int "+err.Error(), lt)
		log.AppTraceLog.Debug(logmessage)
		offset = -1
	}
	h.offset = offset
}

func (h *PartyHandler) SaveIndividual(c echo.Context) error {
	lt := log.LogTracing{ApiName: TMFno + "-" + "SaveIndividual"}
	h.callAPIStart(c, lt)
	return SaveIndividualService(h, c, lt)
}

func (h *PartyHandler) UpdateIndividual(c echo.Context) error {
	lt := log.LogTracing{ApiName: TMFno + "-" + "UpdateIndividual"}
	h.callAPIStart(c, lt)
	id := c.Param("id")
	return UpdateIndividualService(h, c, id, lt)
}

func (h *PartyHandler) GetIndividual(c echo.Context) error {
	lt := log.LogTracing{ApiName: TMFno + "-" + "GetIndividual"}

	h.callAPIStart(c, lt)
	return GetIndividualService(h, c, lt)
}

func (h *PartyHandler) GetIndividualById(c echo.Context) error {
	id := c.Param("id")
	lt := log.LogTracing{ApiName: TMFno + "-" + "GetIndividualById", CustNumb: id}
	h.callAPIStart(c, lt)
	return GetIndividualByIdService(h, c, id, lt)
}

func (h *PartyHandler) DeleteIndividual(c echo.Context) error {
	id := c.Param("id")
	lt := log.LogTracing{ApiName: TMFno + "-" + "DeleteIndividualService", CustNumb: id}
	h.callAPIStart(c, lt)
	return DeleteIndividualService(h, c, id, lt)
}

func (h *PartyHandler) SaveOrganization(c echo.Context) error {
	lt := log.LogTracing{ApiName: TMFno + "-" + "SaveOrganization"}
	h.callAPIStart(c, lt)
	return SaveOrganizationService(h, c, lt)
}

func (h *PartyHandler) UpdateOrganization(c echo.Context) error {
	lt := log.LogTracing{ApiName: TMFno + "-" + "UpdateOrganization"}
	h.callAPIStart(c, lt)
	id := c.Param("id")
	return UpdateOrganizationService(h, c, id, lt)
}

func (h *PartyHandler) GetOrganization(c echo.Context) error {
	lt := log.LogTracing{ApiName: TMFno + "-" + "GetOrganization"}

	h.callAPIStart(c, lt)
	return GetOrganizationService(h, c, lt)
}

func (h *PartyHandler) GetOrganizationById(c echo.Context) error {
	id := c.Param("id")
	lt := log.LogTracing{ApiName: TMFno + "-" + "GetOrganizationById", CustNumb: id}
	h.callAPIStart(c, lt)
	return GetOrganizationByIdService(h, c, id, lt)
}

func (h *PartyHandler) DeleteOrganization(c echo.Context) error {
	id := c.Param("id")
	lt := log.LogTracing{ApiName: TMFno + "-" + "DeleteOrganizationService", CustNumb: id}
	h.callAPIStart(c, lt)
	return DeleteOrganizationService(h, c, id, lt)
}
