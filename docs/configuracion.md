# Configuración del Proyecto

# Configuración del Proyecto

## Claves SSH

SSH permite que cualquiera con la clave privada pueda contribuir en el proyecto.

Clave creada el 17-09-2026:

    ssh-keygen -t ed25519 -C "hugoperezm2004trabajo@gmail.com" -f ~/.ssh/github_ed25519

ed225519 es el procedimiento por el que genero la clave SSH, lo he escogido por ser seguro.

La mitad pública (`~/.ssh/github_ed25519.pub`)
La privada no sale de `~/.ssh` (permisos 600).

La mitad pública está pegada en github mediante el procedimiento de "Crear nueva clave SSH", acredito que he hecho el proceso anterior con la siguiente captura:
![Prueba clave SSH](img/Prueba%20clave%20SSH.jpg)

`~/.ssh/config` asocia la clave al host:

    Host github.com
      HostName github.com
      User git
      IdentityFile ~/.ssh/github_ed25519
      IdentitiesOnly yes

Esto hace me identifica en la terminal, para poder realizar acciones git desde la misma con el mismo nivel de acceso que si lo estuviera haciendo desde la web de github, sin tener que iniciar sesión cada vez.

Comprobación:

    $ ssh -T git@github.com
    Hi Hugopm04! You've successfully authenticated, but GitHub does not provide shell access.

## Identidad en los commits

    $ git config --global user.name
    Hugopm04
    $ git config --global user.email
    hugoperezm2004trabajo@gmail.com

Configurado en `~/.gitconfig` (global, sin sobreescritura local en este repo).

## Herramientas

- **git-iv**: Plugin ofrecido por el profesor para facilitar operaciones relacionadas con la realización de los objetivos. Instalado en `~/.local/bin/git-iv`. Lo uso para crear ramas, cambiar a ellas de manera automática y subir los objetivos.
  `git iv objetivo <n>` y `git iv sube-objetivo`.
- **Avatar**: Hugopm04
- **Nick en la hoja compartida**: `Hugopm04`. 
