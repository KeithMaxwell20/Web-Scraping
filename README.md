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