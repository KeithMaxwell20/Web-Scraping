package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

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

	var mapeo map[string]int
	mapeo = contarTopics(listaArticulos)

	for topic, apariciones := range mapeo {
		fmt.Printf("mapeo[%s] = %d\n", topic, apariciones)
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
