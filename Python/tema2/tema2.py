import utils

'''Función main del programa'''

# El tópico de mi interés es go
url = "https://github.com/topics/go?o=desc&s=updated"

# Obtener la página con todos los "load more"
pagina = utils.obtenerPaginas(url)

# Obtener la cantidad total de artículos actualizados en los últimos 30 días
contador_art = utils.obtenerTotalArticulos(pagina)

# Obtiene la lista total de todos los tópicos ordenada
lista_ordenada = utils.obtenerListaTopics(pagina, contador_art)

# Crea un archivo csv donde guarda los resultados
utils.crear_csv(lista_ordenada)

utils.graficar(lista_ordenada)

