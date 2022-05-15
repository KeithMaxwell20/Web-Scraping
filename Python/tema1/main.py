# Utils contiene la descripción de todas las funciones usadas en el main
import utils

# Este programa es una solución al ejercicio 1 planteado en el trabajo
# práctico de Web Scraping

# Definimos una lista con los 20 lenguajes más populares según TIOBE
lista_populares = ["Python", "C", "Java", "C++", "C#", "Visual Basic", "JS",
                   "Assembly", "SQL", "PHP", "R", "Delphi", "Go", "Swift",
                   "Ruby", "Classic Visual Basic", "Objective-C", "Perl", "Lua",
                   "MATLAB"]

# Lenguajes por su URL en github.com/topic/
lista_topic = ["python", "c", "java", "cpp", "csharp", "visual-basic", "javascript",
		       "assembly", "sql", "php", "r", "delphi", "go", "swift", "ruby",
               "visual-basic-for-applications", "objective-c", "perl", "lua", "matlab"]


# 1.1 Extraer el número de repositorios de los lenguajes populares
# Dirección de la página
url = "https://github.com/topics"
# Recibimos la lista con la cantidad de tópicos por los 20 lenguajes más populares
lista = utils.retornar_lista(url, lista_topic)


# 1.2 Creamos el archivo csv
utils.crear_csv(lista, lista_populares)


# 1.3 Hallar máximo y mínimo número de apariciones de los 20 lenguajes en github
minimo = min(lista)
maximo = max(lista)
# Calculamos el rating de github y lo colocamos en una lista
lista_rating = utils.retornar_rating(minimo, maximo, lista)


# 1.4 Ordenar descendentemente e imprimir
listas_ordenadas = utils.ordenar_listas(lista_rating, lista, lista_populares)
utils.imprimir_lista(listas_ordenadas)


# 1.5 Realizar un gráfico de barras con los 10 primeros lenguajes según
# número de apariciones en Github
utils.crear_grafico(listas_ordenadas)
