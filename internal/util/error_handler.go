package util

import (
	"fmt"
	"path/filepath"
	"runtime"
	"tmfEcho/internal/log"
)

const (
	EXPECTED_ERROR   = "EXPECTED"
	UNEXPECTED_ERROR = "UNEXPECTED"
)

/*
	type OMError interface {
		Error() string
		NewError(eType string, eCode string, er error)
	}
*/
type OMError struct {
	ErrType    string
	ErrCode    string
	ErrText    string
	Path       string
	LineOfCode string
	Err        error
}

func (e *OMError) Error() string {

	return e.ErrType + ":" + e.ErrCode + " (" + e.LineOfCode + ", " + e.Path + "): " + e.ErrText
}

type ErrorResponse struct {
	ErrDetail ErrorDetailResponse `json:"ERROR"`
}

type ErrorDetailResponse struct {
	ErrCode string `json:"ERROR_CODE"`
	ErrType string `json:"ERROR_TYPE"`
	ErrText string `json:"ERROR_TEXT"`
	ErrLoc  string `json:"ERROR_LOC"`
	ErrPath string `json:"ERROR_PATH"`
}

func (err *OMError) ErrorReponsJSON() ErrorResponse {
	if err.Err != nil {
		//fmt.Printf("ERROR HANDLER: %v %v\n", err.ErrCode, ErrorResponse{ErrDetail: ErrorDetailResponse{ErrCode: err.ErrCode, ErrType: err.ErrType, ErrText: err.ErrText, ErrLoc: err.LineOfCode}})
		return ErrorResponse{ErrDetail: ErrorDetailResponse{ErrCode: err.ErrCode, ErrType: err.ErrType, ErrText: err.ErrText, ErrLoc: err.LineOfCode}}
	}
	return ErrorResponse{}
}

func NewOMError(m log.LogMessage) OMError {
	var e OMError
	e.ErrType = m.ErrorType
	e.ErrCode = m.ErrorId

	_, filename, line, _ := runtime.Caller(1)
	e.LineOfCode = fmt.Sprintf("%s:%d", filepath.Base(filename), line)
	m.LineOfCode = e.LineOfCode
	e.ErrText = m.ErrorMessage
	if m.ErrorObject != nil {
		e.ErrText = e.ErrText + ":" + m.ErrorObject.Error()
		e.Err = m.ErrorObject
	}
	//log.AppTraceLog.Error(m)
	return e
}

func NewError(eType string, eCode string, er error) OMError {
	var e OMError
	e.ErrType = eType
	e.ErrCode = eCode
	_, filename, line, _ := runtime.Caller(1)
	e.LineOfCode = fmt.Sprintf("%s:%d", filepath.Base(filename), line)
	e.Path = filepath.Dir(filename)
	if er != nil {
		e.ErrText = er.Error()
		e.Err = er
	}
	return e
}
