from dataclasses import dataclass

from foodmatcher.inventario import Inventario
from foodmatcher.tiempo import Tiempo
from foodmatcher.usuario import Usuario


@dataclass(frozen=True)
class ProcesoEleccion:
    usuario: Usuario
    inventario: Inventario
    limite: Tiempo
