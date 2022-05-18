from selenium import webdriver
import time
from bs4 import BeautifulSoup
from collections import Counter
import itertools
import matplotlib.pyplot as plt
from datetime import datetime
from dateutil import parser
import csv # Para crear archivos del tipo csv
from selenium.webdriver.common.by import By

# Itera a través de las páginas mediante el botón "load more" y devuelve la página total
def obtenerPaginas(url):
    # Para iterar las páginas mediante el botón "load more"
    driver = webdriver.Firefox()
    driver.get(url)

    # Nro de páginas a iterar, la máxima cantidad permitida por github: 34 páginas
    page_num = 0
    total_pag = 33

    # Cliquear el botón de load more hasta que se acabe
    while page_num < total_pag:
        driver.find_element(by=By.CSS_SELECTOR,
                            value='.ajax-pagination-btn.btn.btn-outline.color-border-default.f6.mt-0.width-full').click()
        page_num += 1
        print("Obteniendo la página " + str(page_num))
        time.sleep(5)

    return driver.page_source.encode('utf-8')

# Función para obtener la cantidad total de artículos actualizados en los últimos 30
# días
def obtenerTotalArticulos(pagina):
    # Parser para hallar las fechas
    soup = BeautifulSoup(pagina, "html.parser")

    # Comprobar que la cantidad de días es menor o igual a 30
    # y contar cuántos artículos cumplen
    contador_art = 0

    # Fecha actual
    hoy = datetime.today().date()

    # Profundizamos en los elementos del html para comparar las fechas de los artículos
    resultado_fecha = soup.find_all(class_='border rounded color-shadow-small color-bg-subtle my-4')
    for elemento in resultado_fecha:
        clase = elemento.find(class_='mr-4')


        if not (clase is None):
            fecha = clase.find('relative-time', class_='no-wrap')['datetime']
            fecha_date = parser.parse(fecha).date()
            diferencia = hoy - fecha_date

            if diferencia.days <= 30:  # Si el artículo tiene 30 días o menos
                contador_art += 1

    print("Contador de artículos: ", contador_art)

    return contador_art


# Obtiene todos los tópicos en una sola lista ordenada
def obtenerListaTopics(pagina, contador_art):
    soup = BeautifulSoup(pagina, "html.parser")

    # Lista con la totalidad de los tópicos
    lista_total = []

    # Añadir los artículos a la lista total de tópicos
    articulos = soup.find_all(class_="d-flex flex-wrap border-bottom color-border-muted px-3 pt-2 pb-2")
    for item in itertools.islice(articulos, contador_art):
        resultado = item.find_all(class_="topic-tag topic-tag-link f6 mb-2")
        for topico in resultado:
            lista_total.append(topico.text.strip())

    diccionario_no_ordenado = dict(Counter(lista_total))

    lista_ordenada = sorted(diccionario_no_ordenado.items(), key=lambda x: x[1], reverse=True)

    # Imprimir elementos en la pantalla
    print()
    print("Elementos y su número de apariciones:")

    for item in lista_ordenada:
        print(item[0], ":", item[1])

    return lista_ordenada

# Crear archivo csv para almacenar el nombre de los lenguajes y sus apariciones
def crear_csv(lista_ordenada):
    print("Creando archivo csv...")

    #cantidad_apariciones = []
    lista_topicos = []

    # Guardar resultados en 2 listas distintas para graficar
    for elemento in lista_ordenada:
        lista_topicos.append(elemento[0])

    archivo = open('ResultadosTema2.csv', 'w') #Crear el archivo en modo escritura
    escritor = csv.writer(archivo) # Escritor para el archivo csv

    # Escribimos los nombres de los lenguajes en el archivo csv
    escritor.writerow(lista_topicos)


    print("El archivo csv se ha creado con éxito")


# Grafica los 20 tópicos con mayor número de apariciones de forma descendente
def graficar(lista_ordenada):
    # Crear 2 listas para colocar en los ejes x e y
    lista_topicos = []
    cantidad_topicos = []

    for elemento in itertools.islice(lista_ordenada, 20):
        lista_topicos.append(elemento[0])
        cantidad_topicos.append(elemento[1])

    plt.rcParams["figure.figsize"] = [10.00, 3.50]
    plt.rcParams["figure.autolayout"] = True
    # Etiquetar ejes
    plt.xlabel('Tópicos')
    plt.ylabel('Cantidad')
    # Rotar elementos de los ejes
    plt.xticks(rotation=90)

    plt.bar(lista_topicos, cantidad_topicos)
    plt.show()
