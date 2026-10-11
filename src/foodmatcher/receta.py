from dataclasses import dataclass

from foodmatcher.alimento import Alimento
from foodmatcher.tiempo import Tiempo


class RecetaInvalida(ValueError):
    """Se lanza cuando los datos de una receta no son válidos"""


@dataclass(eq=False)
class Receta:
    """Receta que se identifica por su nombre"""

    nombre: str
    ingredientes: frozenset[Alimento]
    tiempo: Tiempo

    def __post_init(self) -> None:
        if not self.nombre.strip():
            raise RecetaInvalida("La receta debe tener nombre")
        if not self.ingredientes:
            raise RecetaInvalida(f"La receta debe tener ingredientes: {self.nombre}")

    def __eq__(self, otra: object) -> bool:
        return isinstance(otra, Receta) and self.nombre == otra.nombre

    def __hash__(self) -> int:
        return hash(self.nombre)