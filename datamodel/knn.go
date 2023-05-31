package datamodel

import (
	"math"
	"sort"

	"github.com/Parzeval8/proyecto-api-go/config"
	"github.com/Parzeval8/proyecto-api-go/models"
)

func Knn(encuesta *models.NewEncuesta) {
	// Obtener todos los registros de la tabla Encuesta
	var registros []models.Encuesta
	config.DB.Find(&registros)

	// Calcular las distancias euclidianas entre el nuevo registro y los registros existentes
	distancias := make(map[float64]int)
	for _, registro := range registros {
		distancia := math.Sqrt(math.Pow(float64(encuesta.P3Edad-registro.P3Edad), 2) +
			math.Pow(float64(encuesta.P39HabiPrac-registro.P39HabiPrac), 2) +
			math.Pow(float64(encuesta.P1Carrera-registro.P1Carrera), 2))
		distancias[distancia] = registro.P33GraMoti
	}

	// Obtener los K registros más cercanos
	k := 5 // Valor de K (número de vecinos más cercanos)
	var vecinosCercanos []float64
	for distancia := range distancias {
		vecinosCercanos = append(vecinosCercanos, distancia)
	}
	sort.Float64s(vecinosCercanos)
	vecinosCercanos = vecinosCercanos[:k]

	// Calcular el valor de P33GraMoti basado en los vecinos cercanos
	// suma := 0
	// for _, distancia := range vecinosCercanos {
	// 	suma += distancias[distancia]
	// }
	// encuesta.P33GraMoti = suma / k

	// Calcular la moda de P33GraMoti basado en los vecinos cercanos
	frecuencias := make(map[int]int)
	for _, distancia := range vecinosCercanos {
		p33GraMoti := distancias[distancia]
		frecuencias[p33GraMoti]++
	}

	moda := -1
	maxFrecuencia := -1
	for valor, frecuencia := range frecuencias {
		if frecuencia > maxFrecuencia {
			moda = valor
			maxFrecuencia = frecuencia
		}
	}

	encuesta.P33GraMoti = moda
}
