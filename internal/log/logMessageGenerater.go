package log

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
)

type LogTracing struct {
	CcbUser      string
	CustNumb     string
	SubrNumb     string
	MessageId    string
	ApiName      string
	ServiceName  string
	SourceSystem string
}

func GenAppLog(log_message string, lt LogTracing) LogMessage {
	_, filename, line, _ := runtime.Caller(1)
	m := LogMessage{
		LogMessage:       log_message,
		LineOfCode:       fmt.Sprintf("%s:%d", filepath.Base(filename), line),
		CCBUser:          lt.CcbUser,
		CustomerNumber:   lt.CustNumb,
		SubscriberNumber: lt.SubrNumb,
		MessageId:        lt.MessageId,
		ErrorId:          "",
		ErrorMessage:     "",
		ErrorType:        "",
		ErrorObject:      nil,
		ApiName:          lt.ApiName,
		ServiceName:      lt.ServiceName,
		SourceSystem:     lt.SourceSystem,
	}
	m.setLocalTime()
	return m
}
func GenApiInboundLog(log_message string, lt LogTracing) LogMessage {

	_, filename, line, _ := runtime.Caller(1)
	m := LogMessage{
		LogMessage:       log_message,
		LineOfCode:       fmt.Sprintf("%s:%d", filepath.Base(filename), line),
		CCBUser:          lt.CcbUser,
		CustomerNumber:   lt.CustNumb,
		SubscriberNumber: lt.SubrNumb,
		MessageId:        lt.MessageId,
		ErrorId:          "",
		ErrorMessage:     "",
		ErrorType:        "",
		ErrorObject:      nil,
		ApiName:          lt.ApiName,
		ServiceName:      lt.ServiceName,
		SourceSystem:     lt.SourceSystem,
	}
	m.setLocalTime()
	return m
}

func GenErrLog(log_message string, lt LogTracing, e Global_error, er error) LogMessage {

	_, filename, line, _ := runtime.Caller(1)
	m := LogMessage{
		LogMessage:       log_message,
		LineOfCode:       fmt.Sprintf("%s:%d", filepath.Base(filename), line),
		CCBUser:          lt.CcbUser,
		CustomerNumber:   lt.CustNumb,
		SubscriberNumber: lt.SubrNumb,
		MessageId:        lt.MessageId,
		ErrorId:          e.id,
		ErrorMessage:     e.message,
		ErrorType:        e.eType,
		SourceSystem:     "OM",
	}
	if er != nil {
		m.ErrorObject = er
		m.LogMessage = m.LogMessage + " [ " + er.Error() + " ]"
	} else {
		m.ErrorObject = errors.New("Error: " + e.message)
	}
	m.setLocalTime()
	return m
}

func GenLogMessage(log_message string, ccb_user string, cust_numb string, subr_numb string, mssg_id string, err_mssg string, err_type string) LogMessage {

	_, filename, line, _ := runtime.Caller(1)
	m := LogMessage{
		LogMessage:       log_message,
		LineOfCode:       fmt.Sprintf("%s:%d", filepath.Base(filename), line),
		CCBUser:          ccb_user,
		CustomerNumber:   cust_numb,
		SubscriberNumber: subr_numb,
		MessageId:        mssg_id,
		ErrorMessage:     err_mssg,
		ErrorType:        err_type,
		SourceSystem:     "OM",
	}
	m.setLocalTime()
	return m
}

func AppTraceLogInfo(log_message string, ccb_user string, cust_numb string, subr_numb string, mssg_id string, err_mssg string) LogMessage {
	_, filename, line, _ := runtime.Caller(1)
	m := LogMessage{
		LogMessage:       log_message,
		LineOfCode:       fmt.Sprintf("%s:%d", filepath.Base(filename), line),
		CCBUser:          ccb_user,
		CustomerNumber:   cust_numb,
		SubscriberNumber: subr_numb,
		MessageId:        mssg_id,
		ErrorMessage:     err_mssg,
		SourceSystem:     "OM",
	}
	m.setLocalTime()
	return m
}

func ApiTraceLogInfo(log_mssg string, api_name string, lan_user string, ccb_user string, cust_numb string, subr_numb string, call_head string, mssg_id string) LogMessage {
	_, filename, line, _ := runtime.Caller(1)
	m := LogMessage{
		LogMessage:       log_mssg,
		ApiName:          api_name,
		LineOfCode:       fmt.Sprintf("%s:%d", filepath.Base(filename), line),
		LanUser:          lan_user,
		CCBUser:          ccb_user,
		CustomerNumber:   cust_numb,
		SubscriberNumber: subr_numb,
		CallHeaderId:     call_head,
		MessageId:        mssg_id,
		SourceSystem:     "OM",
	}
	m.setLocalTime()
	return m
}
