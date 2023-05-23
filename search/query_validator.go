package search

import "strings"

type QueryValidator struct {
}

func (q QueryValidator) Validate(query string) bool {
	return len(strings.TrimSpace(query)) > 3
}
