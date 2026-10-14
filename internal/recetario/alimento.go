package recetario

import (
	"errors"
	"maps"
	"strings"
)

var (
	ErrAlimentoSinNombre = errors.New("el alimento debe tener nombre")
	ErrContenidoYTrazas  = errors.New("un componente no puede estar contenido y en trazas a la vez")
)

type Alimento struct {
	nombre    string
	padre     *Alimento
	alergenos map[Alergeno]bool
	contiene  map[Intolerancia]bool
	trazas    map[Intolerancia]bool
	dietas    map[Dieta]bool
}

func nuevoAlimento(nombre string, padre *Alimento, alergenos map[Alergeno]bool, contiene map[Intolerancia]bool, trazas map[Intolerancia]bool, dietas map[Dieta]bool) (Alimento, error) {
	if strings.TrimSpace(nombre) == "" {
		return Alimento{}, ErrAlimentoSinNombre
	}
	for componente := range trazas {
		if contiene[componente] {
			return Alimento{}, ErrContenidoYTrazas
		}
	}
	return Alimento{
		nombre:    nombre,
		padre:     padre,
		alergenos: maps.Clone(alergenos),
		contiene:  maps.Clone(contiene),
		trazas:    maps.Clone(trazas),
		dietas:    maps.Clone(dietas),
	}, nil
}

func (a Alimento) Igual(otro Alimento) bool {
	return a.nombre == otro.nombre &&
		a.padre == otro.padre &&
		maps.Equal(a.alergenos, otro.alergenos) &&
		maps.Equal(a.contiene, otro.contiene) &&
		maps.Equal(a.trazas, otro.trazas) &&
		maps.Equal(a.dietas, otro.dietas)
}

func (a Alimento) TieneAlergeno(alergeno Alergeno) bool {
	if a.alergenos[alergeno] {
		return true
	}
	return a.padre != nil && a.padre.TieneAlergeno(alergeno)
}
