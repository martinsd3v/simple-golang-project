package services

import (
	userEntity "github.com/martinsd3v/simple-golang-project/domain/users/entity"
	userRepository "github.com/martinsd3v/simple-golang-project/domain/users/repository_interface"
)

//IndexUserService service responsible for find one register
func IndexUserService(repo userRepository.IUserRepository, userUUID string) userEntity.UserEntity {
	user := userEntity.UserEntity{}
	user, _ = repo.FindByUUID(userUUID)
	return user
}
