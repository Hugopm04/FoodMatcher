package recetario

type Receta struct {
	ingredientes map[string]struct{}
	minutos      int
}
