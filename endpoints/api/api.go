package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Person struct {
}

func New() *Person {
	return &Person{}
}

func (h *Person) Create(c echo.Context) error {
	return echo.NewHTTPError(http.StatusNotImplemented, "")
}

func (h *Person) Get(c echo.Context) error {
	return echo.NewHTTPError(http.StatusNotImplemented, "")
}

func (h *Person) Search(c echo.Context) error {
	return echo.NewHTTPError(http.StatusNotImplemented, "")
}

func (h *Person) Count(c echo.Context) error {
	return echo.NewHTTPError(http.StatusNotImplemented, "")
}
