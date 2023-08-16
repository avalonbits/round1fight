package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	_ "net/http/pprof"

	"github.com/avalonbits/round1fight/endpoints/api"
	"github.com/avalonbits/round1fight/service/person"
	"github.com/avalonbits/round1fight/storage/pg"
	"github.com/avalonbits/round1fight/storage/pg/repo"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/mailru/easyjson"
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
	e.JSONSerializer = easyJsonSerializer{}

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	e.Logger.Fatal(e.Start(":1323"))
}

type easyJsonSerializer struct {
}

func (_ easyJsonSerializer) Serialize(c echo.Context, data any, indent string) error {
	var buf []byte
	var err error

	ejs, ok := data.(easyjson.Marshaler)
	if ok {
		buf, err = easyjson.Marshal(ejs)
	} else {
		buf, err = json.Marshal(data)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	_, err = io.Copy(c.Response(), bytes.NewBuffer(buf))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

func (_ easyJsonSerializer) Deserialize(c echo.Context, data any) error {

	js, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ejs, ok := data.(easyjson.Unmarshaler)
	if ok {
		err = easyjson.Unmarshal(js, ejs)
	} else {
		err = json.Unmarshal(js, data)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
