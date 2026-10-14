package recetario

import (
	"errors"
	"slices"
	"strings"
)

var (
	ErrUsuarioSinNick       = errors.New("el usuario tiene que  tener nick")
	ErrDietaInvalida        = errors.New("la dieta no es válida")
	ErrAlergiaInvalida      = errors.New("la alergia no es válida")
	ErrIntoleranciaInvalida = errors.New("la intolerancia no es válida")
)

type Usuario struct {
	nick          string
	dieta         Dieta
	alergias      []Alergeno
	intolerancias []Intolerancia
}

func NuevoUsuario(nick string, dieta Dieta, alergias []Alergeno, intolerancias []Intolerancia) (Usuario, error) {
	if strings.TrimSpace(nick) == "" {
		return Usuario{}, ErrUsuarioSinNick
	}
	usuario := Usuario{nick: nick}
	if err := usuario.CambiarRestricciones(dieta, alergias, intolerancias); err != nil {
		return Usuario{}, err
	}
	return usuario, nil
}

func (u *Usuario) CambiarRestricciones(dieta Dieta, alergias []Alergeno, intolerancias []Intolerancia) error {
	if dieta < DietaLibre || dieta > DietaVegana {
		return ErrDietaInvalida
	}
	for _, alergia := range alergias {
		if alergia < AlergenoGluten || alergia > AlergenoMoluscos {
			return ErrAlergiaInvalida
		}
	}
	for _, intolerancia := range intolerancias {
		if intolerancia < IntoleranciaGluten || intolerancia > IntoleranciaLactosa {
			return ErrIntoleranciaInvalida
		}
	}
	u.dieta = dieta
	u.alergias = slices.Clone(alergias)
	u.intolerancias = slices.Clone(intolerancias)
	return nil
}

func (u Usuario) MismoUsuario(otro Usuario) bool {
	return u.nick == otro.nick
}
