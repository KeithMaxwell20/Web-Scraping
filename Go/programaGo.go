package main

import (
	"fmt"
	"log"
	"os"
	"unicode"

	"github.com/gocolly/colly/v2"
)

func main() {

	fmt.Println("Iniciando Programa...")
	fmt.Println("Creando Archivo...")
	myfile, err := os.Create("ResultadosGO.txt")

	// error al crear el archivo.
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Archivo Creado con Exito!!!")

	// Definimos la lista de los Lenguajes
	listaLenguajesNombres := [20]string{"Python", "C", "Java", "C++", "C#", "Visual Basic", "JavaScript",
		"Assembly Language", "SQL", "PHP", "R", "Delphi/Object Pascal", "Go", "Swift", "Ruby",
		"Classic Visual Basic", "Objective-C", "Perl", "Lua", "MATLAB"}

	listaLenguajesTopic := [20]string{"python", "c", "java", "cpp", "csharp", "visual-basic", "javascript",
		"assembly", "sql", "php", "r", "object-pascal", "go", "swift", "ruby", "visual-basic", "objective-c",
		"perl", "lua", "matlab"}

	fmt.Println("Procesando...")
	for i := 0; i < 20; i++ {
		procesar(listaLenguajesNombres[i], listaLenguajesTopic[i], myfile)
	}

	fmt.Println("Cerrando Archivo...")
	myfile.Close()
	fmt.Println("Programa Finalizado!")
}

// Realiza el proceso de scrapping
func procesar(lenguajeNomnbre string, lenguajeLink string, archivo *os.File) {
	// Link para el get
	linkTopic := "https://github.com/topics/" + lenguajeLink

	// Scrapping
	c := colly.NewCollector()

	// Evento
	c.OnHTML("h2.h3.color-fg-muted", func(e *colly.HTMLElement) {
		numeroRepositorios := extractNumber(e.Text)
		archivo.WriteString(lenguajeNomnbre + "," + numeroRepositorios)
		archivo.WriteString("\n")
	})

	c.Visit(linkTopic)
}

//Extrae el numero en el formato correcto
func extractNumber(text string) string {
	string_return := ""
	for i := 0; i < len(text); i++ {
		if unicode.IsDigit(rune(text[i])) {
			string_return += string(text[i])
		}
	}
	return string_return
}
