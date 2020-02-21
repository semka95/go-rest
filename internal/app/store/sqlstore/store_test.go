package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		databaseURL = "host=localhost user=semyonz password=pass dbname=restapi_test sslmode=disable"
	}

	os.Exit(m.Run())
}
