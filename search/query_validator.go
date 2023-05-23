package search

import "strings"

type QueryValidator struct {
	MinQueryLength int
}

func (q QueryValidator) Validate(query string) bool {
	return len(strings.TrimSpace(query)) > q.MinQueryLength
}
