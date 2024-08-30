package log

import (
	"fmt"
	"strings"
	"time"
)

const (
	APP_LOG     = 0
	APP_WARN    = 3
	APP_SUCCESS = 4
	APP_ERROR   = 5

	API_IN            = 1
	SERVICE_IN        = 2
	SERVICE_OUT       = 3
	SERVICE_OUT_ERROR = 5
	API_OUT_SUCCESS   = 4
	API_OUT_ERROR     = 5

	API_RETURN_SUCCESS        = "200"
	API_RETURN_FAIL           = "500"
	API_RETURN_BAD_REQUEST    = "400"
	API_RETURN_UNAUTHORIZED   = "401"
	API_RETURN_BUSINESS_ERROR = "520"

	API_RETURN_BAD_REQUEST_MESSAGE = "The server cannot or will not process the request due to something that is perceived to be a client error."

	MessageLanUser          = 0
	MessageCCBUser          = 1
	MessageSourceSystem     = 2
	MessageCustomerNumber   = 3
	MessageSubscriberNumber = 4
	MessageCallHeaderId     = 5
	MessageRequestIp        = 6

	EXPECTED_ERROR   = "EXPECTED"
	UNEXPECTED_ERROR = "UNEXPECTED"
	NO_ERRPR         = "NO_ERROR"
)

type LogMessage struct {
	Step                  int       `bson:"step"` //Step in log (1,2,3,4,5)
	ApiStartTime          time.Time `bson:"apiStartTime"`
	ApiStartLocalTime     string    `bson:"-"`
	LogTime               time.Time `bson:"logTime"`
	LogLocalTime          string    `bson:"-"`
	LogMessage            string    `bson:"-"`
	LanUser               string    `bson:"lanUser"`          //send from client (USERLAN)
	CCBUser               string    `bson:"ccbUser"`          //send from client (CCB USER)
	CustomerNumber        string    `bson:"customerNumber"`   //send from client
	SubscriberNumber      string    `bson:"subscriberNumber"` //send from client
	CallHeaderId          string    `bson:"callHederID"`      //send from client
	MessageId             string    `bson:"messageID"`        //genereate in server (UUID)
	SourceSystem          string    `bson:"sourceSystem"`     //send from client (Format = Source System ID + _ + pagename???
	RequestDateTime       time.Time `bson:"requestDateTime"`
	RequestDateLocalTime  string    `bson:"-"`
	ResponseDateTime      time.Time `bson:"responseDateTime"`
	ResponseDateLocalTime string    `bson:"-"`
	RequestIp             string    `bson:"requestIP"`         //send from http header (client ip)
	ApiName               string    `bson:"apiName"`           //get from API server (from route handler)
	MachineName           string    `bson:"machineName"`       // API server ip address
	ApiStatusCode         string    `bson:"apiStatusCode"`     //get from server (HTTP_CODE 200 or 201 or 500...)
	ServiceName           string    `bson:"serviceName"`       //get from server (ReadCardWS,...)
	ServiceStatusCode     string    `bson:"serviceStatusCode"` //get from server (SALT 200 or 201 or 500...)
	EndPointURL           string    `bson:"endpointUrl"`
	EndPointSystem        string    `bson:"endpointSystem"` //get from server (TUX, SHAREPOINT)
	ErrorType             string
	ErrorId               string        `bson:"errorId"`      //get from server (SALT/External system error code)
	ErrorMessage          string        `bson:"errorMessage"` //get from server (SALT/External system error LogMessage )
	ResponseTime          time.Duration `bson:"responseTime"` //calculate in server (millisecond)
	EndToEnd              time.Duration `bson:"endToEnd"`     //calculate in server (millisecond)
	ResponseTimeMsec      int           `bson:"-"`            //millisecond
	EndToEndMsec          int           `bson:"-"`            //millisecond
	LineOfCode            string        `bson:"LineOfCode"`   //filename and line of code
	ErrorObject           error         `bson:"-"`            //error object
}

func (m *LogMessage) string() string {
	m.setLocalTime()
	m.setTimeMsec()

	v := []string{
		m.LogLocalTime,
		m.LogMessage,
		m.LanUser,
		m.CCBUser,
		m.CustomerNumber,
		m.SubscriberNumber,
		m.CallHeaderId,
		m.MessageId,
		m.SourceSystem,
		m.RequestDateLocalTime,
		m.ResponseDateLocalTime,
		m.RequestIp,
		m.ApiName,
		m.EndPointSystem,
		m.ServiceName,
		m.ApiStatusCode,
		m.ServiceStatusCode,
		m.ErrorId,
		m.ErrorType,
		m.ErrorMessage,
		fmt.Sprintf("%d", m.ResponseTimeMsec),
		fmt.Sprintf("%d", m.EndToEndMsec),
		m.MachineName,
		fmt.Sprintf("%d", m.Step),
		m.LineOfCode,
	}
	joined := fmt.Sprintf("API_LOG|%s", strings.Join(v, "|"))

	return strings.Replace(joined, "\n", " ", -1)
}

