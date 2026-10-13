from dataclasses import dataclass

from foodmatcher.alergeno import Alergeno
from foodmatcher.dieta import Dieta
from foodmatcher.intolerancia import Intolerancia


class UsuarioInvalido(ValueError):
    pass


@dataclass(eq=False)
class Usuario:
    nick: str
    dieta: Dieta = Dieta.LIBRE
    alergias: frozenset[Alergeno] = frozenset()
    intolerancias: frozenset[Intolerancia] = frozenset()

    def __post_init__(self) -> None:
        if not self.nick.strip():
            raise UsuarioInvalido("El usuario debe tener nick")

    def __eq__(self, otro: object) -> bool:
        return isinstance(otro, Usuario) and self.nick == otro.nick

    def __hash__(self) -> int:
        return hash(self.nick)
