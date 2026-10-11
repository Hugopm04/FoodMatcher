from dataclasses import dataclass

from foodmatcher.inventario import Inventario
from foodmatcher.tiempo import Tiempo
from foodmatcher.usuario import Usuario


@dataclass(frozen=True)
class ProcesoEleccion:
    """Datos de una búsqueda de una receta concreta"""

    usuario: Usuario
    inventario: Inventario
    limite: Tiempo
