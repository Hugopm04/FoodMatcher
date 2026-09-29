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
Proceso por el cuál, tras haber obtenido los datos del usuario, el sistema devuelve las recetas potenciales.

## Restricción

Limitación concreta que condiciona la validez de una receta y su probabilidad de ser recomendada.

### Restricción Débil

Limitación que condiciona la probabilidad de que una receta sea recomendada o no, con el siguiente orden de importancia:
    - Apetencia 
    - Tiempo deseado de cocción. 

### Restricción Fuerte

Limitación que invalida una receta, con el siguiente orden de importancia: 
    - Inventario 
    - Alergias, Intolerancias 
    - Dieta 
    - Límite real del tiempo de cocción.