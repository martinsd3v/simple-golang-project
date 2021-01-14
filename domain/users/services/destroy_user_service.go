package services

import (
	userRepository "github.com/martinsd3v/simple-golang-project/domain/users/repository_interface"
)

//DestroyUserService responsible for deleting register
func DestroyUserService(repo userRepository.IUserRepository, userUUID string) bool {
	user, _ := repo.FindByUUID(userUUID)
	if user.UUID != "" {
		err := repo.Destroy(userUUID)
		return err == nil
	}
	return false
}
