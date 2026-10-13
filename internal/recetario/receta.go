package recetario

import (
	"errors"
	"slices"
	"strings"
)

var (
	ErrRecetaSinNombre       = errors.New("la receta tiene que tener nombre")
	ErrRecetaSinIngredientes = errors.New("la receta debe tener al menos un ingrediente")
)

type Receta struct {
	nombre       string
	ingredientes []Alimento
	tiempo       Tiempo
}

func NuevaReceta(nombre string, ingredientes []Alimento, tiempo Tiempo) (Receta, error) {
	if strings.TrimSpace(nombre) == "" {
		return Receta{}, ErrRecetaSinNombre
	}
	if len(ingredientes) == 0 {
		return Receta{}, ErrRecetaSinIngredientes
	}
	return Receta{nombre: nombre, ingredientes: slices.Clone(ingredientes), tiempo: tiempo}, nil
}

func (r Receta) MismaReceta(otra Receta) bool {
	return r.nombre == otra.nombre
}
