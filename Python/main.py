import requests #Para solicitar el HTML de la página
from bs4 import BeautifulSoup #Herramientas de Web Scraping
import re #Para expresiones regulares
import os # Para verificar la existencia de un archivo
import csv # Para crear archivos del tipo csv
import time # Para pausar la ejecución del programa
import matplotlib.pyplot as plt # Para graficar los resultados finales
import itertools # Para iterar una lista




# Una función que retorne la lista con la cantidad de tópicos todos los lenguajes
def retornar_lista(url, lista_topic):
    # Lista donde guardaremos los valores numéricos hallados
    lista_valores = []
    print("Tomando cantidad de repositorios por cada tópico...")

    # Iteramos por cada tópico
    for topico in lista_topic:
        url_busqueda = url + "/" + topico #Url de un lenguaje dado
        pagina = requests.get(url_busqueda) #Recibe el html de la página

        # Acá manejamos el error http 429 por el cual realizamos más peticiones de las permitidas en un rango de tiempo
        while pagina.status_code == 429:
            print("Saltó el código http 429, esperaremos", int(pagina.headers["Retry-after"]) ,"segundos antes de mandar más peticiones")
            time.sleep(int(pagina.headers["Retry-after"]))
            pagina = requests.get(url_busqueda)


        soup = BeautifulSoup(pagina.content, "html.parser")
        # Buscamos la clase que contiene el texto que necesitamos y convertimos a string
        resultado = soup.find(class_="h3 color-fg-muted")
        texto = resultado.text

        # El texto contiene cadenas y un valor numérico: extraemos el valor numérico
        if texto.find(',')!=-1:
            valor = re.findall("\d+\,\d+", texto)
            valor_hallado = valor[0]
            valor_hallado = valor_hallado.replace(',', '') #Quitamos la coma, si tiene
        else:
            valor = re.findall("\d+", texto)
            valor_hallado = valor[0]


        # Convertimos el valor hallado a entero
        valor_num = int(valor_hallado)

        lista_valores.append(valor_num)

    print("El proceso ha concluido con éxitos")
    return lista_valores

# Crear archivo csv para almacenar el nombre de los lenguajes y sus apariciones
def crear_csv(lista_lenguajes, cantidad_repositorios):

    print("Creando archivo csv...")

    archivo = open('Resultados.csv', 'w') #Crear el archivo en modo escritura
    escritor = csv.writer(archivo) #Escritor para el archivo csv

    # Escribimos los nombres de los lenguajes junto a su cantidad de repositorios

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
def ordenar_listas(lista_rating, lista_repositorios, lista_populares):
    # Ordenar lista
    #lista_ordenada = lista.sort(reverse=True)
    #Imprimir lista

    zipeado = zip(lista_rating, lista_repositorios, lista_populares)
    lista_zipeado = list(zipeado)

    lista_zipeado.sort(reverse=True)

    return lista_zipeado

def imprimir_lista(listas_ordenadas):
    print() # Genera un espacio en blanco
    print("Nombre del lenguaje, Número de Apariciones, Rating")

    for elemento in listas_ordenadas:
        print(elemento[2], elemento[1], elemento[0])


# Realizar gráfico de los 10 lenguajes con mayor número de apariciones
def crear_grafico(listas_ordenadas):

    # Crear 2 listas, una con los 10 lenguajes y otra con su nro de repositorios
    lista_lenguajes = []
    cantidad_repositorios = []

    for elemento in itertools.islice(listas_ordenadas, 10):
        lista_lenguajes.append(elemento[2])
        cantidad_repositorios.append(elemento[1])

    # Graficar los resultados
    plt.rcParams["figure.figsize"] = [7.00, 3.50]
    plt.rcParams["figure.autolayout"] = True

    plt.bar(lista_lenguajes, cantidad_repositorios)

    plt.show()


'''Este es un programa que utiliza Web Scraping para analizar y comparar la
popularidad de distintos lenguajes de programación según el índice Tiobe'''

# Definimos una lista con los 20 lenguajes más populares según TIOBE
lista_populares = ["Python", "C", "Java", "C++", "C#", "Visual Basic", "JS",
                   "Assembly", "SQL", "PHP", "R", "Delphi", "Go", "Swift",
                   "Ruby", "Classic Visual Basic", "Objective-C", "Perl", "Lua",
                   "MATLAB"]

# Lenguajes por su URL en Github
lista_topic = ["python", "c", "java", "cpp", "csharp", "visual-basic", "javascript",
		       "assembly", "sql", "php", "r", "delphi", "go", "swift", "ruby",
               "visual-basic-for-applications", "objective-c", "perl", "lua", "matlab"]


# 1.1 Extraer el número de repositorios de los lenguajes populares
# Dirección de la página
url = "https://github.com/topics"

# ¿Es necesario elegir entre leer el archivo y ejecutar las funciones?

# Verificamos si el archivo con los tópicos y sus repositorios existe
#if not os.path.exists('/resultados.csv'):
# Recibimos la lista con la cantidad de tópicos por cada lenguaje
lista = retornar_lista(url, lista_topic)
# 1.2 Creamos el archivo csv
crear_csv(lista, lista_populares)
#else: # Leemos el archivo csv para obtener la lista


#1.3 Hallar máximo y mínimo número de apariciones
# (puede ser convertido a una función que retorne los 2 valores)
minimo = min(lista)
maximo = max(lista)

#Calculamos el rating de github y lo colocamos en una lista
lista_rating = retornar_rating(minimo, maximo, lista)

# 1.4 Ordenar descendentemente e imprimir
listas_ordenadas = ordenar_listas(lista_rating, lista, lista_populares)
# Imprimir listas
imprimir_lista(listas_ordenadas)


# 1.5 Realizar un gráfico de barras con los 10 primeros lenguajes
crear_grafico(listas_ordenadas)

