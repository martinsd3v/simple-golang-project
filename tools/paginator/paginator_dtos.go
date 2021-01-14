package paginator

//PaginatorParams params for feth
type PaginatorParams struct {
	Page   int
	Limit  int
	Offset int
	Where  string
	Params []interface{}
}

//PaginatorDTO results for paginator
type PaginatorDTO struct {
	Paginator struct {
		Limit int
		Page  int
	}
	Filter []struct {
		Field               string
		Value               string
		ComparisonOperators string
		LogicalOperator     string
	}
}

//ShowResultsDTO responsible for standardize show registers
type ShowResultsDTO = struct {
	Results   []interface{}
	Paginator interface{}
}
