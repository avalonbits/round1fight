package main

import (
	"github.com/avalonbits/round1fight/endpoints/api"
	"github.com/avalonbits/round1fight/service/person"
	"github.com/avalonbits/round1fight/storage/pg/repo"
	"github.com/labstack/echo/v4"
)

func main() {
	svc := person.New(repo.New(nil))
	person := api.New(svc)
	e := echo.New()
	e.POST("/pessoas", person.Create)
	e.GET("/pessoas/:pid", person.Get)
	e.GET("/pessoas", person.Search)
	e.GET("/contagem-pessoas", person.Count)
	e.Logger.Fatal(e.Start(":1323"))
}