func (m *LogMessage) toAppString() string {
	m.setLocalTime()
	m.setTimeMsec()
	v := []string{
		m.LogLocalTime,
		m.LogMessage,
		m.CCBUser,
		m.CustomerNumber,
		m.SubscriberNumber,
		m.MessageId,
		m.ServiceName,
		m.ErrorId,
		m.ErrorType,
		m.ErrorMessage,
		m.MachineName,
		fmt.Sprintf("%d", m.Step),
		m.LineOfCode,
	}
	joined := fmt.Sprintf("APP_LOG|%s", strings.Join(v, "|"))

	return strings.Replace(joined, "\n", " ", -1)
}

func Milliseconds(t time.Duration) int {
	return int(t / time.Millisecond)
}

/*
	func (m *LogMessage) setError(err util.OMError) {
		m.ErrorId = err.ErrCode
		m.ErrorMessage = err.ErrText
		m.ErrorType = err.ErrType
		m.ErrorObject = err.Err
	}
*/
func (m *LogMessage) setTimeMsec() {
	m.ResponseTimeMsec = Milliseconds(m.ResponseTime)
	m.EndToEndMsec = Milliseconds(m.EndToEnd)
}

func (m *LogMessage) setLocalTime() {
	m.ApiStartLocalTime = m.ApiStartTime.Local().String()
	m.LogLocalTime = m.LogTime.Local().String()
	m.RequestDateLocalTime = m.RequestDateTime.Local().String()
	m.ResponseDateLocalTime = m.ResponseDateTime.Local().String()

}

func (m *LogMessage) SetStartTime() {
	m.ApiStartTime = time.Now()
	m.ApiStartLocalTime = m.ApiStartTime.Local().String()
	m.RequestDateTime = m.ApiStartTime
	m.RequestDateLocalTime = m.RequestDateTime.Local().String()
	m.Step = API_IN
}

func (m *LogMessage) SetServiceStartTime() {
	m.ApiStartTime = time.Now()
	m.ApiStartLocalTime = m.ApiStartTime.Local().String()
	m.RequestDateTime = m.ApiStartTime
	m.Step = SERVICE_IN
	m.RequestDateLocalTime = m.RequestDateTime.Local().String()
	m.Step = SERVICE_IN
}

func (m *LogMessage) SetEndTime(apiStatus int) {
	m.ResponseDateTime = time.Now()
	m.ResponseDateLocalTime = m.ResponseDateTime.Local().String()
	m.ResponseTime = m.ResponseDateTime.Sub(m.ApiStartTime)
	m.ResponseTimeMsec = Milliseconds(m.ResponseTime)
	m.EndToEnd = m.ResponseDateTime.Sub(m.ApiStartTime)
	m.EndToEndMsec = Milliseconds(m.EndToEnd)
	fmt.Printf("api status %v\n", apiStatus)
	m.Step = apiStatus
}

/*
	func (m *LogMessage) originString() string {
		m.setLocalTime()
		m.setTimeMsec()
		buf := bytes.NewBuffer([]byte{})
		templ := template.Must(template.New("log").Parse(format))
		templ.Execute(buf, m)

		return strings.Replace(buf.String(), "\n", " ", -1)
	}


var format = `|` + //01
	`{{.LogLocalTime}}|` +
	`{{.LanUser}}|` +
	`{{.CCBUser}}|` +
	`{{.CustomerNumber}}|` +
	`{{.SubscriberNumber}}|` + //06
	`{{.CallHeaderId}}|` +
	`{{.MessageId}}|` +
	`{{.SourceSystem}}|` +
	`{{.RequestDateLocalTime}}|` +
	`{{.ResponseDateLocalTime}}|` +
	`{{.RequestIp}}|` + //11
	`{{.ApiName}}|` +
	`{{.MachineName}}|` +
	`{{.Step}}|` +
	`{{.EndPointSystem}}|` +
	`{{.ServiceName}}|` + //16
	`{{.ApiStatusCode}}|` +
	`{{.ServiceStatusCode}}|` +
	`{{.ErrorMessage}}|` +
	`{{.ResponseTimeMsec}}|` +
	`{{.EndToEndMsec}}` + //21
	`{{.LineOfCode}}`
*/
