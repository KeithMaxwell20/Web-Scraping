import requests #Para solicitar el HTML de la página
from bs4 import BeautifulSoup #Herramientas de Web Scraping
import re #Para expresiones regulares
import csv # Para crear archivos del tipo csv
import matplotlib.pyplot as plt # Para graficar los resultados finales
import itertools # Para iterar una lista
import time # Para pausar el proceso



# Retorna la cantidad de apariciones de un lenguaje en Github
def retornar_lista(url, lista_topic):
    # Lista donde guardaremos los valores numéricos hallados
    lista_valores = []
    print("Tomando cantidad de repositorios por cada tópico...")

    contador_pag = 0

    # Iteramos por los 20 lenguajes para hallar su cantidad de apariciones
    for topico in lista_topic:
        url_busqueda = url + "/" + topico #Url de un lenguaje dado
        pagina = requests.get(url_busqueda) #Recibe el html de la página

        contador_pag+=1
        print("Obteniendo página:", contador_pag)

        time.sleep(1) # Para prevenir que aparezca el error http 429

        # Maneja el error http 429
        # El encabezado Retry-after indica cuántos segundos debemos esperar
        # para mandar más solicitudes
        while pagina.status_code == 429:
            print("Saltó el código http 429, esperaremos", int(pagina.headers["Retry-after"]) ,"segundos antes de mandar más peticiones")
            time.sleep(int(pagina.headers["Retry-after"]))
            pagina = requests.get(url_busqueda)

        # Usamos soup para parsear el htmls
        soup = BeautifulSoup(pagina.content, "html.parser")
        # Buscamos la clase que contiene el texto que necesitamos
        resultado = soup.find(class_="h3 color-fg-muted")
        texto = resultado.text

        # El texto contiene cadenas y un valor numérico: extraemos el valor numérico
        # Para ello usamos expresiones regulares que nos permitan extraer este valor
        if texto.find(',')!=-1:
            valor = re.findall("\d+\,\d+", texto)
            valor_hallado = valor[0]
            valor_hallado = valor_hallado.replace(',', '') # Quitamos la coma
        else: # Si el número no tiene coma, usamos una expresión regular diferente
            valor = re.findall("\d+", texto)
            valor_hallado = valor[0]


        # Convertimos el valor hallado a entero y añadimos a la lista
        valor_num = int(valor_hallado)

        lista_valores.append(valor_num)

    print("El proceso ha concluido con éxito")
    return lista_valores


# Crear archivo csv para almacenar el nombre de los lenguajes y sus apariciones
def crear_csv(lista_lenguajes, cantidad_repositorios):
    print()
    print("Creando archivo csv...")

    archivo = open('Resultados.csv', 'w') #Crear el archivo en modo escritura
    escritor = csv.writer(archivo) # Escritor para el archivo csv

    # Escribimos los nombres de los lenguajes junto a su cantidad de repositorios
    # cada uno en una fila distinta
    escritor.writerow(cantidad_repositorios)
    escritor.writerow(lista_lenguajes)

    print("El archivo csv se ha creado con éxito")


# Retornar una lista con el rating de los lenguajes de programación
def retornar_rating(min, max, cantidad_repositorios):
    lista_rat = []
    for elemento in cantidad_repositorios:
        rating = (elemento - min)*100/(max-min)
        lista_rat.append(rating)

    return lista_rat


# Ordenar lista según rating, número de tópicos y lenguaje
# La función zip une las 3 listas en 1 sola
def ordenar_listas(lista_rating, lista_repositorios, lista_populares):
    zipeado = zip(lista_rating, lista_repositorios, lista_populares)
    lista_zipeado = list(zipeado)

    lista_zipeado.sort(reverse=True)

    return lista_zipeado


def imprimir_lista(listas_ordenadas):
    print() # Genera un espacio en blanco
    print("Nombre del lenguaje, Número de Apariciones, Rating")

    for elemento in listas_ordenadas:
        print(elemento[2], "-", elemento[1], "-", elemento[0])


# Realizar gráfico de los 10 lenguajes con mayor número de apariciones
def crear_grafico(listas_ordenadas):

    # Crear 2 listas, una con los 10 lenguajes y otra con su nro de apariciones
    lista_lenguajes = []
    cantidad_apariciones = []

    # islice facilita la iteración de la lista
    for elemento in itertools.islice(listas_ordenadas, 10):
        lista_lenguajes.append(elemento[2])
        cantidad_apariciones.append(elemento[1])

    # Graficar los resultados
    plt.rcParams["figure.figsize"] = [7.00, 3.50]
    plt.rcParams["figure.autolayout"] = True

    # Etiquetar ejes x e y
    plt.xlabel('Lenguajes')
    plt.ylabel('Apariciones')

    plt.bar(lista_lenguajes, cantidad_apariciones)
    plt.show()


