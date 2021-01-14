package data

import (
	userEntity "github.com/martinsd3v/simple-golang-project/domain/users/entity"
	userRepository "github.com/martinsd3v/simple-golang-project/domain/users/repository_interface"
	toolsPaginator "github.com/martinsd3v/simple-golang-project/tools/paginator"

	"gorm.io/gorm"
)

//Users repositorio de usuários
type Users struct {
	database *gorm.DB
}

//Efetuando assinatura do repositorio na interface
var _ userRepository.IUserRepository = &Users{}

//SetupRepository inicializando repository
func SetupRepository(database *gorm.DB) *Users {
	database.AutoMigrate(&userEntity.UserEntity{})
	return &Users{database: database}
}

//DropTable removendo tabela do banco
func DropTable(database *gorm.DB) {
	database.Migrator().DropTable(&userEntity.UserEntity{})
}

//All responsável por listar registros
func (u *Users) All(paginatorDto toolsPaginator.PaginatorDTO) ([]userEntity.UserEntity, toolsPaginator.PaginatorDTO, error) {
	paginator := toolsPaginator.Paginator(paginatorDto)
	paginatorDto.Paginator.Page = paginator.Page
	paginatorDto.Paginator.Limit = paginator.Limit

	users := []userEntity.UserEntity{}
	result := u.database.Where(paginator.Where, paginator.Params...).Limit(paginator.Limit).Offset(paginator.Offset).Find(&users)

	return users, paginatorDto, result.Error
}

//Create responsável por inserir registros
func (u *Users) Create(user userEntity.UserEntity) (userEntity.UserEntity, error) {
	result := u.database.Create(&user)
	return user, result.Error
}

//FindByCpf retorna usuário pelo cpf
func (u *Users) FindByCpf(cpf string) (userEntity.UserEntity, error) {
	user := userEntity.UserEntity{}
	result := u.database.Where("cpf = ?", cpf).First(&user)
	return user, result.Error
}

//FindByEmail retorna usuário pelo e-mail
func (u *Users) FindByEmail(email string) (userEntity.UserEntity, error) {
	user := userEntity.UserEntity{}
	result := u.database.Where("email = ?", email).First(&user)
	return user, result.Error
}

//FindByUUID retorna usuário pelo uuid
func (u *Users) FindByUUID(uuid string) (userEntity.UserEntity, error) {
	user := userEntity.UserEntity{}
	result := u.database.Where("uuid = ?", uuid).First(&user)
	return user, result.Error
}

//Save responsável por salvar registros
func (u *Users) Save(user userEntity.UserEntity) (userEntity.UserEntity, error) {
	result := u.database.Save(&user)
	return user, result.Error
}

//Destroy responsável por salvar registros
func (u *Users) Destroy(uuid string) error {
	user := userEntity.UserEntity{}
	result := u.database.Where("uuid = ?", uuid).Delete(&user)
	return result.Error
}
