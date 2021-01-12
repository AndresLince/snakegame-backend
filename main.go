package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"./routes"
	"github.com/jackc/pgx/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file", err.Error())
	}
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Println(err)
	}

	r := routes.Routers(pool)
	http.ListenAndServe(":3000", r)
}
