package services

import (
	"encoding/json"

	userEntity "github.com/martinsd3v/simple-golang-project/domain/users/entity"
	userRepository "github.com/martinsd3v/simple-golang-project/domain/users/repository_interface"
	infraSecurity "github.com/martinsd3v/simple-golang-project/infra/security"
	toolsErrors "github.com/martinsd3v/simple-golang-project/tools/errors"
	toolsValidations "github.com/martinsd3v/simple-golang-project/tools/validations"
)

//UpdateUserTDO object receiver
type UpdateUserTDO struct {
	Name     string `validate:"isRequired"`
	Email    string `validate:"isRequired|isEmail"`
	Cpf      string `validate:"isRequired|isCpf"`
	Password string `validate:"isPassword"`
}

//UpdateUserService Serviço responsável pela inserção de registros
func UpdateUserService(repo userRepository.IUserRepository, dto UpdateUserTDO, userUUID string) (updated userEntity.UserEntity, err toolsErrors.ErrorDB) {

	err.Fields = toolsValidations.ValidateStruct(&dto, "", "updated")

	//Check e-mail in use
	userFinderEmail, _ := repo.FindByEmail(dto.Email)
	if userFinderEmail.UUID != "" && userFinderEmail.UUID != userUUID {
		inserError := toolsErrors.ErrorFields{Field: "email", Message: toolsErrors.ErrorExisting}
		err.Fields = append(err.Fields, inserError)
	}

	userFinderCpf, _ := repo.FindByCpf(dto.Cpf)
	if userFinderCpf.UUID != "" && userFinderCpf.UUID != userUUID {
		inserError := toolsErrors.ErrorFields{Field: "cpf", Message: toolsErrors.ErrorExisting}
		err.Fields = append(err.Fields, inserError)
	}

	//Check exists user with this identifier
	userFinderUUID, _ := repo.FindByUUID(userUUID)
	if userFinderUUID.UUID == "" {
		inserError := toolsErrors.ErrorFields{Field: "uuid", Message: toolsErrors.ErrorEmptyResults}
		err.Fields = append(err.Fields, inserError)
	}

	if len(err.Fields) > 0 {
		err.Code = 400
		err.Message = toolsErrors.ErrorInFields
		return
	}

	//Apply security hash in password
	if dto.Password != "" {
		hash := infraSecurity.Hash{}
		dto.Password = hash.Create(dto.Password)
	} else {
		dto.Password = userFinderCpf.Password
	}

	//Mergin entity and DTO
	toMerge, _ := json.Marshal(dto)
	json.Unmarshal(toMerge, &userFinderUUID)

	updated, er := repo.Save(userFinderUUID)

	if er != nil {
		err.Code = 500
		err.Message = toolsErrors.ErrorUpdateRegister
	}

	return
}
