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
	return c.JSON(http.StatusOK, fmt.Sprintf("{hello: here}"))
}
