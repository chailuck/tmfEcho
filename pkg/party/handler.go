package party

import (
	"fmt"
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
	fmt.Println("filterField: ", filterField)
	h.fields = util.ConvertCommaStringToMap(filterField)
	limitStr := c.QueryParam("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		logmessage := log.GenAppLog("ERROR:limit is not int "+err.Error(), lt)
		log.AppTraceLog.Debug(logmessage)
		limit = maxLimit
	}
	h.limit = limit
	offsetStr := c.QueryParam("offset")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		logmessage := log.GenAppLog("ERROR:offset is not int "+err.Error(), lt)
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
