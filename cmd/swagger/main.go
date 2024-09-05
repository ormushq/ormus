package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/doc/swagger"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var listInstances = []string{
	"manager",
	"source",
}

func main() {
	cfg := config.C().Swagger
	if !cfg.Expose {
		for {
			fmt.Println("swagger expose is disabled")
			time.Sleep(time.Hour)
		}
	}

	swagger.SwaggerInfomanager.Title = cfg.Manager.Title
	swagger.SwaggerInfomanager.Description = cfg.Manager.Description
	swagger.SwaggerInfomanager.Version = cfg.Manager.Version
	swagger.SwaggerInfomanager.Host = cfg.Manager.Host
	swagger.SwaggerInfomanager.BasePath = cfg.Manager.BasePath
	swagger.SwaggerInfomanager.Schemes = []string{"http"}

	swagger.SwaggerInfosource.Title = cfg.Source.Title
	swagger.SwaggerInfosource.Description = cfg.Source.Description
	swagger.SwaggerInfosource.Version = cfg.Source.Version
	swagger.SwaggerInfosource.Host = cfg.Source.Host
	swagger.SwaggerInfosource.BasePath = cfg.Source.BasePath
	swagger.SwaggerInfosource.Schemes = []string{"http"}

	e := echo.New()

	e.GET("/", home)

	for _, instance := range listInstances {
		e.GET(fmt.Sprintf("/%s/*", instance), echoSwagger.EchoWrapHandler(func(cfg *echoSwagger.Config) {
			cfg.InstanceName = instance
		}))
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.Port)))
}

func home(ctx echo.Context) error {
	html := `
		<html>
			<head>
				<title>Swagger API List</title>
			</head>
			<body>
				<h1>Swagger API List</h1>
				<ul>
		`

	for _, instance := range listInstances {
		html += fmt.Sprintf("<li><a href=\"/%s/index.html\">%s</a></li>", instance, instance)
	}
	html += "<ul></body></html>"

	return ctx.HTML(http.StatusOK, html)
}
