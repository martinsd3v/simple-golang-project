package repository

import (
	userEntity "github.com/martinsd3v/simple-golang-project/domain/users/entity"
	toolsPaginator "github.com/martinsd3v/simple-golang-project/tools/paginator"
)

// IUserRepository interface para padronizar
type IUserRepository interface {
	All(paginatorDto toolsPaginator.PaginatorDTO) ([]userEntity.UserEntity, toolsPaginator.PaginatorDTO, error)
	Create(userEntity.UserEntity) (userEntity.UserEntity, error)
	FindByEmail(email string) (userEntity.UserEntity, error)
	FindByCpf(cpf string) (userEntity.UserEntity, error)
	FindByUUID(string) (userEntity.UserEntity, error)
	Destroy(uuid string) error
	Save(userEntity.UserEntity) (userEntity.UserEntity, error)
}
