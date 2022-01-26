package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
}

type Cat struct {
	Name string `json:"name"`
	Type string `json:"type"`
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

func (h *Handler) Cats(c echo.Context) error {
	cat := Cat{}
	err := c.Bind(&cat)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	return c.String(http.StatusOK, fmt.Sprintf("We got your cat! %s %s", cat.Name, cat.Type))
}
