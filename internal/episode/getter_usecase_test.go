package episode_test

import (
	"context"
	"errors"
	"testing"

	"deuna-rickandmorty-api/internal/episode"
	"deuna-rickandmorty-api/internal/episode/mocks"

	"github.com/bxcodec/faker/v4"
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

func TestGetterUseCase_GetAll(t *testing.T) {

	t.Run("should return episodes", func(t *testing.T) {
		var (
			ctx          = context.Background()
			mockRepo     = new(mocks.EpisodesRepositoryMock)
			useCase      = episode.NewGetterUseCase(mockRepo)
			episodesLen  = 10
			mockEpisodes = fakeDataEpisodes(episodesLen)
		)

		mockRepo.On("GetAllEpisodes", mock.Anything).Return(mockEpisodes, nil)

		episodes, err := useCase.GetAll(ctx)
		assert.NoError(t, err)
		assert.Len(t, episodes, episodesLen)
		assert.ElementsMatch(t, mockEpisodes, episodes)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error", func(t *testing.T) {
		var (
			ctx         = context.Background()
			mockRepo    = new(mocks.EpisodesRepositoryMock)
			useCase     = episode.NewGetterUseCase(mockRepo)
			expectedErr = errors.New("some error")
		)

		mockRepo.On("GetAllEpisodes", mock.Anything).Return(nil, expectedErr)

		episodes, err := useCase.GetAll(ctx)
		assert.NotNil(t, err)
		assert.Nil(t, episodes)
		assert.Equal(t, expectedErr, err)
		mockRepo.AssertExpectations(t)
	})

}

func TestGetterUseCase_GetByID(t *testing.T) {
	t.Run("should return episode by id", func(t *testing.T) {
		var (
			ctx         = context.Background()
			mockRepo    = new(mocks.EpisodesRepositoryMock)
			useCase     = episode.NewGetterUseCase(mockRepo)
			episodeID   = int64(1)
			mockEpisode = episode.Episode{ID: episodeID, Name: "Episode 1"}
		)

		mockRepo.On("GetSingleEpisode", mock.Anything, episodeID).Return(mockEpisode, nil)

		ep, err := useCase.GetByID(ctx, episodeID)
		assert.NoError(t, err)
		assert.Equal(t, mockEpisode, ep)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error", func(t *testing.T) {
		var (
			ctx         = context.Background()
			mockRepo    = new(mocks.EpisodesRepositoryMock)
			useCase     = episode.NewGetterUseCase(mockRepo)
			episodeID   = int64(1)
			expectedErr = errors.New("some error")
		)

		mockRepo.On("GetSingleEpisode", mock.Anything, episodeID).Return(episode.Episode{}, expectedErr)

		_, err := useCase.GetByID(ctx, episodeID)
		assert.NotNil(t, err)
		assert.Equal(t, expectedErr, err)
		mockRepo.AssertExpectations(t)
	})
}
