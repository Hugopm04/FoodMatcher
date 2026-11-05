package recetario

type Cocina struct {
	restriccion_por_usuario map[string]RestriccionUsuario
	recetas                 map[string]Receta
	alergeno_por_alimento   map[string]map[Alergeno]struct{}
	inventarios             map[string][]string
}
