package recetario

type Usuario struct {
	alergias      []Alergeno
	intolerancias []Intolerancia
	dieta         Dieta
	inventario    []Alimento
}
