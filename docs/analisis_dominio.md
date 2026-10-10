# Análisis del dominio de la HU001

Análisis de la [HU001](historias_usuario.md#hu001) para el Milestone 0.

## Lenguaje (#7)

Se usa Python 3.14.8 por los siguientes motivos:

- Permite indicar los tipos de los datos y comprobarlos con mypy antes de ejecutar  para que no se confundan conceptos distintos del modelo.
- Con las dataclasses inmutables se pueden crear objetos que no cambian una vez creados, como los objetos valor.
- Permite definir excepciones propias para rechazar datos inválidos como un tiempo menor o igual que cero.
- Se instala en el espacio de usuario con mise y la versión queda fijada en `.python-version`.
- El estilo del código se comprueba con ruff.


## Tiempo (#8)

El glorsario habla de tiempo de cocción, pero se interpreta como el tiempo total de preparación porque la HU pide que la receta no tarde más de 4 horas. Y hay recetas que llevan tiempo aunque no se cuezan. 
Se mide en minutos enteros y siempre es mayor que cero. Se trata de un objeto valor porque incluye esa validación y lo usan tanto la receta como el límite de tiempo del usuario.


## Dieta (#9)

Se tienen tres dietas, a saber, libre, vegetariana y vegana. Libre es quien puede comer de todo y se relaciona con quien no sigue una dieta específica dentro de las personas. Cada una es más restrictiva que la anterior, así que una receta vegana vale también para un vegetariano y para alguien con dieta libre.
Al ser una lista cerrada de valores se representa con un enumerado.

## Alergias (#10)

Solo se admiten los alérgenos de la lista de Open Food Facts, ya que según el README es con la que se etiquetan las recetas. De hecho una alergia fuera de esa lista no se podría comprobar.
Al ser una lista cerrada, se representan con un enumerado. Por otro lado un alimento hereda los alérgenos de su alimento padre. Por último destacar que una persona puede tener ninguna o varias alergias.
