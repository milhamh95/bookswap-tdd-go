package search

import (
	"fmt"
	"strings"
)

type Repository struct {
	AvailableQueries      []string
	InMemorySearchService InMemorySearchService
}

func NewRepository(availableQueries []string) Repository {
	return Repository{
		AvailableQueries:      availableQueries,
		InMemorySearchService: InMemorySearchService{AvailableQueries: availableQueries},
	}
}

func (r Repository) PerformSearch(query string) ([]string, error) {
	result := r.InMemorySearchService.FindMatches(query)

	if len(result) == 0 {
		return []string{}, fmt.Errorf("no match found for %s", query)
	}

	return result, nil
}

type InMemorySearchService struct {
	AvailableQueries []string
}

func (i InMemorySearchService) FindMatches(query string) []string {
	result := []string{}

	for _, v := range i.AvailableQueries {
		if strings.Contains(
			strings.ToLower(v),
			strings.ToLower(query),
		) {
			result = append(result, v)
		}
	}

	return result
}
