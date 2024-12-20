package handler

import (
	"errors"
	"strconv"
	"strings"
)

type IDsParam string

func (idsp IDsParam) Values() ([]int64, error) {
	if idsp == "" {
		return nil, errors.New("ids parameter is required")
	}

	var (
		IdsStr = strings.Split(string(idsp), ",")
		IDs    []int64
	)

	for _, idStr := range IdsStr {
		ID, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return nil, errors.New("invalid id format, must be a integer")
		}
		IDs = append(IDs, ID)
	}

	return IDs, nil
}
