package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mahdirezaie336/IMDB/database"
	"net/http"
)

type Handler struct {
	db database.Database
}

func New() (Handler, error) {
	db, err := database.New("root@tcp(172.17.0.2:3306)/imdb")
	if err != nil {
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
	return nil
}

func (h *Handler) GetMovies(c echo.Context) error {
	return nil
}

func (h *Handler) GetAMovie(c echo.Context) error {
	return nil
}
