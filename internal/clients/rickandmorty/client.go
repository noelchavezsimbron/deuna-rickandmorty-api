package rickandmorty

import (
	"context"
	"errors"
	"fmt"

	"deuna-rickandmorty-api/internal/episode"
	"deuna-rickandmorty-api/internal/tracer"

	"github.com/go-resty/resty/v2"
	"go.opentelemetry.io/otel/codes"
)

const (
	episodes = "episode"
)

type (
	APIConfig struct {
		BaseURL string
	}

	Client struct {
		rc   *resty.Client
		conf APIConfig
	}
)

func NewClient(c *resty.Client, conf APIConfig) *Client {
	return &Client{rc: c, conf: conf}
}

func (c *Client) GetAllEpisodes(ctx context.Context) ([]episode.Episode, error) {
	ctx, span := tracer.Start(ctx, getSpanName("Client.GetAllEpisodes"))
	defer span.End()

	var (
		URL          = fmt.Sprintf("%s/%s", c.conf.BaseURL, episodes)
		episodesResp episodesResponse
	)

	res, err := c.rc.R().
		SetContext(ctx).
		SetResult(&episodesResp).
		Get(URL)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	if res.IsError() {
		err := errors.New(fmt.Sprintf("and error ocurred on try to get episodes, %s", res.String()))
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	return mapResultsToEpisodes(episodesResp.Results), nil
}

func (c *Client) GetSingleEpisode(ctx context.Context, ID int64) (episode.Episode, error) {
	ctx, span := tracer.Start(ctx, getSpanName("Client.GetSingleEpisode"))
	defer span.End()

	var (
		URL          = fmt.Sprintf("%s/%s/%d", c.conf.BaseURL, episodes, ID)
		episodesResp episodeResult
	)

	res, err := c.rc.R().
		SetContext(ctx).
		SetResult(&episodesResp).
		Get(URL)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return episode.Episode{}, err
	}

	if res.IsError() {
		err := errors.New(fmt.Sprintf("and error ocurred on try to get episode by id %d, %s", ID, res.String()))
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return episode.Episode{}, err
	}

	return mapResultToEpisode(episodesResp), nil
}
