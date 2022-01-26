package handler

import (
	"github.com/labstack/echo/v4"
)

func New() Handler {
	return Handler{}
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

func (h *Handler) UpdateComment(c echo.Context) error {
	return nil
}

func (h *Handler) DeleteComment(c echo.Context) error {
	return nil
}
