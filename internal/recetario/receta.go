package recetario

import (
	"errors"
	"slices"
)

type Receta struct {
	ingredientes []Alimento
	tiempo       Tiempo
}

func NuevaReceta(ingredientes []Alimento, tiempo Tiempo) (Receta, error) {
	if len(ingredientes) == 0 {
		return Receta{}, errors.New("la receta debe tener al menos un ingrediente")
	}
	return Receta{ingredientes: slices.Clone(ingredientes), tiempo: tiempo}, nil
}
