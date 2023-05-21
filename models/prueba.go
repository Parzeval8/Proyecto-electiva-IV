package models

import "gorm.io/gorm"

//se crea la estructura del json
type Prueba struct {
	gorm.Model
	Pregunta1 string `json:"pregunta1"`
	Pregunta2 string `json:"pregunta2"`
	Pregunta3 string `json:"pregunta3"`
}
