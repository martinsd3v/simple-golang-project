package services

import (
	"testing"

	userDB "github.com/martinsd3v/simple-golang-project/domain/users/infra/data"
	infraData "github.com/martinsd3v/simple-golang-project/infra/data"

	"github.com/stretchr/testify/assert"
)

func TestAuthenticateUserService(t *testing.T) {
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
		test     string
		email    string
		password string
		wantOk   bool
	}{
		{
			test:     "Must be able to authenticate a user",
			email:    dto.Email,
			password: dto.Password,
			wantOk:   true,
		},
		{
			test:     "Must not authenticate user with invalid email",
			email:    "invalid",
			password: dto.Password,
			wantOk:   true,
		},
		{
			test:     "Must not authenticate user with nonexistent email",
			email:    "invalid@gmail.com",
			password: dto.Password,
			wantOk:   true,
		},
		{
			test:     "Must not authenticate user with invalid password",
			email:    dto.Email,
			password: "invalid",
			wantOk:   true,
		},
	} {
		t.Run(multTest.test, func(t *testing.T) {
			t.Helper()

			token, authErrors := AuthenticateUserService(repo, AuthenticateUserDTO{
				Email:    multTest.email,
				Password: multTest.password,
			})

			if multTest.wantOk {
				assert.NotEmpty(t, token.Token)
				assert.Empty(t, authErrors)
			} else {
				assert.Empty(t, token.Token)
				assert.NotEmpty(t, authErrors)
			}
		})
	}
}
