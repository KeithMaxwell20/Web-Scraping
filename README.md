# Web-Scraping

## Para ejecutar en Go

### Requerimientos
- Tener instalados los archivos binarios del compilador de go
    
    (Puede ser descargado de la web oficial [aquí](https://go.dev/dl/)).
- Mantener la configuración estándar de entorno

    (Las variables GOROOT y GOBIN apuntan a los directorios por defecto)

### Ejecución
#### Para la inicialización del programa
```
cd Go
go mod init Go
go mod tidy
go mod install .
```
#### Para ejecutar:
```
go run main.go X
```

(Donde X es el nro de ejercicio)
#### Para el tema 1:
```
go run main.go 1
```
#### Para el tema 2:
```
go run main.go 2
```

Obs.: Inicialmente desde el directorio raíz

## Para ejecutar en JavaScript

### Prerrequisitos

#### Se debe tener instalado los siguientes módulos:

- **NodeJs** 
- **npm**

Para la instalación de _Node.js_ y _npm_ en Windows, macOS y Linux se recomienda el siguiente enlace: [kinsta](https://kinsta.com/es/blog/como-instalar-node-js/#cmo-instalar-nodejs-y-npm)

#### Instalación de la librería puppeteer

Dentro del directorio **JavaScript** se debe instalar la librería puppeteer con el siguiente comando:
```
npm install puppeteer
```

### Ejecución

#### **Para el tema 1**

Dentro de la consola hay que ubicarse dentro del directorio del tema 1:
```
Web-Scraping/JavaScript/Tema 1/
```

Y posteriormente ejecutar el siguiente comando:
```
node app.js
```

Una vez terminada la ejecución se abre el archivo **index.html** corriendo sobre un servidor.

Si se utiliza VS Code se recomienda instalar la extensión: **Live Server**

#### **Para el tema 2**

Dentro de la consola hay que ubicarse dentro del directorio del tema 1:
```
Web-Scraping/JavaScript/Tema 2/
```

Y posteriormente ejecutar el siguiente comando:
```
node app.js
```

Una vez terminada la ejecución se abre el archivo **index.html** corriendo sobre un servidor.

Si se utiliza VS Code se recomienda instalar la extensión: **Live Server**

## Para ejecutar en Java

### Requerimientos
- Tener instalado el entorno de ejecución de Java (versión openjdk 11.0.9. o superiores)

- Tener instalado la última versión del IDE Apache Netbeans (versión 13)

    (Puede ser descargado de la web oficial [aquí](https://netbeans.apache.org/download/index.html)).

### Limpiar y Compilar
- Abrir el proyecto Maven en Netbeans (directorio Java)
- Seleccionar la opción de Limpiar y Compilar
#### Para ejecutar:
- Seleccionar la opción de Run (en el IDE Netbeans)


El programa pregunta en consola cuál programa ejecutar:

#### Para el tema 1:
- Ingresar 1 y ENTER
#### Para el tema 2:
- Ingresar 2 y ENTER
