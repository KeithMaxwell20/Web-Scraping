
package lib;

/**
 * Genera salida en consola con formato tabla.
 */
public class Tabla {

	/**
	 * Genera la tabla en consola 
	 * @param encabezados Nombres de las columnas
 	 * @datos datos a imprimir en cada fila
	 * 
	*/
    public void generarTabla(String[] encabezados, String[][] datos) {
        System.out.printf("\n\t");
	for (int i = 0; i < encabezados.length; i++) {
		System.out.printf("|-----------------------------------------------------------");	
	}
	System.out.printf("%s", "|");

	System.out.printf("\n\t|");
	for (int i = 0; i < encabezados.length; i++) {
		System.out.printf("%30s%30s",encabezados[i], "|");	
	}

	System.out.printf("\n\t|");
	for (int i = 0; i < encabezados.length; i++) {
		System.out.printf("-----------------------------------------------------------|");	
	}

	// Imprimiendo la lista de datos
	for (int i = 0; i < datos.length; i++) {
		
		System.out.printf("\n\t|");
		for (int j = 0; j < datos[0].length; j++) {
			System.out.printf("%30s%30s",datos[i][j], "|");	
		}

		System.out.printf("\n\t|");
		for (int j = 0; j < datos[0].length; j++) {
			System.out.printf("-----------------------------------------------------------|");			
		}
	}
	System.out.printf("\n");
    }
}
