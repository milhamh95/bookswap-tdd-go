package search

import "fmt"

type Repository struct {
}

func (r Repository) PerformSearch(query string) string {
	availableQueries := map[string]string{
		"item":    "Item 1",
		"another": "Another Item",
	}

	item, ok := availableQueries[query]
	if !ok {
		return fmt.Sprintf("No match found for %s", query)
	}

	return item
}
