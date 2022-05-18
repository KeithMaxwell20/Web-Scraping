package tema1

import (
	"Go/grafico"
	"Go/tabla"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"

	"github.com/gocolly/colly/v2"
)

// Para almacenar el Nombre y el topic a buscar de cada lenguaje
type NombreTopic struct {
	nombre string
	topic  string
}

// Para almacenar el Nombre del lenguaje, rating y la cantidad de repeticiones
type InfoLenguaje struct {
	nombre   string
	rating   float64
	cantidad int
}

var CANTIDAD_BARRAS int = 10

func Run(nombreaArchivo1, directorioArchivo1, nombreArchivo2, directorioArchivo2 string) {

	fmt.Println("Iniciando Programa...")

	verificarExisteDirectorio(directorioArchivo1 + "/")

	//Leyendo del archivo .csv
	archivoEntrada := ".././tiobe-list.csv"
	listaLenguajes := extraerDatosEntrada(archivoEntrada)

	archivoSalida, err := os.Create(directorioArchivo1 + "/" + nombreaArchivo1)

	// error al crear el archivo.
	if err != nil {
		log.Fatal(err)
	}

	// Extraer la cantidad de repositorios para cada lenguaje
	var listaResultados [20]InfoLenguaje //Guardamos los resultados para las tablas
	fmt.Println("Procesando...")
	for i := 0; i < 20; i++ {
		listaResultados[i] = procesar(listaLenguajes[i].nombre, listaLenguajes[i].topic, archivoSalida)
		fmt.Printf("... ")
	}
	archivoSalida.Close()
	fmt.Printf("\n")
	fmt.Println("Proceso Completo!")

	// Ordenando Lista
	ordenarListaResultados(&listaResultados)
	calcularRating(&listaResultados)

	// Encabezados
	var encabezado = []string{"LENGUAJE", "RATING", "REPOSITORIOS"}
	// Lista de Lenguajes con su rating
	matrizResultados := make([][]string, 0)
	for i := 0; i < len(listaResultados); i++ {
		nombre := listaResultados[i].nombre
		rating := fmt.Sprintf("%.3f", listaResultados[i].rating)
		cantidad := fmt.Sprintf("%d", listaResultados[i].cantidad)

		var fila = []string{nombre, rating, cantidad}
		matrizResultados = append(matrizResultados, fila)
	}

	tabla.GenerarTabla(encabezado, matrizResultados)

	fmt.Println("A continuación, se genera el archivo del gráfico de barras.")
	listaGraficar := make(grafico.Resultado, 0)
	for i := 0; i < CANTIDAD_BARRAS; i++ {
		aux := grafico.Datos{Nombre: listaResultados[i].nombre, NroApariciones: listaResultados[i].cantidad}
		listaGraficar = append(listaGraficar, aux)
	}

	texto := grafico.Info{
		Titulo:    "Lenguajes con Mayor Nº de Apariciones",
		Subtitulo: "Los 10 primeros lenguajes, ordenados de mayor a menor",
		EjeX:      "NOMBRE",
		EjeY:      "Nº de Apariciones"}

	grafico.GenerarGraficoBarras(listaGraficar, directorioArchivo2, nombreArchivo2, CANTIDAD_BARRAS, texto)
	fmt.Println("Programa Finalizado!")
}

// Verifica si existe el directorio "dir"
// Si no existe, lo crea
// Si ya existe, no hace nada
func verificarExisteDirectorio(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
}

// Ordena la lista de resultados, de mayor a menor
func ordenarListaResultados(lista *[20]InfoLenguaje) {
	for i := 0; i < 19; i++ {
		for j := i + 1; j < 20; j++ {
			if lista[i].cantidad < lista[j].cantidad {
				aux := lista[i]
				lista[i] = lista[j]
				lista[j] = aux
			}
		}
	}
}

// Calcula el rating y agrega el resultado al campo de cada registro
func calcularRating(lista *[20]InfoLenguaje) {
	max := maximo(*lista)
	min := minimo(*lista)
	for i := 0; i < 20; i++ {
		lista[i].rating = (float64(lista[i].cantidad-min) / float64(max-min)) * 100.0
	}
}

// Retorna el valor de la cantidad Maxima de repositorios
func maximo(lista [20]InfoLenguaje) int {
	valorRetorno := lista[0].cantidad
	for i := 1; i < 20; i++ {
		if valorRetorno < lista[i].cantidad {
			valorRetorno = lista[i].cantidad
		}
	}
	return valorRetorno
}

// Retorna el valor de la cantidad Minima de repositorios
func minimo(lista [20]InfoLenguaje) int {
	valorRetorno := lista[0].cantidad
	for i := 1; i < 20; i++ {
		if valorRetorno > lista[i].cantidad {
			valorRetorno = lista[i].cantidad
		}
	}
	return valorRetorno
}

// Extrae los registros de archivo.csv y retorna en un
// array struct
func extraerDatosEntrada(archivoEntrada string) [20]NombreTopic {
	var listaRegistro [20]NombreTopic
	records, err := readData(archivoEntrada)

	if err != nil {
		log.Fatal(err)
	}
	var ind = 0
	for _, record := range records {
		aux := NombreTopic{
			nombre: record[0],
			topic:  record[1],
		}
		listaRegistro[ind] = aux
		ind++
	}
	return listaRegistro
}

// Lee un registro del archivo .csv
func readData(archivo string) ([][]string, error) {
	f, err := os.Open(archivo)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()
	r := csv.NewReader(f)
	// saltear la primera linea
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}
	records, err := r.ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return records, nil
}

// Realiza el proceso de scraping
func procesar(lenguajeNombre string, lenguajeLink string, archivo *os.File) InfoLenguaje {
	var registro InfoLenguaje // Registro a retornar

	// Link para el get
	linkTopic := "https://github.com/topics/" + lenguajeLink

	c := colly.NewCollector()

	// Evento
	c.OnHTML("h2.h3.color-fg-muted", func(e *colly.HTMLElement) {
		numeroRepositorios := extractNumber(e.Text)
		registro.nombre = lenguajeNombre
		registro.cantidad, _ = strconv.Atoi(numeroRepositorios)
		archivo.WriteString(lenguajeNombre + "," + numeroRepositorios)
		archivo.WriteString("\n")
	})

	c.Visit(linkTopic)

	return registro
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
