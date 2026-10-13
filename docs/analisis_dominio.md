# Análisis del dominio de la HU001

Análisis de la [HU001](historias_usuario.md#hu001) para el Milestone 0.

## Lenguaje (#7)

Se usa Go, decisión conjunta tomada con el propietario en #7 , por los siguientes motivos:

- Los tipos son obligatorios y el compilador los comprueba antes de ejecutar luego no se confunden conceptos distintos del modelo.
- Los campos que empiezan por minúscula son privados por tanto cada objeto solo se crea con su constructor, que valida los datos y no se puede modificar desde fuera.
- Los errores se devuelven como valores lo que permite rechazar datos inválidos como un tiempo menor o igual que cero.
- El formato y las comprobaciones vienen incluidos en el propio lenguaje, sin necesidad de instalar herramientas aparte. Destacar sintaxis sencilla.
- Se instala en el espacio de usuario con mise.

## Tiempo (#8)

El glosario habla de tiempo de cocción, pero se interpreta como el tiempo total de preparación porque la HU pide que la receta no tarde más de 4 horas. Y hay recetas que llevan tiempo aunque no se cuezan. 
Se mide en minutos enteros y siempre es mayor que cero. Se trata de un objeto valor porque incluye esa validación y lo usan tanto la receta como el límite de tiempo del proceso de elección. Destacar que pese a que tenga nada más que un atributo se crea una estructura propia porque tiene una regla asociada(mayor que cero) y se usa en dos sitios, lo que permite que la validación no se repita.


## Dieta (#9)

Se tienen tres dietas, a saber, libre, vegetariana y vegana. Libre es quien puede comer de todo y se relaciona con quien no sigue una dieta específica dentro de las personas. Cada una es más restrictiva que la anterior, así que una receta vegana vale también para un vegetariano y para alguien con dieta libre.
Es un objeto valor y al ser una lista cerrada de valores se representa con un enumerado.

## Alergias (#10)

Solo se admiten los alérgenos de la lista de Open Food Facts, ya que según el README es con la que se etiquetan las recetas. De hecho una alergia fuera de esa lista no se podría comprobar.
Es un objeto valor y al ser una lista cerrada, se representa con un enumerado. Por otro lado un alimento hereda los alérgenos de su alimento padre. Por último destacar que una persona puede tener ninguna o varias alergias.

## Intolerancias (#11)

Se nombran por el componente que no se tolera. Solo se admiten las intolerancias cuyo componente aparece en la lista de alérgenos de Open Food Facts, que es la misma con la que se etiquetan las recetas, ya que otro componente no se podría comprobar. Por eso son el gluten (celiaquía) y la lactosa, que se comprueba con la leche, coincidiendo con las intolerancias que aparecen en las personas. Cada alimento puede contener ese componente, tener trazas o estar libre de él. Como se trata de una restricción fuerte, solo es apta una receta libre del componente, es decir, las trazas también la descartan. Destacar que una persona puede tener varias intolerancias o ninguna. Es un objeto valor y al ser una lista cerrada de componentes, se representan con un enumerado.

## Alimentos (#12)

Un alimento se trata de uno de los que aparecen en la taxonomía de Open Food Facts, de manera que nombres distintos de un mismo alimento corresponden al mismo alimento. Destacar que un alimento conoce a su alimento padre, sus alérgenos, los componentes que contiene o  de los que tiene trazas y las dietas para las que es apto. Es un objeto valor, ya que para la HU solo importan sus propiedades que no cambian. No cambia una vez creado y dos alimentos con los mismos datos son el mimso alimento. No puede tener nombre vacío. Destacar que un alimento que no está en la taxonomía no se admite.

## Inventario (#13)

Se trata del conjunto de alimentos que hay en casa, sin cantidades, porque la HU solo habla de los alimentos que tenga en casa. Puede estar vacío, ya que es una situación real, en dicho caso ninguna receta cumplirá el inventario. Una receta solo lo cumple si todos sus ingredientes están en él, ya que según el glosario el inventario se trata de una restricción fuerte. Se trata de un objeto valor , puesto que dos inventarios con los mismos alimentos serían iguales. Como no tiene más datos ni reglas propias, no se crea una clase asociada.

## Receta (#14)

Una receta tiene nombre, ingredientes y tiempo. Destacar que los pasos son necesarios porque solo con los ingrediente sy el tiempo no se puede preparar. Sus alérgenos, intorelancias y dietas no se guardan aparte, sino que se deducen de sus ingredientes. Se trata de un objeto valor porque no cambia una vez creada y se define por su contenido. De manera que dos recetas con los mismos ingredientes, pasos y tiempo son la misma receta. Debe tener nombre, al menos un ingrediente y al menos un paso.

## Usuario (#15)

El usuario es una entidad identificada por su nick, que elige en la configuración inicial del user journey. Contiene sus alergias, intolerancias y dieta. La dieta siempre se indica porque quien no sigue ninguna elige la dieta libre, luego no hay una dieta por defecto. Según la configuración opcional del user journey, sus restricciones pueden cambiarse después, sin que deje de ser el mismo usuario.

## Proceso de elección (#15)

El inventario y el límite de tiempo pertenecen a cada proceso de elección, porque cambian según el momento. Según el user journey, los alimentos se actualizan en cada uso, y en la HU el límite de 4 horas corresponde a una mañana concreta.El proceso de elección es un objeto valor que agrupa el usuario, el inventario y el límite de tiempo de una búsqueda.



