package services

import (
	"testing"

	userDB "github.com/martinsd3v/simple-golang-project/domain/users/infra/data"
	infraData "github.com/martinsd3v/simple-golang-project/infra/data"

	"github.com/stretchr/testify/assert"
)

func TestCreateUserService(t *testing.T) {
	database, err := infraData.SetupTestDB("./../../../.env")
	userDB.DropTable(database)
	repo := userDB.SetupRepository(database)

	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}

	t.Run("It should be able to insert a register a user", func(t *testing.T) {
		t.Helper()

		var dto = CreateUserTDO{}
		dto.Name = "marcelo"
		dto.Cpf = "161.286.310-88"
		dto.Email = "marcelo@gmail.com"
		dto.Password = "123456"

		created, infaErrors := CreateUserService(repo, dto)

		assert.Empty(t, infaErrors.Message)           //It should not contain any error messages
		assert.Empty(t, infaErrors.Fields)            //Must not contain any fields with errors
		assert.EqualValues(t, created.Name, dto.Name) //The name must be the same as the past
		assert.NotEmpty(t, created.UUID)              //The UUID must be generated
	})

	t.Run("You should not be able to add two users with the same email or cpf", func(t *testing.T) {
		t.Helper()

		var dto = CreateUserTDO{}
		dto.Name = "marcelo"
		dto.Cpf = "161.286.310-88"
		dto.Email = "marcelo@gmail.com"
		dto.Password = "123456"

		_, infaErrors := CreateUserService(repo, dto)

		assert.NotEmpty(t, infaErrors.Message) //Must contain error message
		assert.NotEmpty(t, infaErrors.Fields)  //Must contain fields with errors
	})
}
