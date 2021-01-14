package entity

import (
	"time"

	"github.com/twinj/uuid"
	"gorm.io/gorm"
)

//UserEntity is the struct default to user
type UserEntity struct {
	UUID      string    `gorm:"type:uuid;primaryKey;"`
	Name      string    `gorm:"notNull;" validate:"isRequired"`
	Email     string    `gorm:"notNull;unique;" validate:"isRequired|isEmail"`
	Cpf       string    `gorm:"notNull;unique;" validate:"isRequired|isCpf"`
	Password  string    `gorm:"notNull;" validate:"isRequired|isPassword"`
	CreatedAt time.Time `gorm:"notNull;"`
}

// BeforeCreate execultado algumas coisas antes de criar
func (u *UserEntity) BeforeCreate(db *gorm.DB) (err error) {
	u.UUID = uuid.NewV4().String()
	u.CreatedAt = time.Now()
	return
}

//PublicUser Formato limpo de retorno
type PublicUser struct {
	Name  string
	Email string
	UUID  string
}

//PublicUser result clean we dont expose any date for front-end
func (u *UserEntity) PublicUser() interface{} {
	return &PublicUser{
		Name:  u.Name,
		Email: u.Email,
		UUID:  u.UUID,
	}
}

//Users list of registers
type Users []UserEntity

//PublicUsers results clean
func (users Users) PublicUsers() []interface{} {

	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.PublicUser()
	}
	return result
}
