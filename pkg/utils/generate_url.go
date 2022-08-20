package utils

import (
	"fmt"
	"web-standalone-template/pkg/configs"
)

func GetMySQLUrl(db string) string {
	prefix := fmt.Sprintf("%s:%s@tcp(%s)/", configs.MySQLUser, configs.MySQLPassword, configs.MySQLHost)
	switch db {
	case "read":
		return prefix + "read"
	default:
		return prefix + configs.MySQLDatabase
	}
}
