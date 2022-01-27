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

	_, err = h.db.Query(fmt.Sprintf("insert into movies (name, description) values (%s, %s)", movie.Name, movie.Description))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, makeResponse("server-error"))
	}

	return c.JSON(http.StatusOK, makeResponse("ok"))
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
