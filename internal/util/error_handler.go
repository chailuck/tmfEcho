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
	ErrMessage string
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
	ErrType   string `json:"ERROR_TYPE"`
	ErrCode   string `json:"ERROR_CODE"`
	ErrText   string `json:"ERROR_TEXT"`
	ErrDetail string `json:"ERROR_DETAIL"`
	ErrLoc    string `json:"ERROR_LOC"`
	ErrPath   string `json:"ERROR_PATH"`
}

type TMFErrorResponse struct {
	ErrType      string `json:"@type"`
	ErrBase      string `json:"@baseType,omitempty"`
	ErrSchema    string `json:"@schemaLocation,omitempty"`
	ErrCode      string `json:"code"`
	ErrReason    string `json:"reason,omitempty"`
	ErrMessage   string `json:"message,omitempty"`
	ErrStatus    string `json:"status,omitempty"`
	ErrReference string `json:"referenceError,omitempty"`
}

/*
	"@type": "Error",
	"@baseType": "string",
	"@schemaLocation": "string",
	"code": "string",
	"reason": "string",
	"message": "string",
	"status": "string",
	"referenceError": "string"
*/

func (err *OMError) ErrorReponsJSON() ErrorResponse {
	if err.Err != nil {
		//fmt.Printf("ERROR HANDLER: %v %v\n", err.ErrCode, ErrorResponse{ErrDetail: ErrorDetailResponse{ErrCode: err.ErrCode, ErrType: err.ErrType, ErrText: err.ErrText, ErrLoc: err.LineOfCode}})
		return ErrorResponse{ErrDetail: ErrorDetailResponse{ErrCode: err.ErrCode, ErrType: err.ErrType, ErrText: err.ErrText, ErrDetail: err.ErrMessage, ErrLoc: err.LineOfCode}}
	}
	return ErrorResponse{}
}

func (err *OMError) ErrorReponsTMFJSON() TMFErrorResponse {
	if err.Err != nil {
		//fmt.Printf("ERROR HANDLER: %v %v\n", err.ErrCode, ErrorResponse{ErrDetail: ErrorDetailResponse{ErrCode: err.ErrCode, ErrType: err.ErrType, ErrText: err.ErrText, ErrLoc: err.LineOfCode}})
		return TMFErrorResponse{ErrCode: err.ErrCode, ErrType: "ERROR", ErrReason: err.ErrText, ErrMessage: err.ErrMessage, ErrStatus: err.ErrType, ErrReference: err.LineOfCode + "(" + err.Path + ")"}
	}
	return TMFErrorResponse{}
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
		e.ErrMessage = m.LogMessage
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
		e.ErrMessage = er.Error()
		e.Err = er
	}
	return e
}
