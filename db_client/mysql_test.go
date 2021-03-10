package db_client

import "testing"

func TestInitDatabaseClient(t *testing.T) {
	InitDatabaseClient("root:12345678@tcp(127.0.0.1:3306)/dealmaker", nil, "mysql")
}