package handler

import (
	"database/sql"
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

	rows, err := h.db.Mariadb.Query("select id from movies where id = ? and deleted_at is null", vote.MovieID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, makeResponse("server-error"))
	}

	if !rows.Next() {
		return c.JSON(http.StatusBadRequest, makeResponse("id-not-found"))
	}

	_, err = h.db.Mariadb.Query("update movies set rating = rating + ? where id = ?", vote.Rating, vote.MovieID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, makeResponse("server-error"))
	}

	return c.String(http.StatusNoContent, "")

}

func (h *Handler) Comment(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*auth.JWTClaims)
	userID := claims.UserID
	cc := new(commentCreating)
	err := c.Bind(cc)
	if err != nil {
		return c.JSON(http.StatusBadRequest, createBadRequestResp("bad request"))
	}
	db, err := sql.Open("mysql", "root:faraz@tcp(172.17.0.2:3306)/cinema")
	defer db.Close()

	rows, err := db.Query("select id from movie where id = ? and deleted_at is null", cc.MovieID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, createInternalServerErrorResp("internal server error"))
	}

	ans := rows.Next()
	if ans == false {
		return c.JSON(http.StatusBadRequest, createBadRequestResp("movie id doesn't exist"))
	}

	_, err = db.Query("insert into comment (comment, movieID, userID) values (?, ?, ?)", cc.CommentBody, cc.MovieID, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, createInternalServerErrorResp("internal server error"))
	}
	return c.String(http.StatusNoContent, "")

}
