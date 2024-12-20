package rickandmorty

import "deuna-rickandmorty-api/internal/episode"

func mapResultsToEpisodes(rs []episodeResult) []episode.Episode {
	var episodes = make([]episode.Episode, 0, len(rs))
	for _, e := range rs {
		episodes = append(episodes, episode.Episode(e))
	}
	return episodes
}

func mapResultToEpisode(er episodeResult) episode.Episode {
	return episode.Episode(er)
}
