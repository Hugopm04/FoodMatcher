package recetario

import (
	"errors"
	"slices"
	"strings"
)

type Alimento struct {
	nombre    string
	alergenos []string
}

func NuevoAlimento(nombre string, alergenos []string) (Alimento, error) {
	if strings.TrimSpace(nombre) == "" {
		return Alimento{}, errors.New("el alimento debe tener nombre")
	}
	return Alimento{nombre: nombre, alergenos: slices.Clone(alergenos)}, nil
}
