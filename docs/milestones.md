# Milestones

Glosario de términos: [docs/glosario.md](glosario.md)

## [M0]
Entrega un módulo con las siguientes capacidades:
Listar las recetas. Almacenar y entregar los ingredientes, alergias, intolerancias y dieta de un usuario.
El proceso de validación comienza por seguir las HU e intentar representarlas usando el módulo entregado. Si no es posible, el módulo está mal.
Continúa comprobando que efectivamente el M1 parte de las funcionalidades necesarias para desempeñar su función descrita sin necesidad de modificar M0. Si no es posible, el módulo está mal.

## [M1]
Entrega un módulo con las siguientes capacidades:
Dados los datos de un usuario almacenados a través del módulo de M0, devuelve las recetas potenciales para dicho usuario.
Si existe un caso en el que el orden de las recetas no respete la jerarquía de restricciones establecida, los tests deben fallar.