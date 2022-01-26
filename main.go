package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mahdirezaie336/IMDB/handler"
)

func main() {
	var e = echo.New()
	var h = handler.New()

	adminGroup := e.Group("/admin")
	adminGroup.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))
	adminGroup.Use(middleware.BasicAuth(func(username string, password string, c echo.Context) (bool, error) {
		return username == "admin" && password == "admin", nil
	}))

	adminGroup.GET("/main", h.MainPage)

	e.GET("/", h.MainPage)
	e.GET("/list", h.List)
	e.POST("/cats", h.Cats)

	err := e.Start("0.0.0.0:8080")
	if err != nil {
		fmt.Println("Problem starting server")
	}
}
