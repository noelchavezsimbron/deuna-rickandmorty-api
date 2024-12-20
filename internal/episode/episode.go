package episode

import (
	"sort"
	"time"
)

type (
	Episode struct {
		ID         int64     `json:"id"`
		Name       string    `json:"name"`
		AirDate    string    `json:"air_date"`
		Episode    string    `json:"episode"`
		Characters []string  `json:"characters"`
		URL        string    `json:"url"`
		Created    time.Time `json:"created" faker:"-"`
	}
	Episodes []Episode
)

func (es Episodes) SortByID() {
	sort.Slice(es, func(i, j int) bool {
		return es[i].ID < es[j].ID
	})
}
