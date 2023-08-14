package api

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/avalonbits/round1fight/service/person"
	"github.com/labstack/echo/v4"
)

type Person struct {
	svc *person.Service
}

func New(svc *person.Service) *Person {
	return &Person{
		svc: svc,
	}
}

type jsonDate struct {
	time.Time
}

func (jd *jsonDate) UnmarshalJSON(b []byte) error {
	in := strings.Trim(strings.TrimSpace(string(b)), "\"\\")
	if in == "" {
		return errors.New("erro em data")
	}
	t, err := time.Parse("2006-01-02", in)
	if err != nil {
		return err
	}
	jd.Time = t
	return nil
}

func (jd *jsonDate) MarshalJSON() ([]byte, error) {
	return []byte(jd.Format("2006-01-02")), nil
}

type personJSON struct {
	ID       string   `json:"id"`
	Nickname string   `json:"apelido"`
	Name     string   `json:"nome"`
	Birthday jsonDate `json:"nascimento"`
	Stack    []string `json:"stack"`
}

func (p *personJSON) validateCreate() error {
	p.ID = ""
	p.Nickname = strings.TrimSpace(p.Nickname)
	if p.Nickname == "" {
		return errors.New("erro em apelido")
	}
	p.Name = strings.TrimSpace(p.Name)
	if p.Name == "" {
		return errors.New("erro em nome")
	}
	if p.Birthday.IsZero() {
		return errors.New("erro em nascimento")
	}
	for i, item := range p.Stack {
		p.Stack[i] = strings.TrimSpace(item)
	}
	return nil
}

func (p *Person) Create(c echo.Context) error {
	in := &personJSON{}
	if err := c.Bind(in); err != nil {
		return httpErr(http.StatusBadRequest, err.Error())
	}
	if err := in.validateCreate(); err != nil {
		return httpErr(http.StatusBadRequest, err.Error())
	}

	return nil
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
