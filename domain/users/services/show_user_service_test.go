package services

import (
	"testing"

	userDB "github.com/martinsd3v/simple-golang-project/domain/users/infra/data"
	infraData "github.com/martinsd3v/simple-golang-project/infra/data"
	toolsPaginator "github.com/martinsd3v/simple-golang-project/tools/paginator"

	"github.com/stretchr/testify/assert"
)

func TestShowUserService(t *testing.T) {
	database, err := infraData.SetupTestDB("./../../../.env")
	userDB.DropTable(database)
	repo := userDB.SetupRepository(database)

	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}

	var dto = CreateUserTDO{}
	dto.Name = "marcelo"
	dto.Cpf = "161.286.310-88"
	dto.Email = "marcelo@gmail.com"
	dto.Password = "123456"
	CreateUserService(repo, dto)

	for _, multTest := range []struct {
		test   string
		wantOk bool
	}{
		{
			test:   "Must be able to find users",
			wantOk: true,
		},
	} {
		t.Run(multTest.test, func(t *testing.T) {
			t.Helper()
			findUser, _ := ShowUserService(repo, toolsPaginator.PaginatorDTO{})

			if multTest.wantOk {
				assert.NotEmpty(t, findUser)
			}
		})
	}
}
