package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/joho/godotenv"

	_ "crowdfunding-campaign-management/docs"
	"crowdfunding-campaign-management/internal/server"
	"crowdfunding-campaign-management/internal/store"
)

// @version         1.0
// @description     A REST API in Go using Gin framework.
// @host      localhost:3000
// @BasePath  /
// @title           crowdfunding-campaign-management Test API
func main() {
	ctx := context.Background()
	godotenv.Load(".env")
	url := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s", os.Getenv("DB_LOGIN"),os.Getenv("DB_PASSWORD"),os.Getenv("DB_NAME"))
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Printf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		fmt.Printf("Unable to get new pool: %v\n", err)
	}

	db := store.NewDatabase(pool)
	h := server.NewHandler(db)

	server.NewServer(h)
}
