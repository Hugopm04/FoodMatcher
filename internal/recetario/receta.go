package recetario

import (
	"errors"
	"slices"
	"strings"
)

var (
	ErrRecetaSinNombre       = errors.New("la receta tiene que tener nombre")
	ErrRecetaSinIngredientes = errors.New("la receta debe tener al menos un ingrediente")
	ErrRecetaSinPasos        = errors.New("la receta tiene que tener al menos un paso")
)

type Receta struct {
	nombre       string
	ingredientes []Alimento
	pasos        []string
	tiempo       Tiempo
}

func NuevaReceta(nombre string, ingredientes []Alimento, pasos []string, tiempo Tiempo) (Receta, error) {
	if strings.TrimSpace(nombre) == "" {
		return Receta{}, ErrRecetaSinNombre
	}
	if len(ingredientes) == 0 {
		return Receta{}, ErrRecetaSinIngredientes
	}
	if len(pasos) == 0 {
		return Receta{}, ErrRecetaSinPasos
	}
	return Receta{nombre: nombre, ingredientes: slices.Clone(ingredientes), pasos: slices.Clone(pasos), tiempo: tiempo}, nil
}
