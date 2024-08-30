package database

import "GOKIT_v001/internal/conf"

func IsDB(dbName string) bool {
	return conf.Get("db.DBMS.name") == dbName
}
