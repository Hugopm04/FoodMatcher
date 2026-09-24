# Configuración del Proyecto

## Claves SSH

Las dos opciones para contribuir a un proyecto desde la terminal son SSH o HTTPS. Mediante HTTPS necesitas un token de acceso que solo sirve temporalmente, teniendo que renovarlo de vez en cuando. Por el otro lado la clave SSH solo requiere una configuración inicial y ya se puede usar indefinidamente. 
Y, aunque se pueden crear tokens sin límite de duración, suponen un riesgo de seguridad, ya que estos dependen de un gestor de credenciales que puede dejarlos expuestos.

Utilizo el algoritmo ed25519 para generar la clave porque es el más moderno y estandarizado además de ser muy seguro para lo poco que ocupa.

La mitad pública de la clave está pegada en GitHub mediante el procedimiento de "Crear nueva clave SSH", mientras que la privada está guardada localmente en mi ordenador y nunca sale a la red.

## Identidad en los commits

Debo configurar mi correo y mi usuario de GitHub en la terminal para que se me puedan atribuir a mí las contribuciones que haga al proyecto. La clave sirve para identificar a mi dispositivo, pero el usuario y correo para identificarme a mí como colaborador. Además, git no comprueba el email, lo añade tal cual al commit, recalcando así la necesidad de que el mail coincida con el vinculado con el usuario de GitHub, para que se te pueda identificar correctamente.

## Herramientas

- **git-iv**: Plugin ofrecido por el profesor para facilitar operaciones relacionadas con la realización de los objetivos. Instalado en `~/.local/bin/git-iv`. Lo uso para crear ramas, cambiar a ellas de manera automática y subir los objetivos.
  `git iv objetivo <n>` y `git iv sube-objetivo`.
- **Avatar**: Una foto mía editada por IA en la que salgo ligeramente más fuerte de lo que soy en realidad.
- **Nick en la hoja compartida**: `Hugopm04`. Ya apuntado en la hoja compartida.

## Elección de Licencia

He escogido una licencia [AGPL-3.0](../LICENSE) porque no quiero que nadie pueda crear un producto privado basado en mi proyecto. Elijo AGPL-3.0 en vez de GPL-3.0 porque esta última solo prohíbe la distribución, y, al ser mi aplicación también un servicio web, quiero restringir también el uso de dicho código. 
Quiero estas restricciones porque mi objetivo es privatizar el repositorio al terminar la asignatura con la idea de continuarlo como proyecto personal y lanzarlo yo.