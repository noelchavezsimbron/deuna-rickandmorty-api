package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"deuna-rickandmorty-api/internal/episode"
	"deuna-rickandmorty-api/internal/http/handler/mocks"

	"github.com/bxcodec/faker/v4"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func fakeDataEpisodes(size int) []episode.Episode {
	data := make([]episode.Episode, size)
	for i := 0; i < size; i++ {
		var ep episode.Episode
		_ = faker.FakeData(&ep)
		data[i] = ep
	}
	return data
}

func TestEpisode_HandleGetAll(t *testing.T) {

	t.Run("should return success", func(t *testing.T) {
		var (
			e            = echo.New()
			mockUseCase  = new(mocks.EpisodeGetterUseCaseMock)
			handler      = NewEpisode(mockUseCase)
			mockEpisodes = fakeDataEpisodes(5)
			req          = httptest.NewRequest(http.MethodGet, "/episodes", nil)
			rec          = httptest.NewRecorder()
			c            = e.NewContext(req, rec)
		)

		mockUseCase.On("GetAll", mock.Anything).Return(mockEpisodes, nil)

		err := handler.HandleGetAll(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		var response []episode.Episode
		_ = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.ElementsMatch(t, mockEpisodes, response)

		mockUseCase.AssertExpectations(t)
	})

	t.Run("should return error", func(t *testing.T) {
		var (
			e           = echo.New()
			mockUseCase = new(mocks.EpisodeGetterUseCaseMock)
			handler     = NewEpisode(mockUseCase)
			req         = httptest.NewRequest(http.MethodGet, "/episodes", nil)
			rec         = httptest.NewRecorder()
			c           = e.NewContext(req, rec)
			expectedErr = errors.New("some error")
		)

		mockUseCase.On("GetAll", mock.Anything).Return(nil, expectedErr)

		err := handler.HandleGetAll(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		var responseErr string
		_ = json.Unmarshal(rec.Body.Bytes(), &responseErr)
		assert.Equal(t, expectedErr.Error(), responseErr)
		mockUseCase.AssertExpectations(t)
	})
}

func TestEpisode_HandleGetByID(t *testing.T) {
	t.Run("should return success", func(t *testing.T) {
		var (
			e           = echo.New()
			mockUseCase = new(mocks.EpisodeGetterUseCaseMock)
			handler     = NewEpisode(mockUseCase)
			episodeID   = int64(1)
			mockEpisode = episode.Episode{ID: episodeID, Name: "Episode 1"}
			req         = httptest.NewRequest(http.MethodGet, "/episodes/1", nil)
			rec         = httptest.NewRecorder()
			c           = e.NewContext(req, rec)
		)

		mockUseCase.On("GetByID", mock.Anything, episodeID).Return(mockEpisode, nil)

		c.SetParamNames("id")
		c.SetParamValues("1")

		err := handler.HandleGetByID(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		var response episode.Episode
		_ = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, mockEpisode, response)

		mockUseCase.AssertExpectations(t)
	})

	t.Run("should return error bad request", func(t *testing.T) {
		var (
			e           = echo.New()
			mockUseCase = new(mocks.EpisodeGetterUseCaseMock)
			handler     = NewEpisode(mockUseCase)
			req         = httptest.NewRequest(http.MethodGet, "/episodes/any", nil)
			rec         = httptest.NewRecorder()
			c           = e.NewContext(req, rec)
			expectedErr = errors.New("id is not a valid integer")
		)

		c.SetParamNames("id")
		c.SetParamValues("any")

		err := handler.HandleGetByID(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		var responseErr string
		_ = json.Unmarshal(rec.Body.Bytes(), &responseErr)
		assert.Equal(t, expectedErr.Error(), responseErr)

		mockUseCase.AssertExpectations(t)
	})

	t.Run("should return error internal error", func(t *testing.T) {
		var (
			e           = echo.New()
			mockUseCase = new(mocks.EpisodeGetterUseCaseMock)
			handler     = NewEpisode(mockUseCase)
			req         = httptest.NewRequest(http.MethodGet, "/episodes/1", nil)
			rec         = httptest.NewRecorder()
			c           = e.NewContext(req, rec)
			episodeID   = int64(1)
			expectedErr = errors.New("ups")
		)

		c.SetParamNames("id")
		c.SetParamValues("1")

		mockUseCase.On("GetByID", mock.Anything, episodeID).Return(episode.Episode{}, expectedErr)

		err := handler.HandleGetByID(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		var responseErr string
		_ = json.Unmarshal(rec.Body.Bytes(), &responseErr)
		assert.Equal(t, expectedErr.Error(), responseErr)

		mockUseCase.AssertExpectations(t)
	})
}
