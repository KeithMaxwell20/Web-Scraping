package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/gocolly/colly/v2"
)

type Resultado struct {
	nombre         string
	nroApariciones int
}

type Resultados struct {
	horaUltimaAct string
	listaTopics   []string
}

func main() {

	linkTopics := "https://github.com/topics/nodejs?o=desc&s=updated"

	count := 34
	for i := 2; i <= count; i++ {
		linkTopics = linkTopics + "&page=" + strconv.Itoa(i)
	}
	fmt.Println(linkTopics)
	var listaArticulos []Resultados

	listaArticulos = getHorasTopics(linkTopics)
	fmt.Println("Cant. Total de Articulos: " + strconv.Itoa(len(listaArticulos)))
	var mapeo map[string]int
	mapeo = contarTopics(listaArticulos)

	// Pasando a un array para ordenar
	lista := make([]Resultado, 0, len(mapeo))

	for topic, apariciones := range mapeo {
		lista = append(lista, Resultado{nombre: topic, nroApariciones: apariciones})
	}

	// Ordenando
	ordenarListaResultados(&lista)
	fmt.Println("Lista Ordenada de Resultados: ")
	imprimirLista(lista)

	directorio := "../Resultados/"
	archivoGrafico := "GraficosTema2GO.html"
	verificarExisteDirectorio(directorio)
	generarGraficoBarras(lista, directorio+archivoGrafico)
	fmt.Println("A continuación, se genera el archivo del gráfico de barras.")
	fmt.Println("Presione ENTER para continuar...")
	bufio.NewReader(os.Stdin).ReadString('\n')
	openBrowser(directorio, archivoGrafico)
}

// Genera la lista de cantidades para el gráfico de barras
func generarItems(lista []Resultado) []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < 20; i++ {
		items = append(items, opts.BarData{Value: lista[i].nroApariciones})
	}
	return items
}

func generarGraficoBarras(lista []Resultado, direccionArchivo string) {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Temas relacionados con Nodejs",
			Subtitle: "Los primeros 20 temas, de mayor a menor",
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

	var nombres [20]string
	// Generando los valores del eje X
	for i := 0; i < 20; i++ {
		nombres[i] = lista[i].nombre
	}

	bar.SetXAxis(nombres).
		AddSeries("Lenguaje", generarItems(lista))

	f, _ := os.Create(direccionArchivo)
	bar.Render(f)
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

func imprimirLista(lista []Resultado) {
	for i := 0; i < len(lista); i++ {
		fmt.Printf("\n%d) %s:\t%d", i+1, lista[i].nombre, lista[i].nroApariciones)
	}
}

// Ordena la lista de resultados, de mayor a menor
func ordenarListaResultados(lista *[]Resultado) {
	for i := 0; i < len(*lista)-1; i++ {
		for j := i + 1; j < len(*lista); j++ {
			if (*lista)[i].nroApariciones < (*lista)[j].nroApariciones {
				aux := (*lista)[i]
				(*lista)[i] = (*lista)[j]
				(*lista)[j] = aux
			}
		}
	}
}

func contarTopics(lista []Resultados) map[string]int {
	var aux map[string]int
	aux = make(map[string]int)

	// Tiempo Actual y de Hace Mes
	tiempoActual := time.Now().UTC()
	tiempoHaceUnMes := tiempoActual.AddDate(0, -1, 0)

	for i := 0; i < len(lista); i++ {
		tiempoArticulo, _ := time.Parse(time.RFC3339, lista[i].horaUltimaAct)
		if tiempoArticulo.After(tiempoHaceUnMes) == true {
			for j := 0; j < len(lista[i].listaTopics); j++ {
				// Si el texto aun no existe en el mapeo
				if aux[lista[i].listaTopics[j]] == 0 {
					aux[lista[i].listaTopics[j]] = 1
				} else { // Si el texto ya existe en el mapeo
					aux[lista[i].listaTopics[j]]++
				}
			}
		}
	}
	return aux
}

//
func getHorasTopics(link string) []Resultados {

	var listaRetorno []Resultados

	c := colly.NewCollector()

	c.OnHTML("article.border.rounded.color-shadow-small.color-bg-subtle.my-4", func(e *colly.HTMLElement) {
		r := Resultados{horaUltimaAct: strings.TrimSpace(e.ChildAttr("relative-time.no-wrap", "datetime"))}
		listaTopics := e.ChildTexts("a.topic-tag.topic-tag-link.f6.mb-2")

		if listaTopics != nil {
			for i := 0; i < len(listaTopics); i++ {
				listaTopics[i] = strings.TrimSpace(listaTopics[i])
			}
			r.listaTopics = listaTopics
			listaRetorno = append(listaRetorno, r)
		}

	})
	c.Visit(link)

	return listaRetorno
}
