package rickandmorty

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"deuna-rickandmorty-api/internal/episode"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTransport struct {
	mock.Mock
}

func (m *MockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	args := m.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}

func createMockClient(transport http.RoundTripper) *resty.Client {
	client := resty.New()
	client.SetTransport(transport)
	return client
}
func createHTTPResponse(status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Body:       io.NopCloser(strings.NewReader(body)),
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
	}
}

func TestClient_GetAllEpisodes(t *testing.T) {
	t.Run("Should return all episodes from the api", func(t *testing.T) {
		var (
			mockTransport = new(MockTransport)
			mockClient    = createMockClient(mockTransport)
			client        = NewClient(mockClient, APIConfig{BaseURL: "http://localhost:8080"})
			ctx           = context.Background()
			mockResponse  = `	{	"results": [	{"id": 1, "name": "Pilot", "air_date": "December 2, 2013", "episode": "S01E01"},	{"id": 2, "name": "Lawnmower Dog", "air_date": "December 9, 2013", "episode": "S01E02"}]}`
		)
		mockTransport.On("RoundTrip", mock.Anything).Return(createHTTPResponse(http.StatusOK, mockResponse), nil)

		result, err := client.GetAllEpisodes(ctx)
		assert.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, "Pilot", result[0].Name)
		assert.Equal(t, "Lawnmower Dog", result[1].Name)
		mockTransport.AssertExpectations(t)
	})

	t.Run("Should return error", func(t *testing.T) {
		var (
			mockTransport = new(MockTransport)
			mockClient    = createMockClient(mockTransport)
			client        = NewClient(mockClient, APIConfig{BaseURL: "http://localhost:8080"})
			ctx           = context.Background()
			expectedError = errors.New("network error")
		)

		mockTransport.On("RoundTrip", mock.Anything).Return(&http.Response{}, expectedError)

		result, err := client.GetAllEpisodes(ctx)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), expectedError.Error())
		mockTransport.AssertExpectations(t)
	})
}

func TestClient_GetSingleEpisode(t *testing.T) {
	t.Run("Should return single episode from the api", func(t *testing.T) {
		var (
			mockTransport = new(MockTransport)
			mockClient    = createMockClient(mockTransport)
			client        = NewClient(mockClient, APIConfig{BaseURL: "http://localhost:8080"})
			ctx           = context.Background()
			mockResponse  = `{"id": 1,"name": "Pilot","air_date": "December 2, 2013","episode": "S01E01"}`
		)
		mockTransport.On("RoundTrip", mock.Anything).Return(createHTTPResponse(http.StatusOK, mockResponse), nil)

		result, err := client.GetSingleEpisode(ctx, 1)

		assert.NoError(t, err)
		assert.Equal(t, "Pilot", result.Name)
		assert.Equal(t, int64(1), result.ID)
		mockTransport.AssertExpectations(t)
	})

	t.Run("Should return error", func(t *testing.T) {
		var (
			mockTransport = new(MockTransport)
			mockClient    = createMockClient(mockTransport)
			client        = NewClient(mockClient, APIConfig{BaseURL: "http://localhost:8080"})
			ctx           = context.Background()
			expectedError = errors.New("network error")
		)
		mockTransport.On("RoundTrip", mock.Anything).Return(&http.Response{}, expectedError)

		result, err := client.GetSingleEpisode(ctx, 1)

		assert.Error(t, err)
		assert.Equal(t, result, episode.Episode{})
		assert.Contains(t, err.Error(), expectedError.Error())
		mockTransport.AssertExpectations(t)
	})
}
