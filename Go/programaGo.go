package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"unicode"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/gocolly/colly/v2"
)

// Para almacenar el Nombre y el topic a buscar de cada lenguaje
type NombreTopic struct {
	nombre string
	topic  string
}

// Para almacenar el Nombre del lenguaje y la cantidad de repeticiones
type NombreCant struct {
	nombre   string
	cantidad int
}

func main() {

	fmt.Println("Iniciando Programa...")

	directorio := "../Resultados/"
	archivoGrafico := "GraficosGO.html"
	verificarExisteDirectorio(directorio)

	//Leyendo del archivo .csv
	archivoEntrada := "../tiobe-list.csv"
	listaLenguajes := extraerDatosEntrada(archivoEntrada)

	fmt.Println("Creando Archivo...")
	archivoSalida, err := os.Create("../Resultados/ResultadosGO.csv")

	// error al crear el archivo.
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Archivo Creado con Exito!!!")

	var listaResultados [20]NombreCant //Guardamos los resultados para las tablas
	fmt.Println("Procesando...")
	for i := 0; i < 20; i++ {
		listaResultados[i] = procesar(listaLenguajes[i].nombre, listaLenguajes[i].topic, archivoSalida)
		fmt.Printf("... ")
	}
	fmt.Println("\nCerrando Archivo...")
	archivoSalida.Close()

	// Ordenando Lista
	ordenarListaResultados(&listaResultados)
	generarGraficoBarras(listaResultados, directorio+archivoGrafico)
	openBrowser(directorio, archivoGrafico)
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

// Selecciona el OS actual y abre la direccion en el navegador por defecto
func openBrowser(directorio string, archivo string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", directorio+archivo).Start()
	case "windows":
		cmd := exec.Command("rundll32", "url.dll,FileProtocolHandler", archivo)
		cmd.Dir = directorio
		cmd.Output()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}

func generarItems(listaLenguajes [20]NombreCant) []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < 10; i++ {
		items = append(items, opts.BarData{Value: listaLenguajes[i].cantidad})
	}
	return items
}

func generarGraficoBarras(listaLenguajes [20]NombreCant, direccionArchivo string) {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Lenguajes con Mayor Nº de Apariciones",
			Subtitle: "Los 10 primeros lenguajes, ordenados de mayor a menor",
			Right:    "30%"}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
		charts.WithLegendOpts(opts.Legend{Right: "80%"}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "NOMBRE",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "Nº APARICIONES",
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Width:  "1500px",
			Height: "600px",
		}),
	)

	var nombres [10]string
	// Generando los valores del eje X
	for i := 0; i < 10; i++ {
		nombres[i] = listaLenguajes[i].nombre
	}

	bar.SetXAxis(nombres).
		AddSeries("Lenguaje", generarItems(listaLenguajes))

	// Where the magic happens
	f, _ := os.Create(direccionArchivo)
	bar.Render(f)
}

// Ordena la lista de resultados, de mayor a menor
func ordenarListaResultados(lista *[20]NombreCant) {
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

// Realiza el proceso de scrapping
func procesar(lenguajeNombre string, lenguajeLink string, archivo *os.File) NombreCant {
	var registro NombreCant // Registro a retornar

	// Link para el get
	linkTopic := "https://github.com/topics/" + lenguajeLink

	// Scrapping
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
