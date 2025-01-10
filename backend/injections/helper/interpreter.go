package helper

import "strings"

func Interpreter(dbms string) string {
	dbms = strings.ToLower(dbms)
	switch dbms {
	case "mysql", "oracle", "mariadb":
		return "MySQL"
	case "mssql":
		return "MsSQL"
	case "postgresql":
		return "PostgresSQL"
	case "sqlite":
		return "SQLite"
	default:
		return "MySQL (default)"
	}
}
