package apihelper

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"tmfEcho/internal/log"
	"tmfEcho/internal/util"
)

func JSONOmitFilteredData(fieldFilter map[string]bool, data interface{}) {

	fmt.Println("omit Start fieldFilter: ", fieldFilter)
	if len(fieldFilter) != 0 {
		val := reflect.ValueOf(data).Elem()

		for i := 0; i < val.NumField(); i++ {
			ft := val.Type().Field(i)
			f := val.Field(i)
			jsonField := strings.Split(ft.Tag.Get("json"), ",")[0]
			valBool, ok := fieldFilter[jsonField]
			if !ok {
				if f.Kind() == reflect.String {
					f.SetString("")
					log.AppTraceLog.Debug(log.GenAppLog("AFTER:"+jsonField+" - "+f.String()+" - "+strconv.FormatBool(valBool), log.LogTracing{}))
				} else if f.Kind() == reflect.Int {
					f.SetInt(0)
					log.AppTraceLog.Debug(log.GenAppLog("AFTER:"+jsonField+" - "+strconv.Itoa(int(f.Int()))+" - "+strconv.FormatBool(valBool), log.LogTracing{}))
				}
			}
		}
	}
}

func JSONconverToUpdateValue(fieldFilter map[string]interface{}, data interface{}, lt log.LogTracing) ([]string, util.OMError) {
	var sqlUpdate []string
	fmt.Println("omit Start fieldFilter: ", fieldFilter)
	lenReq := len(fieldFilter)
	if lenReq != 0 {
		val := reflect.ValueOf(data).Elem()

		for i := 0; i < val.NumField(); i++ {
			ft := val.Type().Field(i)
			f := val.Field(i)
			dbField := strings.Split(ft.Tag.Get("db"), ",")[0]
			jsonField := strings.Split(ft.Tag.Get("json"), ",")[0]
			log.AppTraceLog.Debug(log.GenAppLog(fmt.Sprintf("1) json: %v, db: %v , val: %v", jsonField, dbField, fieldFilter[jsonField]), lt))
			isMapDB := util.IsNotEmptyString(dbField)
			val, ok := fieldFilter[jsonField]
			if ok { // there are fields in request message

				sqlTerm := ""
				valStr := fmt.Sprintf("%v", val)
				if f.Kind() == reflect.String {
					f.SetString(valStr)
				} else if f.Kind() == reflect.Int {
					valInt, err := strconv.ParseInt(valStr, 10, 64)
					if err != nil {
						lg := log.GenErrLog("COVERT ERROR:"+valStr, lt, log.E000000, err)
						log.AppTraceLog.Error(lg)
						omErr := util.NewOMError(lg)
						return nil, omErr
					}
					f.SetInt(valInt)
				}
				if isMapDB {
					sqlTerm = dbField + " = '" + valStr + "' "
					sqlUpdate = append(sqlUpdate, sqlTerm)
				}

				log.AppTraceLog.Debug(log.GenAppLog("SQL TERM: NAME:"+valStr+" JSON:"+jsonField+"SQL "+sqlTerm, lt))
			}

		}
	}

	return sqlUpdate, util.OMError{}
}
