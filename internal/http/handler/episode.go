package handler

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"deuna-rickandmorty-api/internal/episode"
	"deuna-rickandmorty-api/internal/tracer"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel/codes"
)

type (
	episodeGetterUseCase interface {
		GetAll(ctx context.Context) ([]episode.Episode, error)
		GetByID(ctx context.Context, ID int64) (episode.Episode, error)
		GetMultipleByIDs(ctx context.Context, IDs []int64) ([]episode.Episode, error)
	}

	Episode struct {
		episodeGetterUseCase
	}
)

func NewEpisode(episodeGetterUseCase episodeGetterUseCase) *Episode {
	return &Episode{episodeGetterUseCase: episodeGetterUseCase}
}

// HandleGetAll godoc
// @Summary      List episodes
// @Description  get all episodes
// @Tags         episodes
// @Produce      json
// @Success      200  {array}   episode.Episode
// @Failure      400  {string}  string
// @Failure      404  {string}  string
// @Failure      500  {string}  string
// @Router       /episodes [get]
func (e *Episode) HandleGetAll(c echo.Context) error {
	ctx := c.Request().Context()
	ctx, span := tracer.Start(ctx, "handler.Episode.HandleGetAll")
	defer span.End()

	episodes, err := e.episodeGetterUseCase.GetAll(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, episodes)
}

// HandleGetByID godoc
// @Summary      Get episode by id
// @Description  get episode by id
// @Tags         episodes
// @Produce      json
// @Param        id    path     int  true  "episode id"
// @Success      200  {object}   episode.Episode
// @Failure      400  {string}  string
// @Failure      404  {string}  string
// @Failure      500  {string}  string
// @Router       /episodes/{id} [get]
func (e *Episode) HandleGetByID(c echo.Context) error {
	ctx := c.Request().Context()
	ctx, span := tracer.Start(ctx, "handler.Episode.HandleGetAll")
	defer span.End()

	ID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return c.JSON(http.StatusBadRequest, errors.New("id is not a valid integer").Error())
	}

	ep, err := e.episodeGetterUseCase.GetByID(ctx, ID)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, ep)
}

// HandleGetMultipleByIDs godoc
// @Summary      Get multiple episodes by ids
// @Description  get multiple episodes by ids
// @Tags         episodes
// @Produce      json
// @Param        ids    query     string  true  "episode ids delimited with comma ,"
// @Success      200  {array}   episode.Episode
// @Failure      400  {string}  string
// @Failure      404  {string}  string
// @Failure      500  {string}  string
// @Router       /episodes/multiple [get]
func (e *Episode) HandleGetMultipleByIDs(c echo.Context) error {
	ctx := c.Request().Context()
	ctx, span := tracer.Start(ctx, "handler.Episode.HandleGetMultipleByIDs")
	defer span.End()

	idsParam := IDsParam(c.QueryParam("ids"))

	episodeIDs, err := idsParam.Values()
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	eps, err := e.episodeGetterUseCase.GetMultipleByIDs(ctx, episodeIDs)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, eps)
}
