package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
}

func (h *Handler) MainPage(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome!")
}

func (h *Handler) GetComments(c echo.Context) error {
	return nil
}

func (h *Handler) GetMovies(c echo.Context) error {
	return nil
}

func (h *Handler) GetAMovie(c echo.Context) error {
	return nil
}
