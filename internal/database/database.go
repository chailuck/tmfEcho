package database

import (
	"GOKIT_v001/internal/conf"
	"GOKIT_v001/internal/global"
	"GOKIT_v001/internal/log"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	// Import the Oracle driver for sqlx
	_ "github.com/godror/godror"
	// Import sqlx and the PostgreSQL drivertest
	_ "github.com/lib/pq"
)

var DB_CONST_TERM_CURRENT_TIME string
var DB_CONST_TERM_CURRENT_DATE string
var DB_CONST_TERM_VAR_PREFIX string
var DB_CONST_TERM_RETURN_INTO string
var DB_CONST_TERM_TODTTM_FORMAT string

func NewDB() (*sqlx.DB, error) {

	var db *sqlx.DB
	var err error
	fmt.Printf("DBMS configuration: %v\n", conf.Get("db.DBMS.name"))
	if conf.Get("db.DBMS.name") == global.DBMS_NAME_POSTGRESQL {
		db, err = NewPostgresDB()
	} else {
		db, err = NewOracleDB()
	}
	if err != nil {
		log.AppTraceLog.Error(log.AppTraceLogInfo("CONNECT DB ERROR", "ADMIN", "", "", "", err.Error()))
		return nil, err
	}
	return db, nil
}

func NewPostgresDB() (*sqlx.DB, error) {

	dbmsConnStr := conf.Of("db.PostgreSQL.connect")
	log.AppTraceLog.Debug(log.AppTraceLogInfo(fmt.Sprintf("CONNECT POSTGRESQL DB: h=%v:%s u=%s db=%s ssl=%s",
		dbmsConnStr["host"], dbmsConnStr["port"], dbmsConnStr["user"], dbmsConnStr["dbname"], dbmsConnStr["sslmode"]), "ADMIN", "", "", "", ""))
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbmsConnStr["host"], dbmsConnStr["port"], dbmsConnStr["user"], dbmsConnStr["password"], dbmsConnStr["dbname"], dbmsConnStr["sslmode"])

	db, err := sqlx.Connect("postgres", dbInfo)
	if err != nil {
		log.AppTraceLog.Error(log.AppTraceLogInfo(fmt.Sprintf("ERROR CONNECT DB: h=%v:%s u=%s db=%s ssl=%s",
			dbmsConnStr["host"], dbmsConnStr["port"], dbmsConnStr["user"], dbmsConnStr["dbname"], dbmsConnStr["sslmode"]), "ADMIN", "", "", "", err.Error()))
		return nil, err
	}

	log.AppTraceLog.Debug(log.AppTraceLogInfo(fmt.Sprintf("CONNECT DB SUCCESS: h=%v:%s u=%s db=%s ssl=%s", dbmsConnStr["host"], dbmsConnStr["port"], dbmsConnStr["user"], dbmsConnStr["dbname"], dbmsConnStr["sslmode"]), "ADMIN", "", "", "", ""))
	DB_CONST_TERM_CURRENT_TIME = "now()"
	DB_CONST_TERM_CURRENT_DATE = "TODAY"
	DB_CONST_TERM_VAR_PREFIX = "$"
	DB_CONST_TERM_RETURN_INTO = " "

	DB_CONST_TERM_TODTTM_FORMAT = "'YYYY-MM-DD HH24:MI:SS'"
	return db, nil
}

func NewOracleDB() (*sqlx.DB, error) {
	// Initialize Oracle DB connection using sqlx
	dbmsConnStr := conf.Of("db.Oracle.connect")
	log.AppTraceLog.Debug(log.AppTraceLogInfo(fmt.Sprintf("CONNECT ORACLE DB: h=%v:%s u=%s db=%s sysdba=%s",
		dbmsConnStr["host"], dbmsConnStr["port"], dbmsConnStr["user"], dbmsConnStr["dbname"], dbmsConnStr["sysdba"]), "ADMIN", "", "", "", ""))

	dbInfo := fmt.Sprintf("user=\"%s\" password=\"%s\" connectString=\"%s:%s/%s\" sysdba=%s",
		//cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.DBName)
		dbmsConnStr["user"], dbmsConnStr["password"], dbmsConnStr["host"], dbmsConnStr["port"], dbmsConnStr["dbname"], dbmsConnStr["sysdba"])
	db, err := sqlx.Connect("godror", dbInfo)
	db.MapperFunc(strings.ToUpper)

	if err != nil {
		log.AppTraceLog.Error(log.AppTraceLogInfo(fmt.Sprintf("CONNECT DB: h=%v:%s u=%s db=%s ssl=%s",
			dbmsConnStr["host"], dbmsConnStr["port"], dbmsConnStr["user"], dbmsConnStr["dbname"], dbmsConnStr["sslmode"]), "ADMIN", "", "", "", err.Error()))
		return nil, err
	}
	DB_CONST_TERM_CURRENT_TIME = "CURRENT_DATE"
	DB_CONST_TERM_CURRENT_DATE = "TODAY"
	DB_CONST_TERM_VAR_PREFIX = ":"
	DB_CONST_TERM_RETURN_INTO = " INTO :1"
	DB_CONST_TERM_TODTTM_FORMAT = "'YYYY-MM-DD HH24:MI:SS'"
	return db, nil
}
