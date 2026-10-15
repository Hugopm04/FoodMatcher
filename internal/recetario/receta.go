package recetario

import (
	"errors"
	"maps"
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
	ingredientes map[*Alimento]bool
	pasos        []string
	tiempo       Tiempo
}

func NuevaReceta(nombre string, ingredientes map[*Alimento]bool, pasos []string, tiempo Tiempo) (Receta, error) {
	if strings.TrimSpace(nombre) == "" {
		return Receta{}, ErrRecetaSinNombre
	}
	if len(ingredientes) == 0 {
		return Receta{}, ErrRecetaSinIngredientes
	}
	if len(pasos) == 0 {
		return Receta{}, ErrRecetaSinPasos
	}
	return Receta{nombre: nombre, ingredientes: maps.Clone(ingredientes), pasos: slices.Clone(pasos), tiempo: tiempo}, nil
}

func (r Receta) Igual(otra Receta) bool {
	return maps.Equal(r.ingredientes, otra.ingredientes) &&
		slices.Equal(r.pasos, otra.pasos) &&
		r.tiempo == otra.tiempo
}
