package paginator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	t.Run("Should be able formmating paginator parameters", func(t *testing.T) {
		t.Helper()
		dto := PaginatorDTO{
			Paginator: struct {
				Limit int
				Page  int
			}{
				Limit: 30,
				Page:  11,
			},
			Filter: []struct {
				Field               string
				Value               string
				ComparisonOperators string
				LogicalOperator     string
			}{
				{
					Field:               "Name",
					Value:               "%Marcelo%",
					ComparisonOperators: "LIKE",
				},
				{
					Field: "Email",
					Value: "email@gmail.com",
				},
			},
		}
		parameters := Paginator(dto)

		assert.EqualValues(t, "Name LIKE ? AND Email = ?", parameters.Where)
		assert.EqualValues(t, dto.Paginator.Limit, parameters.Limit)
		assert.EqualValues(t, dto.Paginator.Limit*dto.Paginator.Page-dto.Paginator.Limit, parameters.Offset)
	})

	t.Run("Should be able return defaults parameters ", func(t *testing.T) {
		dto := PaginatorDTO{}
		parameters := Paginator(dto)
		assert.EqualValues(t, 20, parameters.Limit)
		assert.EqualValues(t, 1, parameters.Page)
	})
}
