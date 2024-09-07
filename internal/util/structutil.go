package util

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"tmfEcho/internal/global"
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

func StructGenerateSQLUpdate(data interface{}, lt log.LogTracing) (map[string][]string, map[string][]interface{}, map[string][]string, OMError) {
	log.AppTraceLog.Debug(log.GenAppLog(fmt.Sprintf("Start StartGenerateSQLUpdate %v - %v\n", data, reflect.ValueOf(data).Kind()), lt))
	var val reflect.Value
	var mulWhereStmt map[string][]string
	var mulFieldStmt map[string][]string
	var mulValueStmt map[string][]interface{}
	sqlWhereMap := map[string][]string{}
	sqlFieldMap := map[string][]string{}
	sqlValueMap := map[string][]interface{}{}

	if reflect.ValueOf(data).Kind() == reflect.Ptr {
		val = reflect.ValueOf(data).Elem()
	} else {
		val = reflect.ValueOf(data)
	}

	for i := 0; i < val.NumField(); i++ {
		ft := val.Type().Field(i)
		f := val.Field(i)
		structField := ft.Name
		structValue := val.FieldByName(structField)
		structValueStr := strings.Trim(fmt.Sprintf("%v", structValue), " ")
		dbTableField := strings.Split(ft.Tag.Get("dbTable"), ",")[0]
		dbField := strings.Split(ft.Tag.Get("db"), ",")[0]
		jsonField := strings.Split(ft.Tag.Get("json"), ",")[0]
		sqlWhereTerm := ""
		sqlFieldTerm := ""
		sqlValueTerm := ""

		isGenerateSQL := false
		isMulStmt := false
		mulWhereStmt = map[string][]string{}
		mulFieldStmt = map[string][]string{}
		mulValueStmt = map[string][]interface{}{}
		switch f.Kind() {

		case reflect.String:
			if structValueStr != "" && dbField != "" && dbTableField != "" {
				if structValueStr == global.JSON_TERM_EMPTY_STR {
					structValueStr = ""
				}
				isGenerateSQL = true
				sqlWhereTerm = fmt.Sprintf("%v = '%v'", dbField, structValueStr)
				sqlFieldTerm = fmt.Sprintf("%v", dbField)
				sqlValueTerm = fmt.Sprintf("%v", structValueStr)
				log.AppTraceLog.Debug(log.GenAppLog(fmt.Sprintf("UPDATE: STR- %v ", structValueStr), lt))
				//f.SetString(structValueStr)
				log.AppTraceLog.Debug(log.GenAppLog(fmt.Sprintf("UPDATE: STR- %v  %v : %v (json: '%v', db:'%v'", structField, f.Kind(), structValueStr, jsonField, dbField), lt))
			}

		case reflect.Int:
			if structValueStr != "" && structValueStr != "0" && dbField != "" && dbTableField != "" {

				valInt, err := strconv.ParseInt(structValueStr, 10, 64)
				if err == nil {
					if valInt == global.JSON_TERM_ZERO_INT {
						structValueStr = "0"
						valInt = 0
					}
					//f.SetInt(valInt)
					isGenerateSQL = true
					sqlWhereTerm = fmt.Sprintf("%v = %v", dbField, structValueStr)
					sqlFieldTerm = fmt.Sprintf("%v", dbField)
					sqlValueTerm = fmt.Sprintf("%v", structValueStr)

					log.AppTraceLog.Debug(log.GenAppLog(fmt.Sprintf("UPDATE: INT- %v  %v : %v (json: '%v', db:'%v'", structField, f.Kind(), structValueStr, jsonField, dbField), lt))
				}
			}

		case reflect.Int64:
			if structValueStr != "" && structValueStr != "0" && dbField != "" && dbTableField != "" {
				valInt, err := strconv.ParseInt(structValueStr, 10, 64)
				if err == nil {
					if valInt == global.JSON_TERM_ZERO_INT {
						structValueStr = "0"
						valInt = 0
					}
					//f.SetInt(valInt)
					isGenerateSQL = true
					sqlWhereTerm = fmt.Sprintf("%v = %v", dbField, structValueStr)
					sqlFieldTerm = fmt.Sprintf("%v", dbField)
					sqlValueTerm = fmt.Sprintf("%v", structValueStr)
					log.AppTraceLog.Debug(log.GenAppLog(fmt.Sprintf("UPDATE: I64- %v  %v : %v (json: '%v', db:'%v'", structField, f.Kind(), structValueStr, jsonField, dbField), lt))
				}
			}
		case reflect.Slice:
			log.AppTraceLog.Debug(log.GenAppLog(fmt.Sprintf("SLI- %v  %v : %v (json: '%v', db:'%v'", structField, f.Kind(), structValueStr, jsonField, dbField), lt))
			cap := structValue.Cap()
			for i := 0; i < cap; i++ {
				sliceValue := f.Slice(0, cap).Index(i)
				var omErr OMError
				mulFieldStmt, mulValueStmt, mulWhereStmt, omErr = StructGenerateSQLUpdate(sliceValue.Interface(), lt)

				if omErr.Err != nil {
					return nil, nil, nil, omErr
				}

				if mulFieldStmt != nil {
					isMulStmt = true
					isGenerateSQL = true
				}
				log.AppTraceLog.Debug(log.GenAppLog(fmt.Sprintf("--> RETURN MAP:%v (isMulStmt: %v) ", sqlFieldMap, isMulStmt), lt))
			}
		case reflect.Struct:
			log.AppTraceLog.Debug(log.GenAppLog(fmt.Sprintf("ARY- %v  %v : %v (json: '%v', db:'%v'", structField, f.Kind(), structValueStr, jsonField, dbField), lt))
			var omErr OMError
			mulFieldStmt, mulValueStmt, mulWhereStmt, omErr = StructGenerateSQLUpdate(&structValue, lt)
			if omErr.Err != nil {
				return nil, nil, nil, omErr
			}
			if mulFieldStmt != nil {
				isMulStmt = true
				isGenerateSQL = true
			}
		default:
			log.AppTraceLog.Debug(log.GenAppLog(fmt.Sprintf("DEF- %v  %v : %v (json: '%v', db:'%v'", structField, f.Kind(), structValueStr, jsonField, dbField), lt))

		}
		if isGenerateSQL {

			if !isMulStmt {
				sqlFieldSlice, ok := sqlFieldMap[dbTableField]
				if !ok {
					sqlFieldMap[dbTableField] = []string{}
				}
				sqlFieldSlice = append(sqlFieldSlice, sqlFieldTerm)
				sqlFieldMap[dbTableField] = sqlFieldSlice

				sqlValueSlide, ok := sqlValueMap[dbTableField]
				if !ok {
					sqlValueMap[dbTableField] = []interface{}{}
				}
				sqlValueSlide = append(sqlValueSlide, sqlValueTerm)
				sqlValueMap[dbTableField] = sqlValueSlide

				sqlWhereSlide, ok := sqlWhereMap[dbTableField]
				if !ok {
					sqlWhereMap[dbTableField] = []string{}
				}
				sqlWhereSlide = append(sqlWhereSlide, sqlWhereTerm)
				sqlWhereMap[dbTableField] = sqlWhereSlide

			} else {
				log.AppTraceLog.Debug(log.GenAppLog(fmt.Sprintf("--> APPEND MAP:%v to %v", sqlFieldMap, mulFieldStmt), lt))

				for k, v := range mulFieldStmt {
					sqlFieldSlice, ok := sqlFieldMap[k]
					if !ok {
						sqlFieldMap[k] = []string{}
						log.AppTraceLog.Debug(log.GenAppLog(fmt.Sprintf("-->INITIAL SQLSLICE:%v", k), lt))
					}
					sqlFieldSlice = append(sqlFieldSlice, v...)
					log.AppTraceLog.Debug(log.GenAppLog(fmt.Sprintf("-->ADD SQL SLICE:%v to %v", v, sqlFieldSlice), lt))
					sqlFieldMap[k] = sqlFieldSlice
				}
				for k, v := range mulValueStmt {
					sqlValueSlice, ok := sqlValueMap[k]
					if !ok {
						sqlValueMap[k] = []interface{}{}
						log.AppTraceLog.Debug(log.GenAppLog(fmt.Sprintf("-->INITIAL SQLSLICE:%v", k), lt))
					}
					sqlValueSlice = append(sqlValueSlice, v...)
					log.AppTraceLog.Debug(log.GenAppLog(fmt.Sprintf("-->ADD SQL SLICE:%v to %v", v, sqlValueSlice), lt))
					sqlValueMap[k] = sqlValueSlice
				}
				for k, v := range mulWhereStmt {
					sqlWhereSlice, ok := sqlWhereMap[k]
					if !ok {
						sqlWhereMap[k] = []string{}
						log.AppTraceLog.Debug(log.GenAppLog(fmt.Sprintf("-->INITIAL SQLSLICE:%v", k), lt))
					}
					sqlWhereSlice = append(sqlWhereSlice, v...)
					log.AppTraceLog.Debug(log.GenAppLog(fmt.Sprintf("-->ADD SQL SLICE:%v to %v", v, sqlWhereSlice), lt))
					sqlWhereMap[k] = sqlWhereSlice
				}
			}
			log.AppTraceLog.Debug(log.GenAppLog(fmt.Sprintf("-->ASSING VAL %v MAP: %v SQLSLICE:%v", dbField, sqlFieldMap, mulFieldStmt), lt))
		}

	}
	return sqlFieldMap, sqlValueMap, sqlWhereMap, OMError{}
}
