package recetario

type Cocina struct {
	alimentos_registrados   map[string]struct{}
	recetas_registradas     map[string]struct{}
	usuarios_registrados    map[string]struct{}
	restriccion_por_usuario map[string]RestriccionUsuario
	recetas_por_nombre      map[string]Receta
	alergeno_por_alimento   map[string]map[Alergeno]struct{}
	inventario_por_usuario  map[string]map[string]struct{}
}
