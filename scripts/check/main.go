package main

import (
	"fmt"
	"os"
	"strings"

	"foodmatcher/scripts/ejecutar"
)

func main() {
	salida, err := ejecutar.Salida("gofmt", "-l", "-e", ".")
	if err != nil {
		os.Exit(1)
	}
	mal := strings.TrimSpace(salida)
	if mal != "" {
		fmt.Println(mal)
		os.Exit(1)
	}
}
