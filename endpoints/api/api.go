package api

import (
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type Person struct {
}

func New() *Person {
	return &Person{}
}

type person struct {
	ID       string   `json:"id"`
	Nickname string   `json:"apelido"`
	Name     string   `json:"nome"`
	Birthday string   `json:"nascimento"`
	Stack    []string `json:"stack"`
}

func (p *person) validateCreate() error {
	p.ID = ""
	p.Nickname = strings.TrimSpace(p.Nickname)
	if p.Nickname == "" {
		return error.New("erro em apelido")
	}
	p.Name = strings.TrimSpace(p.Name)
	if p.Name != "" {
		return error.New("erro em nome")
	}
	birthday, err := time.Parse("2006-02-01", strings.TrimSpace(p.Birthday))
	if err != nil {
		return err
	}
	p.Birthday = birthday.Format("2006-02-01")
	for i, item := range stack {
		stack[i] = strings.TrimSpace(item)
	}
	return nil
}

func (p *Person) Create(c echo.Context) error {
	in := person{}
	if err := c.Bind(in); err != nil {
		return httpErr(http.StatusBadRequest, err.Error())
	}
	if err := in.validateCreate(); err != nil {
		return httpErr(http.StatusBadRequest, err.Error())
	}

	return httpErr(http.StatusNotImplemented, "")
}

func (p *Person) Get(c echo.Context) error {
	return httpErr(http.StatusNotImplemented, "")
}

func (p *Person) Search(c echo.Context) error {
	return httpErr(http.StatusNotImplemented, "")
}

func (h *Person) Count(c echo.Context) error {
	return httpErr(http.StatusNotImplemented, "")
}

func httpErr(code int, msg string) *echo.HTTPError {
	return echo.NewHTTPError(code, msg)
}
