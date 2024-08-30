package global

const (
	NOERROR          = 0
	UNKNOWN_ERROR    = -99
	UNEXPECTED_ERROR = 5
	EXPECTED_ERROR   = 1
	TP_ERROR         = -1
	FML32_ERROR      = -2
	DB_ERROR         = -3
	OS_ERROR         = -4
	APPL_ERROR       = -5

	// define for REC_STAT
	ALL  = 1
	EFFC = 2
	EXPR = 3

	// define for READ_FLAG
	CHECK   = 1
	GETINFO = 2

	// define for values indicator
	NOTNULL = 0
	SETNULL = -1

	LOG_LEVEL_APP_TRACE = -1 // debug: -1, info = 0, Warn = 1, error = 2, DPanic = 3, Panic = 4, Fetal = 5
	LOG_LEVEL_API_TRACE = 0  // debug: -1, info = 0, Warn = 1, error = 2, DPanic = 3, Panic = 4, Fetal = 5

	CONF_DEFAULT_FILE     = "\\..\\..\\configs\\default.conf"
	CONF_LOCAL_FILE       = "\\..\\..\\configs\\local.conf"
	CONF_LOCAL_DEBUG_FILE = "\\..\\..\\configs\\local_debug.conf"

	DBMS_NAME_ORACLE     = "Oracle"
	DBMS_NAME_POSTGRESQL = "PostgreSQL"
)
