package util

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"tmfEcho/internal/log"

	"github.com/go-playground/validator"
)

func ValidateStruct(data interface{}, lt log.LogTracing) OMError {
	log.AppTraceLog.Debug(log.GenAppLog(fmt.Sprintf("Start Validation %v - %v\n", data, reflect.ValueOf(data).Kind()), lt))
	var val reflect.Value
	if reflect.ValueOf(data).Kind() == reflect.Ptr {
		val = reflect.ValueOf(data).Elem()
	} else {
		val = reflect.ValueOf(data)
	}

	validate := validator.New()
	vErr := validate.Struct(data)
	if vErr != nil {
		lg := log.GenErrLog("Required fields", lt, log.E100009, vErr)
		log.AppTraceLog.Error(lg)
		omErr := NewOMError(lg)
		return omErr
	}

	for i := 0; i < val.NumField(); i++ {
		ft := val.Type().Field(i)
		f := val.Field(i)
		structField := ft.Name
		structValue := val.FieldByName(structField)
		structvalueStr := fmt.Sprintf("%v", structValue)
		dbField := strings.Split(ft.Tag.Get("db"), ",")[0]
		jsonField := strings.Split(ft.Tag.Get("json"), ",")[0]
		//isMapDB := util.IsNotEmptyString(dbField)
		//log.AppTraceLog.Debug(log.GenAppLog(fmt.Sprintf("1) json: %v, db: %v , val: %v", jsonField, dbField, fieldFilter[jsonField]), lt))
		switch f.Kind() {

		case reflect.String:
			log.AppTraceLog.Debug(log.GenAppLog(fmt.Sprintf("STR- %v  %v : %v (json: '%v', db:'%v'", structField, f.Kind(), structvalueStr, jsonField, dbField), lt))
		case reflect.Int:
			log.AppTraceLog.Debug(log.GenAppLog(fmt.Sprintf("INT- %v  %v : %v (json: '%v', db:'%v'", structField, f.Kind(), structvalueStr, jsonField, dbField), lt))
		case reflect.Int64:
			log.AppTraceLog.Debug(log.GenAppLog(fmt.Sprintf("I64- %v  %v : %v (json: '%v', db:'%v'", structField, f.Kind(), structvalueStr, jsonField, dbField), lt))
		case reflect.Slice:
			log.AppTraceLog.Debug(log.GenAppLog(fmt.Sprintf("SLI- %v  %v : %v (json: '%v', db:'%v'", structField, f.Kind(), structvalueStr, jsonField, dbField), lt))
			cap := structValue.Cap()
			maxArray := ft.Tag.Get("maxArray")
			if maxArray != "" {
				iMaxArray, err := strconv.Atoi(ft.Tag.Get("maxArray"))
				if err != nil {
					lg := log.GenErrLog("Required fields", lt, log.E100009, err)
					log.AppTraceLog.Error(lg)
					omErr := NewOMError(lg)
					return omErr
				}

				if cap > iMaxArray {
					lg := log.GenErrLog("Array instances is higher than "+maxArray, lt, log.E207196, nil)
					log.AppTraceLog.Error(lg)
					omErr := NewOMError(lg)
					return omErr
				}
			}
			for i := 0; i < cap; i++ {
				sliceValue := f.Slice(0, cap).Index(i)
				omErr := ValidateStruct(sliceValue.Interface(), lt)
				if omErr.Err != nil {
					return omErr
				}
			}

		case reflect.Struct:
			log.AppTraceLog.Debug(log.GenAppLog(fmt.Sprintf("ARY- %v  %v : %v (json: '%v', db:'%v'", structField, f.Kind(), structvalueStr, jsonField, dbField), lt))
			omErr := ValidateStruct(&structValue, lt)
			if omErr.Err != nil {
				return omErr
			}

		default:
			log.AppTraceLog.Debug(log.GenAppLog(fmt.Sprintf("DEF- %v  %v : %v (json: '%v', db:'%v'", structField, f.Kind(), structvalueStr, jsonField, dbField), lt))

		}

	}
	return OMError{}
}
