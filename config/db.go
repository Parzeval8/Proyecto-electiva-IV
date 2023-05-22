package config

import (
	"github.com/Parzeval8/proyecto-api-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// var DSN = "apigo://apigo:postgres@localhost:5432/encuesta"
var DSN = "user=apigo password=apigo dbname=encuesta port=5432"

// Conexion con postgres mediante gorm
func Connect() {
	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// db.AutoMigrate(&models.Prueba{})
	db.AutoMigrate(&models.Encuesta{})
	DB = db
}
