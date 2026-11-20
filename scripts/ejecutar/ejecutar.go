package ejecutar

import (
	"os"
	"os/exec"
)

const directorioProyecto = "../internal"

func preparar(nombre string, argumentos []string) *exec.Cmd {
	comando := exec.Command(nombre, argumentos...)
	comando.Dir = directorioProyecto
	comando.Stderr = os.Stderr
	return comando
}

func Comando(nombre string, argumentos ...string) error {
	comando := preparar(nombre, argumentos)
	comando.Stdout = os.Stdout
	return comando.Run()
}

func Salida(nombre string, argumentos ...string) (string, error) {
	salida, err := preparar(nombre, argumentos).Output()
	return string(salida), err
}

func Salir(err error) {
	if err != nil {
		os.Exit(1)
	}
}
