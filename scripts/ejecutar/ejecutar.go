package ejecutar

import (
	"os"
	"os/exec"
)

const directorioProyecto = "../internal"

func Comando(nombre string, argumentos ...string) error {
	comando := exec.Command(nombre, argumentos...)
	comando.Dir = directorioProyecto
	comando.Stdout = os.Stdout
	comando.Stderr = os.Stderr
	return comando.Run()
}

func Salir(err error) {
	if err != nil {
		os.Exit(1)
	}
}
