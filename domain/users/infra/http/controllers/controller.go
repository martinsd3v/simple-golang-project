package controllers

import "gorm.io/gorm"

//Controller responsável por receber os repositories
type Controller struct {
	Connection *gorm.DB
}
