package search

import (
	"fmt"
	"strings"
)

type Repository struct {
	AvailableQueries []string
}

func NewRepository(availableQueries []string) Repository {
	return Repository{
		AvailableQueries: availableQueries,
	}
}

func (r Repository) PerformSearch(query string) ([]string, error) {
	result := []string{}

	for _, v := range r.AvailableQueries {
		if strings.Contains(
			strings.ToLower(v),
			strings.ToLower(query),
		) {
			result = append(result, v)
		}
	}

	if len(result) == 0 {
		return []string{}, fmt.Errorf("no match found for %s", query)
	}

	return result, nil
}
