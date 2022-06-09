package utils

import (
	"fmt"
	"web-standalone-template/pkg/configs"
)

func GetMySQLUrl(db string) string {
	prefix := fmt.Sprintf("%s:%s@tcp(%s)/", configs.MySQLUser, configs.MySQLPassword, configs.MySQLHost)
	switch db {
	case "db1":
		return prefix + configs.MySQLDatabase
	}
	return ""
}
