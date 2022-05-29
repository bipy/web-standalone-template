package utils

import (
	"fmt"
	"web-standalone-template/pkg/configs"
)

func GetMySQLUrl() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s",
		configs.MySQLUser, configs.MySQLPassword, configs.MySQLHost, configs.MySQLDatabase)
}
