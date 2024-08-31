package database

import "strconv"

func AddLimitOffset(sqlStmt string, limit int, offset int) string {
	if offset > 0 {
		sqlStmt += " OFFSET " + strconv.Itoa(offset)
	}
	if limit > 0 {
		sqlStmt += " LIMIT " + strconv.Itoa(limit)
	}
	return sqlStmt
}
