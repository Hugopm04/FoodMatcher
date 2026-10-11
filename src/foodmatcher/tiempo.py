from dataclasses import dataclass


class TiempoInvalido(ValueError):
    """Se lanza cuando el tiempo no es un número entero de minutos mayor que cero."""


@dataclass(frozen=True)
class Tiempo:
    """Tiempo total de preparación de una receta en minutos."""

    minutos: int

    def __post_init__(self) -> None:
        if not isinstance(self.minutos, int) or self.minutos <= 0:
            raise TiempoInvalido(
                f"El tiempo tiene que ser un número entero de minutos mayor que cero: {self.minutos}"
            )
