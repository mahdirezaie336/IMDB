package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
}

func New() Handler {
	return Handler{}
}

func (h *Handler) MainPage(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome!")
}

func (h *Handler) List(c echo.Context) error {
	return c.JSON(http.StatusOK, fmt.Sprintf("list of all movies"))
}

func (h *Handler) PostMovie(c echo.Context) error {
	return nil
}

func (h *Handler) UpdateMovie(c echo.Context) error {
	return nil
}

func (h *Handler) DeleteMovie(c echo.Context) error {
	return nil
}
