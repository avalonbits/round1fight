package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	_ "net/http/pprof"

	"github.com/avalonbits/round1fight/endpoints/api"
	"github.com/avalonbits/round1fight/service/person"
	"github.com/avalonbits/round1fight/storage/pg"
	"github.com/avalonbits/round1fight/storage/pg/repo"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mailru/easyjson"
)

func main() {
	uuid.EnableRandPool()

	ctx := context.Background()
	writeConn, readConn := connect(ctx)
	defer writeConn.Close(ctx)
	defer readConn.Close()

	svc := person.New(repo.New(&sConn{writeConn: writeConn, readConn: readConn}))
	person := api.New(svc)
	e := echo.New()
	e.Use(middleware.Gzip())

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

func connect(ctx context.Context) (*pgx.Conn, *pgxpool.Pool) {
	readConn, err := pgxpool.New(ctx, os.Getenv("POOL_DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	if err := readConn.Ping(ctx); err != nil {
		panic(err)
	}

	writeConn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	if _, err := writeConn.Exec(ctx, pg.Schema); err != nil {
		panic(err)
	}

	return writeConn, readConn
}

type sConn struct {
	mu        sync.Mutex
	writeConn *pgx.Conn
	readConn  *pgxpool.Pool
}

func (c *sConn) Exec(ctx context.Context, q string, args ...any) (pgconn.CommandTag, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.writeConn.Exec(ctx, q, args...)
}

func (c *sConn) Query(ctx context.Context, q string, args ...any) (pgx.Rows, error) {
	return c.readConn.Query(ctx, q, args...)
}

func (c *sConn) QueryRow(ctx context.Context, q string, args ...any) pgx.Row {
	return c.readConn.QueryRow(ctx, q, args...)
}
