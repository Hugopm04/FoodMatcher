package recetario

import "errors"

var ErrAlimentoFueraDeTaxonomia = errors.New("el alimento no está en la taxonomía")

type Inventario struct {
	alimentos map[string]bool
}

func NuevoInventario() Inventario {
	return Inventario{alimentos: map[string]bool{}}
}

func (i Inventario) Añadir(taxonomia Taxonomia, nombre string) error {
	if _, esta := taxonomia.alimentos[nombre]; !esta {
		return ErrAlimentoFueraDeTaxonomia
	}
	i.alimentos[nombre] = true
	return nil
}
