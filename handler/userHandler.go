package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/mahdirezaie336/IMDB/auth"
	"github.com/mahdirezaie336/IMDB/model"
	"net/http"
)

func (h *Handler) Vote(c echo.Context) error {
	vote := model.Vote{}
	err := c.Bind(&vote)
	if err != nil {
		return c.JSON(http.StatusBadRequest, makeResponse("bad-request"))
	}

	movieId := vote.MovieID
	rating := vote.Rating
	rows, err := h.db.Mariadb.Query("select id from movies where id = ? and deleted_at is null", movieId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, makeResponse("server-error"))
	}

	if !rows.Next() {
		return c.JSON(http.StatusBadRequest, makeResponse("id-not-found"))
	}

	_, err = h.db.Mariadb.Query("update movies set rating = rating + ? where id = ?", rating, movieId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, makeResponse("server-error"))
	}

	return c.String(204, "")
}

func (h *Handler) Comment(c echo.Context) error {
	comment := model.Comment{}
	err := c.Bind(comment)
	if err != nil {
		return c.JSON(http.StatusBadRequest, makeResponse("bad-request"))
	}

	movieId := comment.MovieID
	rows, err := h.db.Mariadb.Query("select id from movies where id = ? and deleted_at is null", movieId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, makeResponse("server-error"))
	}

	if !rows.Next() {
		return c.JSON(http.StatusBadRequest, makeResponse("id-not-found"))
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*auth.JWTClaims)
	_, err = h.db.Mariadb.Query("insert into comments (comment, movieID, claims.UserId) values (?, ?, ?)",
		comment.CommentBody, comment.MovieID, claims.UserId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, makeResponse("server-error"))
	}

	return c.String(204, "")
}
