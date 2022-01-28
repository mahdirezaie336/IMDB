package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mahdirezaie336/IMDB/auth"
	"github.com/mahdirezaie336/IMDB/handler"
)

func main() {
	var e = echo.New()
	var h, err = handler.New()

	if err != nil {
		panic(err)
	}

	defer h.Close()

	// Logger
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	// Public group
	e.GET("/comments", h.GetComments)
	e.GET("/movies", h.GetMovies)
	e.GET("/movie/:movieID", h.GetAMovie)
	e.GET("/", h.MainPage)

	// Users group
	userGroup := e.Group("/user")
	userGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:        &auth.JWTClaims{},
		SigningMethod: "HS256",
		SigningKey:    []byte("userSecret"),
	}))

	userGroup.POST("/vote", h.Vote)
	userGroup.POST("/comment", h.Comment)

	// Admin group
	adminGroup := e.Group("/admin")
	adminGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:        &auth.JWTClaims{},
		SigningMethod: "HS256",
		SigningKey:    []byte("adminSecret"),
	}))

	adminGroup.POST("/movie", h.PostMovie)
	adminGroup.PUT("/movie/:movieID", h.UpdateMovie)
	adminGroup.DELETE("/movie/:movieID", h.DeleteMovie)
	adminGroup.PUT("/comment/:commentID", h.UpdateComment)
	adminGroup.DELETE("/comment/:commentID", h.DeleteComment)

	err = e.Start("0.0.0.0:8080")
	if err != nil {
		fmt.Println("Problem starting server")
	}
}
