package store_test

import "testing"

import "os"

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		databaseURL = "host=localhost user=ubn password=pass dbname=restapi_test sslmode=disable"
	}

	os.Exit(m.Run())
}
