package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {
	t.Run("Must be able to create token", func(t *testing.T) {
		newToken := NewToken()

		tokenDetails, _ := newToken.CreateToken(TokenPayload{UserUUID: "123"})
		checkToken, err := newToken.CheckToken(tokenDetails.AccessToken)
		assert.NotEmpty(t, tokenDetails.AccessToken)
		assert.NotEmpty(t, checkToken.UserUUID)
		assert.Empty(t, err)

		_, err2 := newToken.CheckToken("inválid")
		assert.NotEmpty(t, err2)

		validToken := TokenValid(tokenDetails.AccessToken)
		assert.Empty(t, validToken)

		validToken2 := TokenValid("inválid")
		assert.NotEmpty(t, validToken2)
	})
}
