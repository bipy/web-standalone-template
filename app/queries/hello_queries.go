package queries

import (
	"context"
	"math/rand"
	"web-standalone-template/ent/mytable"
	"web-standalone-template/pkg/repository"
)

func (db *DataSetQuery) GetV(n int) ([]string, error) {
	idx := make([]int, n)
	for i := 0; i < n; i++ {
		idx[i] = 1 + rand.Intn(repository.ID)
	}
	return db.MyTable.Query().
		Where(mytable.IDIn(idx...)).
		Select(mytable.FieldName).
		Strings(context.Background())
}
