package http

import (
	"deuna-rickandmorty-api/internal/http/handler"

	"github.com/labstack/echo/v4"
)

type Api struct {
	episodesHAndler *handler.Episode
}

func NewApi(eh *handler.Episode) *Api {
	return &Api{
		episodesHAndler: eh,
	}
}

func (api *Api) Routes(r *echo.Group) {
	r.GET("/health", handler.HealthCheck)
	r.GET("/episodes", api.episodesHAndler.HandleGetAll)
	r.GET("/episodes/:id", api.episodesHAndler.HandleGetByID)
}
