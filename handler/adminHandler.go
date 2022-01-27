package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mahdirezaie336/IMDB/model"
	"net/http"
)

func (h *Handler) PostMovie(c echo.Context) error {
	movie := model.Movie{}

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
	movie := model.Movie{}

	err := c.Bind(&movie)
	if err != nil {
		return c.JSON(http.StatusBadRequest, makeResponse("bad-request"))
	}

	movieID := c.Param("movieID")
	rows, err := h.db.Query(fmt.Sprintf("select id from movies where id = %s and deleted_at is null", movieID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, makeResponse("server-error"))
	}

	if !rows.Next() {
		return c.JSON(http.StatusBadRequest, makeResponse("id-not-found"))
	}

	_, err = h.db.Query(fmt.Sprintf("update movies set deleted_at = now() where id=%s", movieID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, makeResponse("server-error"))
	}

	return c.JSON(http.StatusOK, makeResponse("ok"))
}

func (h *Handler) DeleteMovie(c echo.Context) error {
	movieID := c.Param("movieID")
	rows, err := h.db.Query(fmt.Sprintf("select id from comments where id=%s and deleted_at is null", movieID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, makeResponse("server-error"))
	}

	if !rows.Next() {
		return c.JSON(http.StatusBadRequest, makeResponse("id-not-found"))
	}

	_, err = h.db.Query(fmt.Sprintf("update comments set deleted_at=now() where id=%s", movieID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, makeResponse("server-error"))
	}

	return c.JSON(http.StatusOK, makeResponse("ok"))
}

func (h *Handler) UpdateComment(c echo.Context) error {
	comment := model.Comment{}
	err := c.Bind(&comment)
	if err != nil {
		return c.JSON(http.StatusBadRequest, makeResponse("bad-request"))
	}

	commentID := c.Param("commentID")
	rows, err := h.db.Query(fmt.Sprintf("select id from comments where id = %s and deleted_at is null", commentID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, makeResponse("server-error"))
	}

	if !rows.Next() {
		return c.JSON(http.StatusBadRequest, makeResponse("bad-request"))
	}

	_, err = h.db.Query(fmt.Sprintf("update comments set approved=%v where id=%s", comment.Approved, commentID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, makeResponse("server-error"))
	}

	return c.JSON(http.StatusOK, makeResponse("ok"))
}

func (h *Handler) DeleteComment(c echo.Context) error {
	return nil
}
