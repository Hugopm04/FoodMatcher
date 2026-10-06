# Milestones

Glosario de términos: [docs/glosario.md](glosario.md)

## [M0]
Modelo de HU_001, HU_002 y HU_003 creado a través de la metodología _DDD_ sin incluir la lógica de negocio. Válido si además de seguir la metodología _DDD_ cumple que:
    - Cada _issue_ representa un problema concreto derivado de al menos una _HU_.
    - Todos los problemas contenidos en las distintas _HU_ han sido representados con issues.
    - Cada fragmento del entregable está asociado al _issue_ que resuelve o avanza a resolver.
    - Cada _commit_ referencia o cierra un único _issue_.


## [M1]
Entrega un programa a un usuario que le permita:
Almacenar sus datos (ingredientes disponibles, alergias, intolerancias y dieta).
Dados sus datos, devolver las recetas potenciales para dicho usuario.
El programa posee sus propios tests para verificar su funcionamiento.
Si existe un caso en el que el orden de las recetas no respete la jerarquía de restricciones establecida o entregue a un usuario una receta que no cumple una restricción fuerte habiendo otras que sí la cumplen, los tests deben fallar.
