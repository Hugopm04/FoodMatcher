# Milestones

Glosario de términos: [docs/glosario.md](glosario.md)

## [M0]
Entrega un módulo que permite al producto desarrollado en el M1:
Listar las recetas. Almacenar y entregar los ingredientes, alergias, intolerancias y dieta de un usuario.
El proceso de validación comienza por seguir las HU e intentar representarlas usando el módulo entregado. Si no es posible, el módulo está mal.
Continúa comprobando que efectivamente el M1 parte de las funcionalidades necesarias para desempeñar su función descrita sin necesidad de modificar M0. Si no es posible, el módulo está mal.

## [M1]
Entrega un programa a un usuario que le permita:
Almacenar sus datos (ingredientes disponibles, alergias, intolerancias y dieta).
Dados sus datos, devolver las recetas potenciales para dicho usuario.
El programa posee sus propios tests para verificar su funcionamiento.
Si existe un caso en el que el orden de las recetas no respete la jerarquía de restricciones establecida o entregue a un usuario una receta que no cumple una restricción fuerte habiendo otras que sí las cumplen, los tests deben fallar.
