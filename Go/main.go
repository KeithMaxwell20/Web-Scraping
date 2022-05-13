package main

import (
	"fmt"
	"os"

	"Go/tema1"
	"Go/tema2"
)

func main()  {
	fmt.Println("Iniciando Programa...")

	nro := os.Args[1]

	if  nro == "1" {
		fmt.Println("Programa 1")
		tema1.Run("ResultadosTema1Go.csv", "Resultados", "GraficoTema1GO.html", "Resultados")
	} else if nro == "2" {
		fmt.Println("Programa 2")
		tema2.Run("ResultadosTema2Go.csv", "Resultados", "GraficoTema2GO.html", "Resultados")
	}
}