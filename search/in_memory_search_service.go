package search

import "strings"

type InMemorySearchService struct {
	AvailableQueries []string
}

func NewInMemorySearchService(availableQueries []string) InMemorySearchService {
	return InMemorySearchService{
		AvailableQueries: availableQueries,
	}
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
