package data

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//SetupTestDB Configuração inicial do banco de dados em anbiente de teste
func SetupTestDB(credentialsLocation string) (*gorm.DB, error) {
	if _, err := os.Stat(credentialsLocation); !os.IsNotExist(err) {
		var err error
		err = godotenv.Load(os.ExpandEnv(credentialsLocation))
		if err != nil {
			log.Fatalf("Error getting env %v\n", err)
		}
		return localhostDB()
	}
	return ciDB()
}

//SetupDB Inicializando banco de dados
func localhostDB() (*gorm.DB, error) {

	DbHost := os.Getenv("TEST_DB_HOST")
	DbPort := os.Getenv("TEST_DB_PORT")
	DbUser := os.Getenv("TEST_DB_USER")
	DbName := os.Getenv("TEST_DB_NAME")
	DbPassword := os.Getenv("TEST_DB_PASSWORD")

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

func ciDB() (*gorm.DB, error) {
	return nil, nil
}
