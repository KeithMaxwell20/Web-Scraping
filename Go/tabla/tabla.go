package tabla

import "fmt"

func GenerarTabla(encabezados []string, datos [][]string) {

	fmt.Printf("\n\t")
	for i := 0; i < len(encabezados); i++ {
		fmt.Printf("|-----------------------------------------------------------")	
	}
	fmt.Printf("%s", "|")

	fmt.Printf("\n\t|")
	for i := 0; i < len(encabezados); i++ {
		fmt.Printf("%30s%30s",encabezados[i], "|")	
	}

	fmt.Printf("\n\t|")
	for i := 0; i < len(encabezados); i++ {
		fmt.Printf("-----------------------------------------------------------|")	
	}

	// Imprimiendo la lista de datos
	for i := 0; i < len(datos); i++ {
		
		fmt.Printf("\n\t|")
		for j := 0; j < len(datos[0]); j++ {
			fmt.Printf("%30s%30s",datos[i][j], "|")	
		}

		fmt.Printf("\n\t|")
		for j := 0; j < len(datos[0]); j++ {
			fmt.Printf("-----------------------------------------------------------|")			
		}
	}
	fmt.Printf("\n")
}