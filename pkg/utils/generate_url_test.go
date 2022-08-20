package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"web-standalone-template/pkg/configs"
)

func TestGetMySQLUrl(t *testing.T) {
	configs.MySQLHost = "localhost:3306"
	configs.MySQLUser = "root"
	configs.MySQLPassword = "123"
	configs.MySQLDatabase = "test"
	assert.Equal(t, "root:123@tcp(localhost:3306)/test", GetMySQLUrl("test"))
}
