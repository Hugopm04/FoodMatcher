package recetario

import (
	"errors"
	"slices"
	"strings"
)

var (
	ErrAlimentoSinNombre = errors.New("el alimento debe tener nombre")
	ErrContenidoYTrazas  = errors.New("un componente no puede estar contenido y en trazas a la vez")
)

type Alimento struct {
	nombre    string
	padre     string
	alergenos []Alergeno
	contiene  []Intolerancia
	trazas    []Intolerancia
	dietas    []Dieta
}

func NuevoAlimento(nombre string, padre string, alergenos []Alergeno, contiene []Intolerancia, trazas []Intolerancia, dietas []Dieta) (Alimento, error) {
	if strings.TrimSpace(nombre) == "" {
		return Alimento{}, ErrAlimentoSinNombre
	}
	for _, componente := range trazas {
		if slices.Contains(contiene, componente) {
			return Alimento{}, ErrContenidoYTrazas
		}
	}
	return Alimento{
		nombre:    nombre,
		padre:     padre,
		alergenos: slices.Clone(alergenos),
		contiene:  slices.Clone(contiene),
		trazas:    slices.Clone(trazas),
		dietas:    slices.Clone(dietas),
	}, nil
}
