package recetario

import "slices"

type ProcesoEleccion struct {
	usuario    Usuario
	inventario Inventario
	limite     Tiempo
}

func NuevoProcesoEleccion(usuario Usuario, inventario Inventario, limite Tiempo) ProcesoEleccion {
	return ProcesoEleccion{usuario: usuario, inventario: slices.Clone(inventario), limite: limite}
}
