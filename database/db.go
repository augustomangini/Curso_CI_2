package database

import (
	"log"
	"os"
	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	var host = os.Getenv("HOST")
	var user = os.Getenv("USER")
	var password = os.Getenv("PASSWORD")
	var name = os.Getenv("DBNAME")
	var port = os.Getenv("DBPORT")
	stringDeConexao := "host=" + host + " user=" + user + " password=" + password + " dbname=" + name + " port=" + port + " sslmode=require"
	DB, err := gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})
}
