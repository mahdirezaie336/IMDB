package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mahdirezaie336/IMDB/model"
	"net/http"
)

func (h *Handler) PostMovie(c echo.Context) error {
	movie := new(model.Movie)
	err := c.Bind(&movie)
	if err != nil {
		return c.JSON(http.StatusBadRequest, makeResponse("bad-request"))
	}
	h.db.Query(fmt.Sprintf("insert into movies (name, description) values ()"))
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
