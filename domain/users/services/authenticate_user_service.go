package services

import (
	userEntity "github.com/martinsd3v/simple-golang-project/domain/users/entity"
	userRepository "github.com/martinsd3v/simple-golang-project/domain/users/repository_interface"
	infraAuth "github.com/martinsd3v/simple-golang-project/infra/auth"
	infraSecurity "github.com/martinsd3v/simple-golang-project/infra/security"
	toolsErrors "github.com/martinsd3v/simple-golang-project/tools/errors"
	toolsValidations "github.com/martinsd3v/simple-golang-project/tools/validations"
)

//AuthenticateUserDTO object receiver
type AuthenticateUserDTO struct {
	Email    string `validate:"isRequired|isEmail"`
	Password string `validate:"isRequired|isPassword"`
}

//ResponseAuthenticateDTO response for authenticate
type ResponseAuthenticateDTO struct {
	Token   string
	Message string
}

//AuthenticateUserService service responsible for auth user
func AuthenticateUserService(repo userRepository.IUserRepository, dto AuthenticateUserDTO) (token ResponseAuthenticateDTO, err toolsErrors.ErrorDB) {
	err.Fields = toolsValidations.ValidateStruct(&dto, "", "created")

	if len(err.Fields) > 0 {
		err.Code = 400
		err.Message = toolsErrors.ErrorInFields
		return
	}

	user := userEntity.UserEntity{}
	user, _ = repo.FindByEmail(dto.Email)

	if user.UUID != "" {
		hash := infraSecurity.Hash{}
		checkPassword := hash.Compare(user.Password, dto.Password)

		if checkPassword {
			newToken := infraAuth.NewToken()
			tokenDetails, _ := newToken.CreateToken(infraAuth.TokenPayload{UserUUID: user.UUID})

			token.Token = tokenDetails.AccessToken
			token.Message = toolsErrors.SuccessInAutenticate
		} else {
			token.Message = toolsErrors.ErrorInAutenticate
		}

	} else {
		err.Code = 404
		err.Message = toolsErrors.ErrorEmptyResults
	}

	return
}
