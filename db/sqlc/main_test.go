package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/AnkitNayan83/backend-boilerplate-go/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries
var testDb *pgxpool.Pool

func TestMain(m *testing.M) {

	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	fmt.Printf("DB URL: %s\n", config.DBUrl)

	testDb, err = pgxpool.New(context.Background(), config.DBUrl)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	defer testDb.Close()

	testQueries = New(testDb)

	os.Exit(m.Run())
}
