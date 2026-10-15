package recetario

import (
	"errors"
	"strings"
)

var ErrAlimentoSinNombre = errors.New("el alimento debe tener nombre")

type Alimento struct {
	nombre string
}

func NuevoAlimento(nombre string) (Alimento, error) {
	if strings.TrimSpace(nombre) == "" {
		return Alimento{}, ErrAlimentoSinNombre
	}
	return Alimento{nombre: nombre}, nil
}
