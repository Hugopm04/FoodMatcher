# Servicio Móvil FoodMatcher

## Descripción del problema:

Tomar la decisión de qué cocinar para luego comerlo si debe haber un consenso mutuo entre un grupo de personas familiares entre sí que se encuentran presencialmente juntos presenta dificultades debido a lo compleja que puede volverse dicha decisión.
Esta depende de factores muy limitantes, como los alimentos disponibles o las alergias e intolerancias pasando por factores más situacionales, como el tiempo de cocción o el esfuerzo necesario y, finalmente, los más subjetivos, como lo que le apetece comer a cada individuo. Todos ellos son determinantes para la comida final, pues si se requiere un consenso al 100%, basta con que a una persona no le apetezca X para que dicha comida sea descartada.
Tener en cuenta todos estos factores para llegar a una decisión común no es tarea fácil y más teniendo en cuenta que la complejidad aumenta de forma considerable conforme aumenta el número de personas.

## Los Datos Necesarios

Los usuarios poseen la información de qué alimentos tienen, la situación y preferencias de cada uno y pueden introducirlas manualmente.
La situación se compone de limitaciones fuertes en el siguiente orden:
- Alergias
- Intolerancias
- Estilos de alimentación (vegetarianismo, veganismo, ...)
- Límite real de tiempo de cocción

Las preferencias son limitaciones débiles en el siguiente orden:
- Que apetece o no apetece al usuario
- Límite deseado de tiempo de cocción

El problema se da por resuelto si se encuentra una receta que satisfaga todas las exigencias del grupo y respete todos los límites. De no existir dicha receta, la solución al problema será aquella que satisfaga la mayor cantidad de limitaciones fuertes por orden, y en caso de empate, la mayor cantidad de limitaciones débiles por orden.
Las recetas se escogen de un catálogo del repositorio.
El catálogo del repositorio es obtenido de gousto.co.uk a través de su API: https://production-api.gousto.co.uk/cmsreadbroker/v1/recipe/
Filtraré alérgenos y alimentos a los que se puedan tener intolerancias con ayuda de la base de datos existente en la api api.nal.usda.gov/fdc/v1.
Este proceso se hará una vez y luego se mantendrá el catálogo en el repositorio.

## Configuración del Repositorio
Puede ser encontrada en: [configuración](docs/configuracion.md)

## Tarjetas del Juego de Rol

### Tarjeta 0.2

Tarjeta de cliente actualizada. Entrevista real a una persona con el problema, incluyendo preguntas y respuestas.
![Tarjeta de validación del juego de rol 0.2](img/Juego%20de%20Rol%200.2.jpg)

### Tarjeta 0.1

Acredita que estuve en clase el día que se realizó la actividad, pero ha sido sustituida por otra atendiendo a las peticiones del profesor.

![Tarjeta de validación del juego de rol 0.1](img/Juego%20de%20Rol%200.1.jpg)