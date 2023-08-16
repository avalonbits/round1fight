package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/avalonbits/round1fight/service/person"
	"github.com/labstack/echo/v4"
	"github.com/microcosm-cc/bluemonday"
)

type Person struct {
	svc         *person.Service
	queryPolicy *bluemonday.Policy
}

func New(svc *person.Service) *Person {
	return &Person{
		svc:         svc,
		queryPolicy: bluemonday.UGCPolicy(),
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
	d := jd.Format("2006-01-02")
	fmt.Println(d)
	return []byte(d), nil
}

type personJSON struct {
	ID       string    `json:"id"`
	Nickname string    `json:"apelido"`
	Name     string    `json:"nome"`
	Birthday *jsonDate `json:"nascimento"`
	Stack    []string  `json:"stack"`
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
		return httpErr(http.StatusUnprocessableEntity, err.Error())
	}

	id, err := p.svc.Create(c.Request().Context(), in.Nickname, in.Name, in.Birthday.Time, in.Stack)
	if err != nil {
		return httpErr(http.StatusInternalServerError, err.Error())
	}

	c.Response().Header().Set("Location", "/pessoa/"+id)
	return c.String(http.StatusCreated, "")
}

func (p *Person) Get(c echo.Context) error {
	return httpErr(http.StatusNotImplemented, "")
}

func (p *Person) Search(c echo.Context) error {
	rawQuery := c.QueryParam("t")
	query := p.queryPolicy.Sanitize(strings.TrimSpace(rawQuery))
	if query == "" {
		if rawQuery != "" {
			return c.String(http.StatusOK, "[]")
		}
		return httpErr(http.StatusBadRequest, "missing query")
	}

	results, err := p.svc.Search(c.Request().Context(), query)
	if err != nil {
		return httpErr(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, results)
}

func (h *Person) Count(c echo.Context) error {
	count, err := h.svc.Count(c.Request().Context())
	if err != nil {
		return httpErr(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, fmt.Sprintf("%d", count))
}

func httpErr(code int, msg string) *echo.HTTPError {
	return echo.NewHTTPError(code, msg)
}
