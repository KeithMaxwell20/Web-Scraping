import requests #Para solicitar el HTML de la página
from bs4 import BeautifulSoup #Herramientas de Web Scraping
import re #Para expresiones regulares


'''Este es un programa que utiliza Web Scraping para analizar y comparar la
popularidad de distintos lenguajes de programación según el índice Tiobe'''

# Definimos una lista con los 20 lenguajes más populares según TIOBE
lista_populares = ["Python", "C", "Java", "C++", "C#", "Visual Basic", "JavaScript",
                   "Assembly", "SQL", "PHP", "R", "R", "Delphi", "Go", "Swift",
                   "Ruby", "Classic Visual Basic", "Objective-C", "Perl", "Lua",
                   "MATLAB"]

# Lenguajes por su URL en Github
lista_topic = ["python", "c", "java", "cpp", "csharp", "visual-basic", "javascript",
		       "assembly", "sql", "php", "r", "object-pascal", "go", "swift", "ruby", "visual-basic", "objective-c",
		       "perl", "lua", "matlab"]


# 1.1 Extraer el número de repositorios de los lenguajes populares

# Dirección de la página
url = "https://github.com/topics"

# Lista donde guardaremos los valores numéricos hallados
lista_valores = []

for topico in lista_topic:
    url_busqueda = url + "/" + topico #Url de un lenguaje dado
    pagina = requests.get(url_busqueda) #Recibe el html de la página

    soup = BeautifulSoup(pagina.content, "html.parser")
    # Buscamos la clase que contiene el texto que necesitamos y convertimos a string
    resultado = soup.find(class_="h3 color-fg-muted")
    texto = resultado.text

    # El texto contiene cadenas y un valor numérico: extraemos el valor numérico
    valor = re.findall("\d+\,\d+", texto)

    # Convertimos el valor a un número entero
    valor[0] = valor[0].replace(',', '')
    valor_num = int(valor[0])
    print(topico, valor_num)






