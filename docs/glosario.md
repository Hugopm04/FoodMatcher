# Glosario de Términos

## Índice

[Inventario](#inventario)
[Proceso de Elección](#proceso-de-elección)
[Restricción](#restricción)
    - [Restricción Débil](#restricción-débil)
    - [Restricción Fuerte](#restricción-fuerte)

## Inventario

Ingredientes disponibles por el usuario.

## Proceso de Elección

Proceso por el cual, tras haber obtenido los datos del usuario, el sistema devuelve las recetas potenciales.
Dichas recetas se organizan de la siguiente forman. Las recetas que cumplan las mismas restricciones se agrupan, y se van devolviendo las recetas de cada grupo. El orden en el que se devuelven las recetas de un grupo es aleatorio. El orden en el que se devuelven los grupos es determinado por las restricciones que satisfacen atentiendo a la jerarquía existente.

## Restricción

Limitación concreta que condiciona la validez de una receta y su probabilidad de ser recomendada.

### Restricción Débil

Limitación que condiciona la probabilidad de que una receta sea recomendada o no, con el siguiente orden de importancia:
    - Apetencia (lo que elige el usuario)
    - Tiempo deseado de cocción. 

### Restricción Fuerte

Limitación que en principio invalida una receta. El único caso en el que no invalida una receta que la incumpla es cuando no haya recetas disponibles que no la incumplan. Con el siguiente orden de importancia: 
    - Inventario (excluyente)
    - Alergias (excluyente)
    - Intolerancias 
    - Dieta 
    - Límite real del tiempo de cocción.