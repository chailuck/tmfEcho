package util

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"tmfEcho/internal/log"
)

func JSONOmitFilteredData(fieldFilter map[string]bool, data interface{}) {

	//fmt.Println("omit Start fieldFilter: ", fieldFilter)
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

func JSONconverToUpdateValue(fieldFilter map[string]interface{}, data interface{}, lt log.LogTracing) ([]string, OMError) {
	var sqlUpdate []string
	var val reflect.Value
	assgVal := map[string][]string{}
	var valStrSlice []string

	//fmt.Println("omit Start fieldFilter: ", fieldFilter)
	lenReq := len(fieldFilter)
	if lenReq != 0 {
		if reflect.ValueOf(data).Kind() == reflect.Ptr {
			val = reflect.ValueOf(data).Elem()
		} else {
			val = reflect.ValueOf(data)
		}
		//fmt.Printf("Kind: %v\n", reflect.ValueOf(data).Kind())
		//fmt.Printf("VAL: %v\n", val)
		//fmt.Printf("DATA: %v\n", data)
		for i := 0; i < val.NumField(); i++ {
			ft := val.Type().Field(i)
			f := val.Field(i)

			dbField := strings.Split(ft.Tag.Get("db"), ",")[0]
			dbTableField := strings.Split(ft.Tag.Get("dbTable"), ",")[0]
			jsonField := strings.Split(ft.Tag.Get("json"), ",")[0]
			log.AppTraceLog.Debug(log.GenAppLog(fmt.Sprintf("CHECK DATA: json: %v, db: %v , val: %v, kind: %v", jsonField, dbField, fieldFilter[jsonField], f.Kind()), lt))
			isMapDB := IsNotEmptyString(dbField)
			isSliceMap := false
			valField, ok := fieldFilter[jsonField]
			if ok { // there are fields in request message

				sqlTerm := ""
				valStr := fmt.Sprintf("%v", valField)
				switch f.Kind() {
				case reflect.String:
					f.SetString(valStr)
				case reflect.Int:
					valInt, err := strconv.ParseInt(valStr, 10, 64)
					if err != nil {
						lg := log.GenErrLog("COVERT ERROR:"+valStr, lt, log.E000000, err)
						log.AppTraceLog.Error(lg)
						omErr := NewOMError(lg)
						return nil, omErr
					}
					f.SetInt(valInt)
					/*
						case reflect.Slice:

							structField := ft.Name
							structValue := val.FieldByName(structField)
							cap := len(valField.([]interface{}))

							for i := 0; i < cap; i++ {
								fmt.Printf("SLICE TYPE %v)%v, %v, %v, %v, %v\n", i, valField, structField, cap, structValue, f)

								//filterValue := valField[0].Interface().(map[string]interface{})
								filterValue := reflect.ValueOf(valField).Index(i).Interface().(map[string]interface{})
								//fmt.Printf("SLICE: %v\n", sliceValue)
								//sliceValue := structValue.Slice(0, cap).Index(i)

								fVal := reflect.New(fType)

								//newFilter := sliceValue.Slice(0,i)
								valStrSlice, omErr := JSONconverToUpdateValue(filterValue, &fVal, lt)
								isSliceMap = true
								if omErr.Err != nil {
									return nil, omErr
								}
								fmt.Printf("VAL SLIDE: %v\n", valStrSlice)
							}
					*/
				default:

				}

				if isMapDB {
					if isSliceMap && len(valStrSlice) > 0 {
						sqlUpdate = append(sqlUpdate, valStrSlice...)
					} else {
						sqlTerm = dbField + " = '" + valStr + "' "
						sqlUpdate = append(sqlUpdate, sqlTerm)
					}

					sqlSlice, ok := assgVal[dbTableField]
					if !ok {
						assgVal[dbTableField] = []string{}
					}
					sqlSlice = append(sqlSlice, sqlTerm)
					assgVal[dbTableField] = sqlSlice
					log.AppTraceLog.Debug(log.GenAppLog(fmt.Sprintf("-->ASSING VAL %v MAP: %v SQLSLICE:%v", dbField, assgVal, sqlSlice), lt))

				}

				log.AppTraceLog.Debug(log.GenAppLog("UPDATE TERM: NAME: "+valStr+" JSON:"+jsonField+" (SQL "+sqlTerm+")", lt))
			}

		}
	}

	return sqlUpdate, OMError{}
}
