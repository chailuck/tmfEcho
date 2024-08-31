package apihelper

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"tmfEcho/internal/log"
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
