package recetario

import "errors"

var ErrTiempoInvalido = errors.New("el tiempo tiene que ser un número de minutos mayor que cero")

type Tiempo struct {
	minutos int
}

func NuevoTiempo(minutos int) (Tiempo, error) {
	if minutos <= 0 {
		return Tiempo{}, ErrTiempoInvalido
	}
	return Tiempo{minutos: minutos}, nil
}
