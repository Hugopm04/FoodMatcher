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

### Prestaciones

Issue: [#47](https://github.com/Hugopm04/FoodMatcher/issues/47)

Se comparan tiempos de ejecución, tiempos de descarga y peso.

Make tiene la ventaja inherente de venir instalado por defecto en las distribuciones de ubuntu que usa github para los workflows, en las imágenes de docker aunque no venga instalado por defecto pesa solo 0.386 MB, siendo rápido de descargar e instalar.

Task por otro lado no viene instalado y pesa 15.8 MB, considerablemente más pesado que Make.

Mage no viene por defecto en Go, por lo que hay que considerar también su peso de 1.78 MB.

Los resultados de velocidad son los siguientes (medidos usando oh my posh, en la línea del comando se ve lo que ha tardado el anterior). La lista es el resumen, las capturas reales las pongo debajo.

- Make
  - `installdeps`: 34 ms
  - `check`: 25 ms
  - `install` (frío): 54 ms
  - `install` (caliente): 39 ms
  - `test`: 47 ms
- Task
  - `installdeps`: 93 ms
  - `check`: 103 ms
  - `install` (frío): 109 ms
  - `install` (caliente): 94 ms
  - `test`: 114 ms
- Mage
  - `installdeps`: 90 ms
  - `check`: 82 ms
  - `install` (frío): 3,64 s
  - `install` (caliente): 96 ms
  - `test`: 97 ms

Imágenes en la terminal:

![Tiempos Make](https://github.com/user-attachments/assets/6a35f033-f72f-4406-8816-744875f1a644)

![Tiempos Task](https://github.com/user-attachments/assets/27b76fa1-9d70-45bf-a941-45ef98e1cd50)

![Tiempos de Mage](https://github.com/user-attachments/assets/0e951536-b76b-4463-ae1e-113ddf5ff59f)

Los resultados dejan a Make como ganador indiscutible, tanto en peso como en velocidad y compatibilidad con las mejores prácticas.