from dataclasses import dataclass, field

from foodmatcher.alergeno import Alergeno
from foodmatcher.dieta import Dieta
from foodmatcher.intolerancia import Intolerancia


class AlimentoInvalido(ValueError):
    """Se lanza cuando los datos de un alimento no son válidos"""


@dataclass(frozen=True)
class Alimento:
    """Alimento de la taxonomía de OPen Food Facts que es identificado por su nomnbre"""


    nombre: str
    padre: "Alimento | None" = field(default=None, compare=False)
    alergenos: frozenset[Alergeno] = field(default=frozenset(), compare=False)
    contiene: frozenset[Intolerancia] = field(default=frozenset(), compare=False)
    trazas: frozenset[Intolerancia] = field(default=frozenset(), compare=False)
    dietas: frozenset[Dieta] = field(default=frozenset(), compare=False)

    def __post_init__(self) -> None:
        if not self.nombre.strip():
            raise AlimentoInvalido("El alimento debe tener nombre")
        if self.contiene & self.trazas:
            raise AlimentoInvalido(
                f"Un componente no puede estar contenido y en trazas a la vez: {self.nombre}"
            )