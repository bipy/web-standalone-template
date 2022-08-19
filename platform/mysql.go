package platform

import (
	"context"
	"entgo.io/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
	"web-standalone-template/ent"
)

func GetMySQLClient(url string) *ent.Client {
	client, err := ent.Open(dialect.MySQL, url)
	if err != nil {
		panic(err.Error())
	}
	err = client.Schema.Create(context.Background())
	if err != nil {
		panic(err.Error())
	}
	return client
}
