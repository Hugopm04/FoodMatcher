# Servicio Móvil FoodMatcher

## Descripción del problema

Tomar la decisión de qué cocinar para luego comerlo si debe haber un consenso mutuo entre un grupo de personas familiares entre sí que se encuentran presencialmente juntos presenta dificultades debido a lo compleja que puede volverse dicha decisión.
Esta depende de factores muy limitantes, como los alimentos disponibles o las alergias e intolerancias pasando por factores más situacionales, como el tiempo de cocción o el esfuerzo necesario y, finalmente, los más subjetivos, como lo que le apetece comer a cada individuo. Todos ellos son determinantes para la comida final, pues si se requiere un consenso al 100%, basta con que a una persona no le apetezca X para que dicha comida sea descartada.
Tener en cuenta todos estos factores para llegar a una decisión común no es tarea fácil y más teniendo en cuenta que la complejidad aumenta de forma considerable conforme aumenta el número de personas.

## Los Datos Necesarios

### Introducidos por Usuarios
Los usuarios poseen la información de qué alimentos tienen, la situación y preferencias de cada uno y pueden introducirlas manualmente.
La situación se compone de limitaciones fuertes en el siguiente orden:
- Alergias
- Intolerancias
- Estilos de alimentación (vegetarianismo, veganismo, ...)
- Límite real de tiempo de cocción

Las preferencias son limitaciones débiles en el siguiente orden:
- Que apetece o no apetece al usuario
- Límite deseado de tiempo de cocción

### Introducción por Desarrolladores

Las recetas van a ser construidas y almacenadas en el repositorio siguiendo el siguiente proceso:
    - Obtención de las mismas en crudo través de: [Sección Artes Culinarias de _Wikilibros_ en formato _xml_](https://dumps.wikimedia.org/eswikibooks/latest/eswikibooks-latest-pages-articles.xml.bz2), que contiene una colección de recetas.
    - Procesamiento y estructuración de los datos obtenidos: Primero se procesará el xml con un _script_ en _Python_ para extraer las recetas y su información por campos. Luego, aunque la mayoría de las recetas siguen unos estándares en cuestiones de campos rellenados y formatos, otros, como tiempos de cocción o cantidades, a veces vienen en formatos distintos. De manera que una limpieza posterior es obligatoria.
    - Etiquetado de productos alérgenos y dietas usando las _Taxonomías de Open Food Fact_ [alérgenos](https://github.com/openfoodfacts/openfoodfacts-server/blob/main/taxonomies/allergens.txt) y [taxonomía](https://github.com/openfoodfacts/openfoodfacts-server/blob/main/taxonomies/food/ingredients.txt), que incluye nombres y sinónimos del mismo alimento asociados a la alergia que podrían producir.
        - La taxonomía incluye un amplio vocabulario con nombres de alimentos, sinónimos y de qué alimentos principales desciende. El queso desciende de la leche, heredando sus cualidades (por ejemplo). El documento incluye etiquetas con alérgenos y dietas para cada alimento en múltiples idiomas. El documento permite, dado un alimento en español, encontrar su alimento padre en inglés.
        - El documento de alérgenos posee el nombre del alimento padre vínculado a la alergia que puede provocar junto con su traducción a distintos idiomas.
        - Del primero se obtiene si un alimento es un alérgeno y, siguiendo la herencia y llegando al alimento padre del mismo, se traduce el padre usando la tabla del segundo documento.
        - Si es apto para una dieta o no viene en la taxonomía.
        - Ambos documentos habrán de ser procesados por un _script_ en _Python_ para extraer los datos relevantes (no nos interesan los alimentos en otros idiomas por ejemplo). 
    - Etiquetado de intolerancias, usnado las [_Tablas CIQUAL 2020 (ANSES)_](https://www.data.gouv.fr/api/1/datasets/r/e31dd87c-8ad0-43e4-bdaa-af84ad243dc6) que contiene 5 documentos de los cuáles se usaran 2:
        - El diccionario (archivo const_2020_07_07.xml del zip): La taxonomía previa contiene un código CIQUAL que al contrastarlo con el diccionario devuelve punteros al siguiente archivo.
        - Las composiciones (archivo compo_2020_07_07.xml del zip): Los punteros previos apuntan cada uno a una fila de este documento que indica la cantidad de alérgeno o intolerante que el alimento posee cada 100g. Dichos datos pueden procesarse para establecer un umbral de lo que podría generar intolerancia o no (Contiene X, tiene trazas de X, libre de X).
    - La idea principal es seguir un proceso iterativo en el que se comience haciendo la fusión más básica, y se evolucione a incluir términos más específicos que que no estén contemplados en las bases de datos. Plurales, sinónimos que falten, etc.

## Condiciones de Éxito
El problema se da por resuelto si se encuentra una receta que satisfaga todas las exigencias del grupo y respete todos los límites. De no existir dicha receta, la solución al problema será aquella que satisfaga la mayor cantidad de limitaciones fuertes por orden, y en caso de empate, la mayor cantidad de limitaciones débiles por orden.

## Configuración del Repositorio
Puede ser encontrada en: [configuración](docs/configuracion.md)

## Tarjetas del Juego de Rol

### Tarjeta 0.2

Tarjeta de cliente actualizada. Entrevista real a una persona con el problema, incluyendo preguntas y respuestas.
![Tarjeta de validación del juego de rol 0.2](img/Juego%20de%20Rol%200.2.jpg)

### Tarjeta 0.1

Acredita que estuve en clase el día que se realizó la actividad, pero ha sido sustituida por otra atendiendo a las peticiones del profesor.

![Tarjeta de validación del juego de rol 0.1](img/Juego%20de%20Rol%200.1.jpg)