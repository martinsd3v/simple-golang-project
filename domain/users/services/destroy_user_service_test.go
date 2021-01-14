package services

import (
	"testing"

	userDB "github.com/martinsd3v/simple-golang-project/domain/users/infra/data"
	infraData "github.com/martinsd3v/simple-golang-project/infra/data"

	"github.com/stretchr/testify/assert"
)

func TestDestroyUserService(t *testing.T) {
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
	created, _ := CreateUserService(repo, dto)

	for _, multTest := range []struct {
		test   string
		uuid   string
		wantOk bool
	}{
		{
			test:   "Must be able to delete user",
			uuid:   created.UUID,
			wantOk: true,
		},
		{
			test:   "Must be able to delete non-existent user",
			uuid:   "invalid",
			wantOk: false,
		},
	} {
		t.Run(multTest.test, func(t *testing.T) {
			t.Helper()
			deleted := DestroyUserService(repo, multTest.uuid)
			assert.Equal(t, deleted, multTest.wantOk)
		})
	}
}
