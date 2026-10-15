package recetario

import (
	"errors"
	"slices"
)

var ErrRecetaSinIngredientes = errors.New("la receta debe tener al menos un ingrediente")

type Receta struct {
	ingredientes []Alimento
	tiempo       Tiempo
}

func NuevaReceta(ingredientes []Alimento, tiempo Tiempo) (Receta, error) {
	if len(ingredientes) == 0 {
		return Receta{}, ErrRecetaSinIngredientes
	}
	return Receta{ingredientes: slices.Clone(ingredientes), tiempo: tiempo}, nil
}
