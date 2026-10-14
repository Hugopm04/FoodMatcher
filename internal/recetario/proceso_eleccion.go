package recetario

import "maps"

type ProcesoEleccion struct {
	usuario    *Usuario
	inventario Inventario
	limite     Tiempo
}

func NuevoProcesoEleccion(usuario *Usuario, inventario Inventario, limite Tiempo) ProcesoEleccion {
	return ProcesoEleccion{usuario: usuario, inventario: maps.Clone(inventario), limite: limite}
}
