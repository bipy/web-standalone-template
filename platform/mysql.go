package platform

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func GetNewMySQLConn(url string) *sqlx.DB {
	return sqlx.MustConnect("mysql", url)
}
