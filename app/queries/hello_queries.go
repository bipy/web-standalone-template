package queries

import (
	"github.com/jmoiron/sqlx"
	"math/rand"
	"web-standalone-template/pkg/repository"
)

const (
	getEQuery = `SELECT e FROM my_table WHERE id IN (?)`
)

func (db *DataSetQuery) GetE(n int) ([]string, error) {
	var rt []string
	idx := make([]int, n)
	for i := 0; i < n; i++ {
		idx[i] = 1 + rand.Intn(repository.ID)
	}
	query, args, err := sqlx.In(getEQuery, idx)
	if err != nil {
		return rt, err
	}

	err1 := db.Select(&rt, query, args...)
	if err1 != nil {
		return rt, err1
	}
	return rt, nil
}
