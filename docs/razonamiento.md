# Razonamiento de las decisiones técnicas

Gestor de dependencias: `go mod` [#25](https://github.com/Hugopm04/FoodMatcher/issues/25)
Fichero de tareas: Make con scripts en Go [#24](https://github.com/Hugopm04/FoodMatcher/issues/24)

## Gestor de dependencias

Atendiendo al [portal oficial de desarrollo de go](https://go.dev/ref/mod) go maneja las dependencias en forma de módulos y es una parte que viene integrada en el lenguaje y siendo obligatorio usarlo a partir de la versión 1.16 por defecto en **_Go Modules_**. Otras alternativas existieron anteriormente en go: dep, glide, godep... Pero todas han quedado obsoletas y descontinuadas.

## Fichero de tareas

### Opciones disponibles

Issue: [#44](https://github.com/Hugopm04/FoodMatcher/issues/44)

Según la [documentación oficial de go](https://go.dev/doc/articles/go_command) uno de los objetivos que debía cumplir Go es ser compilable sin la necesidad de un fichero de tareas:

> _"An explicit goal for Go from the beginning was to be able to build Go code using only the information found in the source itself, not needing to write a makefile or one of the many modern replacements for makefiles."_

No obstante, nosotros no vamos solo a compilar, sino también a ejecutar tests, instalar dependencias, etc... Por tanto está justificada la búsqueda de un archivo de tareas.

La opción clásica es [GNU Make](https://www.gnu.org/software/make/manual/make.html), pero también servirían [Task](https://taskfile.dev/) y [Mage](https://magefile.org/).

Parto de estos 3 candidatos.

### Mejores prácticas

Issue: [#45](https://github.com/Hugopm04/FoodMatcher/issues/45)

Lo más parecido a mejores prácticas es el [Standard Go Project Layout](https://github.com/golang-standards/project-layout) que aclara al principio que no es un estándar oficial.

A pesar de que no haya un estándar oficial, el propuesto anteriormente ha sido desarrollado por la comunidad, y se tendrá en cuenta a la hora de tomar las decisiones. En él se habla brevemente de la carpeta _scripts/_ en la que deben ir los ficheros de tareas, de forma que el Makefile (que estaría en la raíz del proyecto) sea más breve. Este proyecto supone entonces el uso de GNU Make junto con scripts auxiliares.

Lo que saco como mejores prácticas es:
- Si necesitas varios ficheros de tareas, van en /scripts
- El fichero de tareas principal debe ir en la raíz del proyecto y ser muy breve

Make es perfectamente capaz de cumplir esos requisitos, pudiendo delegar en scripts de bash, o incluso Go.

Task es perfectamente capaz de cumplir esos requisitos, pudiendo delegar en scripts de Task (.yml) o incluso de Go.

Mage es un poco más problemático para esto. Aunque se pueda tener un único fichero Mage que delegue en otros, sigue rompiendo un poco la filosofía de Mage. Además, deberían estar todos en la carpeta magefiles, sin tener uno en la raíz. Podrían estar en /scripts, pero ya estamos yendo en contra de la filosofía de Mage de nuevo. Por otro lado una ventaja considerable de Mage es que reduce el número de lenguajes de programación usados en el proyecto y por tanto, las dependencias y deuda técnica del mismo (al menos a simple vista).

### Deuda técnica

Issue: [#46](https://github.com/Hugopm04/FoodMatcher/issues/46)

Make tiene el problema de que para ser multiplataforma se debe implementar a conciencia. Make acepta código shell en el Makefile, lo que podría hacerlo incompatible entre SOs. Esto aplica también a los scripts en bash, usar make con bash implica que si se quisiera trabajar en múltiples plataformas habría que crear el mismo script para cada plataforma.

Es cierto que no estamos restringidos a scripts bash, que se pueden usar scripts de go. Además se puede evitar por completo incluir comandos shell (que no sean la invocación a los scripts) para que sea compatible con cualquier plataforma. No obstante esto supone un compromiso que debe de mantenerse en el tiempo y supone una deuda a considerar.

Task es multiplataforma. Es también lo más cercano a Make que hay. Debe mantenerse el Taskfile principal y asegurarse de delegar bien la lógica a los scripts.

Mage es multiplataforma y no requiere de un lenguaje adicional, aunque su sintaxis es muy distinta a la de Make. Dejaría Go como único lenguaje del que dependería el proyecto por el momento. No obstante, es una herramienta mucho más específica que las dos anteriores, lo que también la hace más vulnerable a ser discontinuada acorde al siguiente estudio [Scientific Open-Source Software Is Less Likely to Become Abandoned Than One Might Think! Lessons from Curating a Catalog of Maintained Scientific Software](https://dl.acm.org/doi/10.1145/3729369).

Mi conclusión es que el mayor riesgo de deuda técnica lo presenta Mage, porque a pesar de ser la única opción que mantiene Go como único lenguaje, es una herramienta bastante específica que ya hemos visto que choca con algunos estándares de la comunidad. Diría que el segundo es Make, porque mantenerlo multiplataforma es la diferencia principal respecto a Task en cuestiones de deuda. Es cierto que Make es una herramienta mucho más asentada, pero considero Task lo suficientemente general y con una base lo bastante sólida como para que la diferencia no sea demasiado relevante en cuanto a riesgo de discontinuación.

### Prestaciones

Issue: [#47](https://github.com/Hugopm04/FoodMatcher/issues/47)

Se comparan tiempos de ejecución, tiempos de descarga y peso.

Make tiene la ventaja inherente de venir instalado por defecto en las distribuciones de ubuntu que usa github para los workflows, en las imágenes de docker aunque no venga instalado por defecto pesa solo 0.386 MB, siendo rápido de descargar e instalar.

Task por otro lado no viene instalado y pesa 15.8 MB, considerablemente más pesado que Make.

Mage no viene por defecto en Go, por lo que hay que considerar también su peso de 1.78 MB.

Los resultados de velocidad son los siguientes (medidos usando oh my posh, en la línea del comando se ve lo que ha tardado el anterior). La tabla es el resumen, las capturas reales las pongo debajo.

| Herramienta | `installdeps` | `check` | `install` (frío) | `install` (caliente) | `test` |
|---|---|---|---|---|---|
| Make | 5,302 s | 1,487 s | 357 ms | 138 ms | 1,283 s |
| Task | 5,353 s | 1,589 s | 437 ms | 267 ms | 1,373 s |
| Mage | 7,171 s | 396 ms | 459 ms | 295 ms | 1,401 s |

Imágenes en la terminal:

![Tiempos Make](https://github.com/user-attachments/assets/063a5b98-38bc-4969-ae07-f8cc15c30c35)

![Tiempos Task](https://github.com/user-attachments/assets/28e80af9-6cb5-4062-9cbc-064bba574e74)

![Tiempos de Mage](https://github.com/user-attachments/assets/1d325a81-f0a3-42ff-ac20-657b6bbe06c7)

Los resultados indican que hay poca diferencia entre las herramientas. Aunque Mage es mucho más rápido para hacer el check, es el más lento en los installs en frío y en caliente. Make ha quedado como ganador indiscutible en velocidad.

### Decisión

Issue: [#24](https://github.com/Hugopm04/FoodMatcher/issues/24)

Make es el más rápido con diferencia, ya sea por venir instalado por defecto, por su poco peso o por su tiempo de ejecución. Además permite la convención sugerida por la comunidad de Go respecto a organización de directorios y es un estándar muy aceptado en la industria. La única pega es recordar no utilizar comandos propios del shell dentro del makefile, para no hacerlo dependiente del shell.

Task es ligeramente más lento que make y ocupa bastante espacio, tiene un poco más de riesgo a ser descontinuado, y la única ventaja que ofrece es que no es necesario preocuparse por el shell.

Mage es lento aunque ocupa bastante poco, es el que mayor riesgo tiene a ser descontinuado con diferencia y rompe con los estándares sugeridos por la comunidad.

Tras estos análisis, me decanto por usar Make en conjunto con scripts en Go siguiendo el estándar sugerido por la comunidad, con las mejores prestaciones y menor riesgo de discontinuidad.