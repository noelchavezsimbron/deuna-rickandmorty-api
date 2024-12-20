package episode

import (
	"context"

	"deuna-rickandmorty-api/internal/tracer"

	"go.opentelemetry.io/otel/codes"
)

type EpisodesRepository interface {
	GetAllEpisodes(ctx context.Context) ([]Episode, error)
	GetSingleEpisode(ctx context.Context, ID int64) (Episode, error)
}

type GetterUseCase struct {
	EpisodesRepository
}

func NewGetterUseCase(er EpisodesRepository) *GetterUseCase {
	return &GetterUseCase{EpisodesRepository: er}
}

func (guc *GetterUseCase) GetAll(ctx context.Context) ([]Episode, error) {
	ctx, span := tracer.Start(ctx, "GetterUseCase.GetAll")
	defer span.End()

	res, err := guc.EpisodesRepository.GetAllEpisodes(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	var episodes = make([]Episode, 0, len(res))
	for _, e := range res {
		episodes = append(episodes, Episode(e))
	}

	return episodes, nil
}

func (guc *GetterUseCase) GetByID(ctx context.Context, ID int64) (Episode, error) {
	ctx, span := tracer.Start(ctx, "GetterUseCase.GetByID")
	defer span.End()

	episode, err := guc.EpisodesRepository.GetSingleEpisode(ctx, ID)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return Episode{}, err
	}

	return episode, nil
}
