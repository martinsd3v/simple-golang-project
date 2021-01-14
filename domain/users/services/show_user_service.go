package services

import (
	userEntity "github.com/martinsd3v/simple-golang-project/domain/users/entity"
	userRepository "github.com/martinsd3v/simple-golang-project/domain/users/repository_interface"
	toolsPaginator "github.com/martinsd3v/simple-golang-project/tools/paginator"
)

//ShowUserService serviço responsável por retornar todos registros
func ShowUserService(repo userRepository.IUserRepository, paginatorDTO toolsPaginator.PaginatorDTO) (userEntity.Users, toolsPaginator.PaginatorDTO) {

	users := userEntity.Users{}
	users, paginator, _ := repo.All(paginatorDTO)

	return users, paginator
}
