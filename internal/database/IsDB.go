package database

import "tmfEcho/internal/conf"

func IsDB(dbName string) bool {
	return conf.Get("db.DBMS.name") == dbName
}
