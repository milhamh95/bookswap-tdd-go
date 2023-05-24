package search

import (
	"errors"
	"strings"
)

type InMemorySearchService struct {
	AvailableQueries []string
}

func NewInMemorySearchService(availableQueries []string) InMemorySearchService {
	return InMemorySearchService{
		AvailableQueries: availableQueries,
	}
}

func (i InMemorySearchService) FindMatches(query string) ([]string, error) {
	if query == "" {
		return []string{}, errors.New("bad search")
	}

	result := []string{}

	for _, v := range i.AvailableQueries {
		if strings.Contains(
			strings.ToLower(v),
			strings.ToLower(query),
		) {
			result = append(result, v)
		}
	}

	return result, nil
}
