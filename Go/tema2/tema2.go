package tema2

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	//"os/exec"
	//"runtime"
	"strconv"
	"strings"
	"time"
	"Go/grafico"
	"Go/tabla"

	"github.com/go-rod/rod"
	"github.com/gocolly/colly/v2"
)

var CANTIDAD_BARRAS int = 20

type Resultados struct {
	horaUltimaAct string
	listaTopics   []string
}

func Run(nombreaArchivo1, directorioArchivo1, nombreArchivo2, directorioArchivo2 string) {
	fmt.Println("Empieza el programa...")
	localFileName := "data.html"
	link := "https://github.com/topics/nodejs?o=desc&s=updated"
	cantRepeticiones := 33

	fmt.Println("Generando Archivo.\nEste proceso puede tomar un tiempo...")

	// Generamos el archivo local con todos los articulos
	generandoArchivoHTML(localFileName, link, cantRepeticiones)

	fmt.Println("Teminado el proceso de generacion de archivo!")

	var listaArticulos []Resultados

	listaArticulos = getHorasTopics(localFileName)

	// Guardamos los resultados en un archivo
	guardarListaResultados(directorioArchivo1 + "/" + nombreaArchivo1, listaArticulos)

	// Eliminamos los resultados que tienen mas de 30 dias de antigüedad
	listaArticulos = limpiarLista(listaArticulos)

	fmt.Println("Cant. Total de Articulos: " + strconv.Itoa(len(listaArticulos)))
	var mapeo map[string]int
	mapeo = contarTopics(listaArticulos)

	// Pasando el array para ordenar
	lista := make(grafico.Resultado, 0, len(mapeo))

	for topic, apariciones := range mapeo {
		lista = append(lista, grafico.Datos{Nombre: topic, NroApariciones: apariciones})
	}

	fmt.Println("Presione ENTER para continuar...")
	bufio.NewReader(os.Stdin).ReadString('\n')
	// Ordenando
	grafico.OrdenarListaResultados(&lista)
	fmt.Println("Lista Ordenada de Resultados: ")

	// Encabezados
	var encabezado = []string {"TOPIC", "MENCIONES"}

	matrizResultados := make([][]string, 0)
	for i := 0; i < len(lista); i++ {
		nombre := lista[i].Nombre
		cantidad := fmt.Sprintf("%d", lista[i].NroApariciones)
		var fila = []string {nombre, cantidad} 
		matrizResultados = append(matrizResultados, fila)
	}
	
	tabla.GenerarTabla(encabezado, matrizResultados)
	
	verificarExisteDirectorio(directorioArchivo2)
	fmt.Println("A continuación, se genera el archivo del gráfico de barras.")

	texto := grafico.Info{
		Titulo: "Temas Relacionados con Nodejs",
		Subtitulo: "Los 20 primeros temas, de mayor a menor",
		EjeX: "NOMBRE",
		EjeY: "Nº de Apariciones"}

	grafico.GenerarGraficoBarras(lista, directorioArchivo2, nombreArchivo2, CANTIDAD_BARRAS, texto)
	fmt.Println("Fin del programa")
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

// Contamos el nro de apariciones de cada topic.
// Retornamos los resultados en un mapeo
func contarTopics(lista []Resultados) map[string]int {
	var aux map[string]int
	aux = make(map[string]int)

	for i := 0; i < len(lista); i++ {
		for j := 0; j < len(lista[i].listaTopics); j++ {
			// Si el texto aun no existe en el mapeo
			if aux[lista[i].listaTopics[j]] == 0 {
				aux[lista[i].listaTopics[j]] = 1
			} else { // Si el texto ya existe en el mapeo
				aux[lista[i].listaTopics[j]]++
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

	// Creamos un protocolo de transporte para manipular archivos locales

	t := &http.Transport{}
	t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))
	d := &http.Client{Transport: t}
	c.WithTransport(d.Transport)
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	fmt.Println("Nuevo Path: ")
	pathArchivoLocal := "file://" + pathLocal(path) + link
	c.Visit(pathArchivoLocal)
	return listaRetorno
}

// Cambia los "\" por "/" y borra mencion del disco actual
func pathLocal(oldPath string) string {
	s := strings.ReplaceAll(oldPath, "\\", "/")
	return s[strings.Index(s, ":")+1:] + "/"
}

// Accede a la pagina web, realiza el proceso "Load More"
// de manera consecutivas y retorna el archivo .html generado
func generandoArchivoHTML(nombreArchivo string, link string, cantidad int) {

	page := rod.New().
		MustConnect().
		MustPage(link)

	// Esperamos a que toda la pagina se haya cargado
	page.MustElement(`body > footer`).MustWaitVisible()

	// Realizamos los clicks consecutivos al boton "Load More"
	for i := 0; i < cantidad; i++ {
		fmt.Printf("...")
		page.MustElement(`button.ajax-pagination-btn.btn.btn-outline.color-border-default.f6.mt-0.width-full`).MustClick()
		page.MustElement(`body > footer`).MustWaitVisible()
		time.Sleep(3 * time.Second)
	}

	f, err := os.Create(nombreArchivo)
	// Error
	if err != nil {
		log.Println(err)
	}
	// Recuperamos la pagina resultante completa
	data, _ := page.HTML()
	// Escribimos en el archivo
	f.WriteString(data)

	f.Close()
}

// Guarda los resultados de la lista en un archivo .csv
func guardarListaResultados(direccionNombre string, lista []Resultados) {
	archivo, err := os.Create(direccionNombre)

	if err != nil {
		fmt.Println(err)
	}

	// Concatenamos cada lista y escribimos en el archivo
	for i := 0; i < len(lista); i++ {
		aux := strings.Join(lista[i].listaTopics, ",")
		archivo.WriteString(aux)
	}

	archivo.Close()
}

// Eliminamos los resultados que tienen mas de 30 dias de antigüedad
func limpiarLista(lista []Resultados) []Resultados {
	listaNueva := make([]Resultados, 0)

	// Tiempo Actual y de Hace Mes
	tiempoActual := time.Now().UTC()
	tiempoHaceUnMes := tiempoActual.AddDate(0, -1, 0)

	for i := 0; i < len(lista); i++ {
		tiempoArticulo, _ := time.Parse(time.RFC3339, lista[i].horaUltimaAct)
		if tiempoArticulo.After(tiempoHaceUnMes) == true {
			var aux Resultados
			aux.horaUltimaAct = lista[i].horaUltimaAct
			aux.listaTopics = lista[i].listaTopics
			listaNueva = append(listaNueva, aux)
		}
	}

	return listaNueva
}
