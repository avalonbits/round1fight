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
	"github.com/labstack/echo/v4"
	"github.com/mailru/easyjson"
)

type sConn struct {
	mu   sync.Mutex
	conn *pgx.Conn
}

func (c *sConn) Exec(ctx context.Context, q string, args ...any) (pgconn.CommandTag, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.conn.Exec(ctx, q, args...)
}

func (c *sConn) Query(ctx context.Context, q string, args ...any) (pgx.Rows, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.conn.Query(ctx, q, args...)
}

func (c *sConn) QueryRow(ctx context.Context, q string, args ...any) pgx.Row {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.conn.QueryRow(ctx, q, args...)
}

func main() {
	uuid.EnableRandPool()
	ctx := context.Background()
	dbURL := os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)
	if _, err := conn.Exec(ctx, pg.Schema); err != nil {
		panic(err)
	}

	svc := person.New(repo.New(&sConn{conn: conn}))
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
