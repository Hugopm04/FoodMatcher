package recetario

type Cocina struct {
	restriccion_por_usuario map[string]RestriccionUsuario
	recetas                 map[string]Receta
	alimentos               map[string]Alimento
	inventarios             map[string][]string
}
