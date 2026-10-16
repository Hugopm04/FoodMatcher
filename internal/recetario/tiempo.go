package recetario

import "errors"

type Tiempo struct {
	minutos int
}

func NuevoTiempo(minutos int) (Tiempo, error) {
	if minutos <= 0 {
		return Tiempo{}, errors.New("el tiempo tiene que ser un número de minutos mayor que cero")
	}
	return Tiempo{minutos: minutos}, nil
}
