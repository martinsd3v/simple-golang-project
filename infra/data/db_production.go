package data

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//SetupDB Configuração inicial do banco de dados
func SetupDB(credentialsLocation string) (*gorm.DB, error) {
	if _, err := os.Stat(credentialsLocation); !os.IsNotExist(err) {
		var err error
		err = godotenv.Load(os.ExpandEnv(credentialsLocation))
		if err != nil {
			log.Fatalf("Error getting env %v\n", err)
		}
		return productionDB()
	}
	return nil, errors.New("Credentials .env not found")
}

//SetupDB Inicializando banco de dados
func productionDB() (*gorm.DB, error) {

	DbHost := os.Getenv("DB_HOST")
	DbPort := os.Getenv("DB_PORT")
	DbUser := os.Getenv("DB_USER")
	DbName := os.Getenv("DB_NAME")
	DbPassword := os.Getenv("DB_PASSWORD")

	DBURL := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		DbHost,
		DbPort,
		DbUser,
		DbName,
		DbPassword)

	db, err := gorm.Open(
		postgres.Open(DBURL),
		&gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
