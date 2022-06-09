package queries

import (
	"github.com/jmoiron/sqlx"
	"web-standalone-template/pkg/utils"
	"web-standalone-template/platform"
)

type DataSetQuery struct {
	*sqlx.DB
}

var (
	DataSetDB *DataSetQuery
)

func init() {
	DataSetDB = &DataSetQuery{DB: platform.GetNewMySQLConn(utils.GetMySQLUrl("db1"))}
}
