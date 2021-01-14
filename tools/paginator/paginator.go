package paginator

import (
	"fmt"

	"github.com/asaskevich/govalidator"
)

//Paginator responsible for mount paginator
func Paginator(dto PaginatorDTO) (paginatorParams PaginatorParams) {

	if dto.Paginator.Limit == 0 {
		dto.Paginator.Limit = 20
	}

	if dto.Paginator.Page == 0 {
		dto.Paginator.Page = 1
	}

	paginatorParams.Page = dto.Paginator.Page
	paginatorParams.Limit = dto.Paginator.Limit
	paginatorParams.Offset = dto.Paginator.Limit*dto.Paginator.Page - dto.Paginator.Limit

	acceptsComparisonOperators := []string{"=", "!=", "LIKE", "IN", "ILIKE"}
	accetpsLogicalOperator := []string{"AND", "OR", "NOT"}

	for key := range dto.Filter {
		item := dto.Filter[key]
		//Cleaning input data for mitigate sql injection
		item.Field = govalidator.ReplacePattern(item.Field, "[^0-9A-Za-z_-]", "")
		item.Value = govalidator.ReplacePattern(item.Value, "[^0-9A-Za-z@_.-]", "")

		//Setting default comparison operator
		if !contains(acceptsComparisonOperators, item.ComparisonOperators) {
			item.ComparisonOperators = "="
		}

		//Setting default logical operator
		if !contains(accetpsLogicalOperator, item.LogicalOperator) {
			item.LogicalOperator = "AND"
		}

		if item.ComparisonOperators == "LIKE" || item.ComparisonOperators == "ILIKE" {
			item.Value = "%" + item.Value + "%"
		}

		if paginatorParams.Where != "" {
			paginatorParams.Where += fmt.Sprintf(" %s %s %s ?", item.LogicalOperator, item.Field, item.ComparisonOperators)
		} else {
			paginatorParams.Where += fmt.Sprintf("%s %s ?", item.Field, item.ComparisonOperators)
		}

		paginatorParams.Params = append(paginatorParams.Params, item.Value)
	}

	return
}

func contains(s []string, searchterm string) bool {
	for k := range s {
		if s[k] == searchterm {
			return true
		}
	}
	return false
}
