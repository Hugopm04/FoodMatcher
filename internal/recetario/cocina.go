package recetario

type Cocina struct {
	usuarios    map[string]Usuario
	recetas     map[string]Receta
	alimentos   map[string]Alimento
	inventarios map[string][]string
}
