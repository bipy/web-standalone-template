package queries

import (
	"web-standalone-template/ent"
	"web-standalone-template/pkg/utils"
	"web-standalone-template/platform"
)

type DataSetQuery struct {
	*ent.Client
}

var (
	DataSetDB *DataSetQuery
)

func init() {
	DataSetDB = &DataSetQuery{platform.GetMySQLClient(utils.GetMySQLUrl("db1"))}
}
