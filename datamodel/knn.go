package datamodel

import (
	"math/rand"
	"time"

	"github.com/Parzeval8/proyecto-api-go/models"
)

func Knn(encuesta *models.NewEncuesta) {
	rand.Seed(time.Now().Unix())
	encuesta.P33GraMoti = rand.Intn(11)
}
