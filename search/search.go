package search

import "fmt"

type Searcher struct {
	Validator  QueryValidator
	Repository Repository
}

var searchStateLiveData = SearchState{}

func (s Searcher) Search(query string) {
	if !s.Validator.Validate(query) {
		searchStateLiveData.Error = fmt.Errorf("error: bad query")
		return
	}

	results, err := s.Repository.PerformSearch(query)
	searchStateLiveData.Result = results
	searchStateLiveData.Error = err
}

func (s Searcher) ResultState() SearchState {
	return searchStateLiveData
}
