package recetario

import (
	"errors"
	"slices"
)

type Receta struct {
	ingredientes []Alimento
	minutos      int
}

func NuevaReceta(ingredientes []Alimento, minutos int) (Receta, error) {
	if len(ingredientes) == 0 {
		return Receta{}, errors.New("la receta debe tener al menos un ingrediente")
	}
	if minutos <= 0 {
		return Receta{}, errors.New("el tiempo tiene que ser un número de minutos mayor que cero")
	}
	return Receta{ingredientes: slices.Clone(ingredientes), minutos: minutos}, nil
}
