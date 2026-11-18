package main

import "foodmatcher/scripts/ejecutar"

func main() {
	ejecutar.Salir(ejecutar.Comando("go", "test", "./..."))
}
