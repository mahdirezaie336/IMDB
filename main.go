package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mahdirezaie336/IMDB/handler"
)

func main() {
	var s = echo.New()
	var h = handler.New()

	fmt.Println("Starting server ...")

	s.GET("/", h.MainPage)
	err := s.Start("0.0.0.0:8080")
	if err != nil {
		fmt.Println("Error starting server.")
	}
}
