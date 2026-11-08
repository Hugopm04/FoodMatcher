package recetario

type RestriccionUsuario struct {
	alergias      map[Alergeno]struct{}
	intolerancias map[Intolerancia]struct{}
	dieta         Dieta
}
