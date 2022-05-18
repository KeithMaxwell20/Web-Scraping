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
$ cd Go
$ go mod init Go
$ go mod tidy
$ go mod install .
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