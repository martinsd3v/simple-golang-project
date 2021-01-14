package services

import (
	"encoding/json"
	"fmt"

	userEntity "github.com/martinsd3v/simple-golang-project/domain/users/entity"
	userRepository "github.com/martinsd3v/simple-golang-project/domain/users/repository_interface"
	infraSecurity "github.com/martinsd3v/simple-golang-project/infra/security"
	toolsErrors "github.com/martinsd3v/simple-golang-project/tools/errors"
	toolsValidations "github.com/martinsd3v/simple-golang-project/tools/validations"
)

//CreateUserTDO object receiver
type CreateUserTDO struct {
	Name     string `validate:"isRequired"`
	Email    string `validate:"isRequired|isEmail"`
	Cpf      string `validate:"isRequired|isCpf"`
	Password string `validate:"isRequired|isPassword"`
}

//CreateUserService Serviço responsável pela inserção de registros
func CreateUserService(repo userRepository.IUserRepository, dto CreateUserTDO) (created userEntity.UserEntity, err toolsErrors.ErrorDB) {

	err.Fields = toolsValidations.ValidateStruct(&dto, "", "created")

	//Check e-mail in use
	userFinderEmail, _ := repo.FindByEmail(dto.Email)
	if userFinderEmail.UUID != "" {
		inserError := toolsErrors.ErrorFields{Field: "email", Message: toolsErrors.ErrorExisting}
		err.Fields = append(err.Fields, inserError)
	}

	userFinderCpf, _ := repo.FindByCpf(dto.Cpf)
	if userFinderCpf.UUID != "" {
		inserError := toolsErrors.ErrorFields{Field: "cpf", Message: toolsErrors.ErrorExisting}
		err.Fields = append(err.Fields, inserError)
	}

	if len(err.Fields) > 0 {
		err.Code = 400
		err.Message = toolsErrors.ErrorInFields
		return
	}

	//Apply security hash in password
	hash := infraSecurity.Hash{}
	dto.Password = hash.Create(dto.Password)

	user := userEntity.UserEntity{}

	//Mergin entity and DTO
	toMerge, _ := json.Marshal(dto)
	json.Unmarshal(toMerge, &user)

	created, er := repo.Create(user)

	if er != nil {
		err.Code = 500
		err.Message = toolsErrors.ErrorCreateRegister
		fmt.Println(er)
	}

	return
}
