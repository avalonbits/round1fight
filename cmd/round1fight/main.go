package main

import (
	"context"
	"os"

	"github.com/avalonbits/round1fight/endpoints/api"
	"github.com/avalonbits/round1fight/service/person"
	"github.com/avalonbits/round1fight/storage/pg"
	"github.com/avalonbits/round1fight/storage/pg/repo"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func main() {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	defer pool.Close()
	if _, err := pool.Exec(ctx, pg.Schema); err != nil {
		panic(err)
	}

	svc := person.New(repo.New(pool))
	person := api.New(svc)
	e := echo.New()
	e.POST("/pessoas", person.Create)
	e.GET("/pessoas/:id", person.Get)
	e.GET("/pessoas", person.Search)
	e.GET("/contagem-pessoas", person.Count)
	e.Logger.Fatal(e.Start(":1323"))
}
