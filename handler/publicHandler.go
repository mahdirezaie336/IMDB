package handler

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mahdirezaie336/IMDB/databases"
	"github.com/mahdirezaie336/IMDB/model"
	"net/http"
)

type Handler struct {
	db databases.Database
}

func New() (Handler, error) {
	db, err := databases.New("root:toor@tcp(172.17.0.2:3306)/imdb")
	if err != nil {
		fmt.Println(err)
		return Handler{}, err
	}
	return Handler{
		db: db,
	}, nil
}

func (h *Handler) Close() error {
	return h.db.Close()
}

func (h *Handler) MainPage(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome!")
}

func (h *Handler) GetComments(c echo.Context) error {
	movieId := c.QueryParam("movie")
	rows, err := h.db.Mariadb.Query("select id, name, description, rating from movies where id = ? and "+
		"deleted_at is null", movieId)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, makeResponse("server-error"))
	}

	if !rows.Next() {
		return c.JSON(http.StatusBadRequest, makeResponse("id-not-found"))
	}

	movie := model.Movie{}
	err = rows.Scan(&(movie.Id), &(movie.Name), &(movie.Description), &(movie.Rating))
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, makeResponse("server-error"))
	}

	allComments := make([]model.Comment, 0)
	rows, err = h.db.Mariadb.Query("select c.id, c.comment, u.username from comments as c join users as u on "+
		"c.userID = u.id where movieID = ? and c.approved is true", movieId)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, makeResponse("server-error"))
	}
	for rows.Next() {
		comment := model.Comment{}
		err := rows.Scan(&(comment.Id), &(comment.CommentBody), &(comment.User))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, makeResponse("server-error"))
		}
		allComments = append(allComments, comment)
	}

	marshalled, err := json.Marshal(allComments)
	return c.JSON(http.StatusOK, map[string]string{
		"movie":    "" + movieId,
		"comments": string(marshalled),
	})
}

func (h *Handler) GetMovies(c echo.Context) error {
	allMovies := make([]model.Movie, 0)
	rows, err := h.db.Mariadb.Query("select id, name, description, rating from movies where deleted_at is null")
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, makeResponse("server-error"))
	}

	for rows.Next() {
		movie := model.Movie{}
		err := rows.Scan(&(movie.Id), &(movie.Name), &(movie.Description), &(movie.Rating))
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusInternalServerError, makeResponse("server-error"))
		}
		allMovies = append(allMovies, movie)
	}

	marshalled, err := json.Marshal(allMovies)
	return c.JSON(http.StatusOK, map[string]string{
		"movies": string(marshalled),
	})
}

func (h *Handler) GetAMovie(c echo.Context) error {
	return nil
}
