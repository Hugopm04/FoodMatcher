package recetario

import (
	"errors"
	"strings"
)

type Alimento struct {
	nombre string
}

func NuevoAlimento(nombre string) (Alimento, error) {
	if strings.TrimSpace(nombre) == "" {
		return Alimento{}, errors.New("el alimento debe tener nombre")
	}
	return Alimento{nombre: nombre}, nil
}
