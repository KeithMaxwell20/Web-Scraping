import requests #Para solicitar el HTML de la página
from bs4 import BeautifulSoup #Herramientas de Web Scraping
import re #Para expresiones regulares
import os # Para verificar la existencia de un archivo
import csv # Para crear archivos del tipo csv


# Una función que retorne la lista con la cantidad de tópicos todos los lenguajes
def retornar_lista(url, lista_topic):
    # Lista donde guardaremos los valores numéricos hallados
    lista_valores = []
    print("Tomando cantidad de repositorios por cada tópico...")

    # Iteramos por cada tópico
    for topico in lista_topic:
        url_busqueda = url + "/" + topico #Url de un lenguaje dado
        pagina = requests.get(url_busqueda) #Recibe el html de la página

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
    escritor.writerow(lista_lenguajes)
    escritor.writerow(cantidad_repositorios)

    print("El archivo csv se ha creado con éxito")


'''Este es un programa que utiliza Web Scraping para analizar y comparar la
popularidad de distintos lenguajes de programación según el índice Tiobe'''

# Definimos una lista con los 20 lenguajes más populares según TIOBE
lista_populares = ["Python", "C", "Java", "C++", "C#", "Visual Basic", "JavaScript",
                   "Assembly", "SQL", "PHP", "R", "Delphi", "Go", "Swift",
                   "Ruby", "Classic Visual Basic", "Objective-C", "Perl", "Lua",
                   "MATLAB"]

# Lenguajes por su URL en Github
lista_topic = ["python", "c", "java", "cpp", "csharp", "visual-basic", "javascript",
		       "assembly", "sql", "php", "r", "delphi", "go", "swift", "ruby",
               "visual-basic", "objective-c", "perl", "lua", "matlab"]


# 1.1 Extraer el número de repositorios de los lenguajes populares
# Dirección de la página
url = "https://github.com/topics"

# Verificamos si el archivo con los tópicos y sus repositorios existe
if not os.path.exists('/resultados.csv'):
    # Recibimos la lista con los tópicos por cada lenguaje
    lista = retornar_lista(url, lista_topic)
    # 1.2 Creamos el archivo csv
    crear_csv(lista, lista_populares) # Se debe añadir nombres a las filas

#1.3





