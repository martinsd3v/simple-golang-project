package routes

import "gorm.io/gorm"

//Router responsible for receive repositories
type Router struct {
	Connection *gorm.DB
}
