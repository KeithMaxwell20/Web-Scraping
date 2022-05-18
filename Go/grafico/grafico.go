package grafico

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type Datos struct {
	Nombre         string
	NroApariciones int
}

type Resultado []Datos

type Info struct {
	Titulo    string
	Subtitulo string
	EjeX      string
	EjeY      string
}

// Genera la lista de cantidades para el gráfico de barras
func generarItems(lista Resultado, cantBarras int) []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < cantBarras; i++ {
		items = append(items, opts.BarData{Value: lista[i].NroApariciones})
	}
	return items
}

// Generar grafico de barras.
func GenerarGraficoBarras(lista Resultado, directorioArchivo, nombreArchivo string, cantBarras int, texto Info) {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			//			Title:    "Temas relacionados con Nodejs",
			Title: texto.Titulo,
			//			Subtitle: "Los primeros 20 temas, de mayor a menor",
			Subtitle: texto.Subtitulo,
			Right:    "30%"}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
		charts.WithLegendOpts(opts.Legend{Right: "80%"}),
		charts.WithXAxisOpts(opts.XAxis{
			//Name: "NOMBRE",
			Name: texto.EjeX,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			//Name: "Nº APARICIONES",
			Name: texto.EjeY,
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Width:  "1500px",
			Height: "600px",
		}),
	)

	nombres := make([]string, cantBarras)

	// Generando los valores del eje X
	for i := 0; i < cantBarras; i++ {
		nombres[i] = lista[i].Nombre
	}

	bar.SetXAxis(nombres).
		AddSeries("Lenguaje", generarItems(lista, cantBarras))

	f, _ := os.Create(directorioArchivo + "/" + nombreArchivo)
	bar.Render(f)

	fmt.Println("Presione ENTER para continuar...")
	bufio.NewReader(os.Stdin).ReadString('\n')
	openBrowser(directorioArchivo, nombreArchivo)
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

func OrdenarListaResultados(lista *Resultado) {
	for i := 0; i < len(*lista)-1; i++ {
		for j := i + 1; j < len(*lista); j++ {
			if (*lista)[i].NroApariciones < (*lista)[j].NroApariciones {
				aux := (*lista)[i]
				(*lista)[i] = (*lista)[j]
				(*lista)[j] = aux
			}
		}
	}
}
