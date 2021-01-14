package services

import (
	"testing"

	userDB "github.com/martinsd3v/simple-golang-project/domain/users/infra/data"
	infraData "github.com/martinsd3v/simple-golang-project/infra/data"

	"github.com/stretchr/testify/assert"
)

func TestUpdateUserService(t *testing.T) {
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

	dto.Name = "marcelo martins"
	dto.Cpf = "939.478.870-00"
	dto.Email = "marcelo.martins@gmail.com"
	dto.Password = "123456"
	CreateUserService(repo, dto)

	t.Run("Must be able to update user data", func(t *testing.T) {
		t.Helper()
		updateDto := UpdateUserTDO{}
		updateDto.Name = "marcelo edited"
		updateDto.Cpf = "161.286.310-88"
		updateDto.Email = "marcelo@gmail.com"

		updated, err := UpdateUserService(repo, updateDto, created.UUID)
		assert.Empty(t, err)
		assert.EqualValues(t, updated.Name, updateDto.Name)
	})

	t.Run("Must be able to update user data and password", func(t *testing.T) {
		t.Helper()
		updateDto := UpdateUserTDO{}
		updateDto.Name = "marcelo edited"
		updateDto.Cpf = "161.286.310-88"
		updateDto.Email = "marcelo@gmail.com"
		updateDto.Password = "123654"

		updated, err := UpdateUserService(repo, updateDto, created.UUID)
		assert.Empty(t, err)
		assert.EqualValues(t, updated.Name, updateDto.Name)
	})

	t.Run("You must not be able to assign email or cpf to another user already registered", func(t *testing.T) {
		t.Helper()
		updateDto := UpdateUserTDO{}
		updateDto.Name = "marcelo martins"
		updateDto.Cpf = "939.478.870-00"
		updateDto.Email = "marcelo.martins@gmail.com"

		_, err := UpdateUserService(repo, updateDto, created.UUID)
		assert.NotEmpty(t, err)
	})

	t.Run("Must not be able to update invalid user", func(t *testing.T) {
		t.Helper()
		updateDto := UpdateUserTDO{}

		_, err := UpdateUserService(repo, updateDto, "inválid")
		assert.NotEmpty(t, err)
	})
}
