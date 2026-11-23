package main

import "foodmatcher/scripts/ejecutar"

func main() {
	_, err := ejecutar.Salida("gofmt", "-e", ".")
	ejecutar.Salir(err)
}
