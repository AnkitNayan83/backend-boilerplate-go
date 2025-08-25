package main

import (
	"context"
	"fmt"
	"log"

	"github.com/AnkitNayan83/backend-boilerplate-go/api"
	db "github.com/AnkitNayan83/backend-boilerplate-go/db/sqlc"
	"github.com/AnkitNayan83/backend-boilerplate-go/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	conn, err := pgxpool.New(context.Background(), config.DBUrl)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer conn.Close()

	store := db.NewStore(conn)

	runGinServer(store, config)
}

func runGinServer(store db.Store, config util.Config) {
	// set gin mode
	// gin.SetMode(gin.ReleaseMode)

	server, err := api.NewServer(store, config)

	if err != nil {
		log.Fatal("Cannot create server:", err)
	}

	fmt.Printf("Starting gin http server at %s\n", config.GinHttpPort)

	err = server.Start(config.GinHttpPort)

	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
