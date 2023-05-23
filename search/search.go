package search

type Searcher struct {
	Validator  QueryValidator
	Repository Repository
}

var searchStateLiveData = ""

func (s Searcher) Search(query string) {
	if !s.Validator.Validate(query) {
		searchStateLiveData = "Error: bad query"
		return
	}

	item := s.Repository.PerformSearch(query)
	searchStateLiveData = item
}

func (s Searcher) ResultState() SearchState {
	return SearchState{
		Result: searchStateLiveData,
	}
}
