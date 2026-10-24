package recetario

import "slices"

type Recetario struct {
	recetas []Receta
}

func NuevoRecetario(recetas []Receta) Recetario {
	return Recetario{recetas: slices.Clone(recetas)}
}
