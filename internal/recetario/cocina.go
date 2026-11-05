package recetario

type Cocina struct {
	restriccion_por_usuario map[string]RestriccionUsuario
	recetas_por_nombre      map[string]Receta
	alergeno_por_alimento   map[string]map[Alergeno]struct{}
	inventario_por_usuario  map[string]map[string]struct{}
}
