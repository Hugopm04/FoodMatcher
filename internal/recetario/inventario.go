package recetario

import "errors"

var ErrAlimentoFueraDeTaxonomia = errors.New("el alimento no está en la taxonomía")

type Inventario map[string]bool

func (i Inventario) Añadir(taxonomia Taxonomia, nombre string) error {
	if _, esta := taxonomia[nombre]; !esta {
		return ErrAlimentoFueraDeTaxonomia
	}
	i[nombre] = true
	return nil
}
