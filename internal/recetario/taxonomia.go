package recetario

import "errors"

var (
	ErrAlimentoRepetido = errors.New("el alimento ya está en la taxonomía")
	ErrPadreDesconocido = errors.New("el alimento padre no está en la taxonomía")
)

type Taxonomia map[string]*Alimento

func (t Taxonomia) Registrar(nombre string, padre string, alergenos map[Alergeno]bool, contiene map[Intolerancia]bool, trazas map[Intolerancia]bool, dietas map[Dieta]bool) error {
	if _, esta := t[nombre]; esta {
		return ErrAlimentoRepetido
	}
	var alimentoPadre *Alimento
	if padre != "" {
		encontrado, esta := t[padre]
		if !esta {
			return ErrPadreDesconocido
		}
		alimentoPadre = encontrado
	}
	alimento, err := nuevoAlimento(nombre, alimentoPadre, alergenos, contiene, trazas, dietas)
	if err != nil {
		return err
	}
	t[nombre] = &alimento
	return nil
}
