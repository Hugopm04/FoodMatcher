# Configuración del Proyecto

## Claves SSH

Las dos opciones para contribuir a un proyecto desde la terminal son SSH o HTTPS. Mediante HTTPS necesitas un token de acceso que solo sirve temporalmente, teniendo que renovarlo de vez en cuando. Por el otro lado la clave SSH solo require una configuración inicial y ya se puede usar indefinidamente. 
Y, aunque se pueden crear tokens sin límite de duración, supone un riesgo de seguridad. Ya que estos depende de un gestor de credenciales que puede dejarlos expuestos.
Por el contrario la parte privada de la clave ssh nunca se envía a la red y se le puede asignar una *passphrase* para que solo tú puedas acceder a ella.
Por último, el token HTTPS sirve solo para autorizarte el acceso a hacer cambios al directorio, mientras que la clave SSH no solo te autoriza sino que además te identifica. Especialmente relevante para que quede constancia de qué contribuciones haces al proyecto. 

Clave creada el 17-09-2026:

    ssh-keygen -t ed25519 -C "hugoperezm2004trabajo@gmail.com" -f ~/.ssh/github_ed25519

Utilizo el algoritmo ed25519 para generar la clave porque es el más moderno y estandarizado además de ser muy seguro para lo poco que ocupa. La clave se divide en dos partes que están en:

- La mitad pública (`~/.ssh/github_ed25519.pub`)
- La privada no sale de `~/.ssh` (permisos 600).

La mitad pública está pegada en github mediante el procedimiento de "Crear nueva clave SSH", acredito que he hecho el proceso anterior con la siguiente captura:
![Prueba clave SSH](../img/Prueba%20clave%20SSH.png)

`~/.ssh/config` asocia la clave al host:

    Host github.com
      HostName github.com
      User git
      IdentityFile ~/.ssh/github_ed25519
      IdentitiesOnly yes

Esto selecciona la clave que hemos creado para usarla en github sin comprobar ninguna otra.

Comprobación:

    $ ssh -T git@github.com
    Hi Hugopm04! You've successfully authenticated, but GitHub does not provide shell access.

## Identidad en los commits

    $ git config --global user.name
    Hugopm04
    $ git config --global user.email
    hugoperezm2004trabajo@gmail.com

Configurado en `~/.gitconfig` (global, sin sobreescritura local en este repo).

Debo introducir mi correo y mi usuario de github para que se me puedan atribuir a mí las contribuciones que haga al proyecto. La clave sirve para identificar a mi dispositivo, pero el usuario y correo para identificarme a mí como colaborador. Además, git no compureba el email, lo añade tal cual al commit, recalcando así la necesidad de que el mail coincida con el vinculado con el usuario de github, para que se te pueda identificar correctamente.

## Herramientas

- **git-iv**: Plugin ofrecido por el profesor para facilitar operaciones relacionadas con la realización de los objetivos. Instalado en `~/.local/bin/git-iv`. Lo uso para crear ramas, cambiar a ellas de manera automática y subir los objetivos.
  `git iv objetivo <n>` y `git iv sube-objetivo`.
- **Avatar**: Una foto mía editada por IA en la que salgo ligeramente más fuerte de lo que soy en realidad.
- **Nick en la hoja compartida**: `Hugopm04`. Ya apuntado en la hoja compartida.
