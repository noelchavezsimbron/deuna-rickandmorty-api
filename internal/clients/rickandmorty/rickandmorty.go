package rickandmorty

import (
	"time"
)

type episodesResponse struct {
	Info    info            `json:"info"`
	Results []episodeResult `json:"results"`
}

type info struct {
	Count int    `json:"count"`
	Pages int    `json:"pages"`
	Next  string `json:"next"`
	Prev  any    `json:"prev"`
}
type episodeResult struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	AirDate    string    `json:"air_date"`
	Episode    string    `json:"episode"`
	Characters []string  `json:"characters"`
	URL        string    `json:"url"`
	Created    time.Time `json:"created"`
}

func getSpanName(str string) string {
	return "clients.rickandmorty." + str
}
