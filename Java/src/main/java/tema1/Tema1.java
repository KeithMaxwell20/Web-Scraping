package tema1;



import com.gargoylesoftware.htmlunit.html.HtmlPage;
import java.io.*;  
import java.util.ArrayList;
import java.util.Scanner;  
import java.util.logging.Level;
import java.util.logging.Logger;
import java.io.IOException;
import com.gargoylesoftware.htmlunit.WebClient;
import com.gargoylesoftware.htmlunit.html.DomNode;
import java.text.DecimalFormat;
import java.text.NumberFormat;
import lib.Grafico;
import lib.Tabla;

public class Tema1 {
    
    private class NombreTopic {
        String nombre = "";
        String topic = "";
        
        NombreTopic(String nombre, String topic) {
            this.nombre = nombre;
            this.topic = topic;
        }
    }
    
    private class InfoLenguaje {
	String nombre;
	Float rating;
	int cantidad;
        
        InfoLenguaje(String nombre, Float rating, int cantidad) {
            this.nombre = nombre;
            this.rating = rating;
            this.cantidad = cantidad;
        }
    }
    
    /**
     * Ejecuta  los metodos para el tema 1
     * @param nombreArchivo1 Nombre del archivo .csv para el item 1.2
     * @param directorioArchivo1 Directorio del archivo 1
     * @param nombreArchivo2 Nombre del archivo .html para el gráfico del item 1.4
     * @param directorioArchivo2 Directorio del archivo 2
     */
    public void run(String nombreArchivo1, String directorioArchivo1, String nombreArchivo2, String directorioArchivo2) {
        System.err.println("Iniciando Programa 1"); 
        String archivoEntrada = ".././tiobe-list.csv";
        ArrayList<NombreTopic> listaEntrada = new ArrayList<NombreTopic>();
        
        
        try {
            listaEntrada = extraerDatosEntrada(archivoEntrada);
        }catch (FileNotFoundException e) {
            System.err.println(e);
            System.exit(0);
        }
        crearDirectorio(directorioArchivo1);
        
        // Generamos el archivo de salida para guardar los resultados
        String encoding = "UTF-8";
        PrintWriter archivoSalida = null;
        try{
            archivoSalida = new PrintWriter(directorioArchivo1 + "/" + nombreArchivo1, encoding);
        } catch (IOException e){ 
            System.out.println("An error occurred.");
            e.printStackTrace();
        }        
        
        // Guardamos los resultados del scraping
        ArrayList<InfoLenguaje> listaResultados = new ArrayList<InfoLenguaje>();
        for(int i = 0; i < listaEntrada.size(); i++) {
            listaResultados.add(procesar(listaEntrada.get(i).nombre, listaEntrada.get(i).topic, archivoSalida));
        }

        // Calcular Rating (ítem 1.3)
        calcularRating(listaResultados);
        // Ordenamos la lista
        ordenarLista(listaResultados);
        
        // Preparamos el formato de valores requerido para el objeto Tabla
        // Encabezados
        String[] encabezado = {"LENGUAJE", "RATING", "REPOSITORIOS"};
        // Lista de Lenguajes con su rating
        String[][] matrizResultados = new String[listaResultados.size()][encabezado.length];
            NumberFormat formato = new DecimalFormat("0.000");
        for (int i = 0; i < listaResultados.size(); i++) {
                matrizResultados[i][0] = listaResultados.get(i).nombre;
                matrizResultados[i][1] = formato.format(listaResultados.get(i).rating);
                matrizResultados[i][2] = Integer.toString(listaResultados.get(i).cantidad);
      	}

        // Imprimiendo resultados (item 1.4)
        System.out.println("Generando Tabla: ");
        Tabla t = new Tabla();
        t.generarTabla(encabezado, matrizResultados);
        
        // Cerrando archivo
        if(archivoSalida != null) archivoSalida.close();
        
        // Generando los datos para el grafico de barras de los 10
        // lenguajes con la mayor cantidad de repositorios
        String[] info = {"Lenguajes con Mayor Nº de Apariciones", "Los 10 primeros lenguajes, ordenados de mayor a menor", "LENGUAJE", "Nº de Apariciones"};
        
        String[][] datos = new String[10][3];
        for(int i = 0; i < datos.length; i++) {
            datos[i][0] = ((Integer)listaResultados.get(i).cantidad).toString();
            datos[i][1] = listaResultados.get(i).nombre;
            datos[i][2] = "";
        }
        Grafico g = new Grafico(info, datos);
        g.Run();
    }


    /***
     * Calcula el rating de cada lenguaje, de acuerdo a la formula dada
     */
    private void calcularRating(ArrayList<InfoLenguaje> lista) {
        Integer max = maximo(lista);
	Integer min = minimo(lista);
	for (int i = 0; i < lista.size(); i++) {
            lista.get(i).rating =  ((float) (lista.get(i).cantidad-min)/(max-min)) * 100f;
	}
    }

    /***
     * Retorna el valor maximo de cantidad
     */
    private Integer maximo(ArrayList<InfoLenguaje> lista) {
        Integer max = 0;
        for(int i = 0; i < lista.size(); i++) {
            if(lista.get(i).cantidad > max) max = lista.get(i).cantidad;
        }
        return max;
    }

    /***
     * Retorna el valor minimo de cantidad
     */
    private Integer minimo(ArrayList<InfoLenguaje> lista) {
        Integer min = Integer.MAX_VALUE;
        for(int i = 0; i < lista.size(); i++) {
            if(lista.get(i).cantidad < min) min = lista.get(i).cantidad;
        }
        return min;
    }


    
    /***
     * Ordenar la lista en forma descendente, de acuerdo al atributo
     * cantidad.
     */
    private void ordenarLista(ArrayList<InfoLenguaje> lista) {
        for (int i = 0; i < lista.size()-1; i++) {
            for (int j = i + 1; j < lista.size(); j++) {
                if (lista.get(i).cantidad < lista.get(j).cantidad) {
                    InfoLenguaje aux = new InfoLenguaje(lista.get(i).nombre, lista.get(i).rating, lista.get(i).cantidad);
                    lista.get(i).cantidad = lista.get(j).cantidad;
                    lista.get(i).nombre = lista.get(j).nombre;
                    lista.get(i).rating = lista.get(j).rating;    
                    lista.get(j).nombre = aux.nombre;
                    lista.get(j).rating = aux.rating;
                    lista.get(j).cantidad = aux.cantidad;
                }
            }
        }
    }
    
    /***
     * Visita el website "https://github.com/topics/" y extrae el nro.
     * de repositorios para el topic recibido.
     */
    private InfoLenguaje procesar(String nombre, String topic, PrintWriter archivo) {
        
        System.out.println("Extrayendo resultados para " + nombre);
        InfoLenguaje registroRetorno = null;
        String link = "https://github.com/topics/" + topic;
        // Iniciamos el cliente web
        WebClient webClient = new WebClient();
        
        // remove log
        java.util.logging.Logger.getLogger("com.gargoylesoftware").setLevel(Level.OFF); 
        System.setProperty("org.apache.commons.logging.Log", "org.apache.commons.logging.impl.NoOpLog");
        
        //configuring options    
        webClient.getOptions().setUseInsecureSSL(true);
        webClient.getOptions().setCssEnabled(false);
        //webClient.getOptions().setJavaScriptEnabled(false);        
        webClient.getOptions().setThrowExceptionOnFailingStatusCode(false);
        webClient.getOptions().setThrowExceptionOnScriptError(false);
        HtmlPage page = null;
        try {
            page = webClient.getPage(link);
        } catch (IOException ex) {
            Logger.getLogger(Tema1.class.getName()).log(Level.SEVERE, null, ex);
        }
        DomNode a = page.querySelector("h2.h3.color-fg-muted"); 
        String repositorios = extraerCantRep(a.getTextContent());
        archivo.println(nombre + "," + repositorios);
        registroRetorno = new InfoLenguaje(nombre, 0.0f, Integer.parseInt(repositorios));
        return registroRetorno;
    }
    
    /***
     * Retorna la concatenacion de los nros. en texto
     */
    private String extraerCantRep(String texto) {
        String retorno = "";
        for(int i = 0; i < texto.length(); i++) {
            if(Character.isDigit(texto.charAt(i))) {
                retorno = retorno.concat(Character.toString(texto.charAt(i)));
            }
        }    
        return retorno;
    }
    
    /**
     * Guarda datos del archivo.csv en un lista a retornar
     * 
     */
    private ArrayList<NombreTopic> extraerDatosEntrada(String archivo) throws FileNotFoundException {
        ArrayList<NombreTopic> listaRetorno = new ArrayList<NombreTopic>();
        Scanner sc = new Scanner(new File(archivo));
        // Descartamos el encabezado
        sc.nextLine();

        String auxNombre, auxTopic, line;
        while(sc.hasNext()) {
            line = sc.nextLine();
            auxNombre = line.substring(0, line.indexOf(","));
            auxTopic = line.substring(line.indexOf(",")+1);
            listaRetorno.add(new NombreTopic(auxNombre, auxTopic));
        }
        return listaRetorno;
    }
    
    /***
     * Crea directorio en la direccion recibida
     */
    private void crearDirectorio(String directorio) {
        File a = new File(directorio);
        if (a.exists()) {
            if (a.isDirectory()) 
                System.out.println("Existe el directorio");
        } else {
            a.mkdir();
        }
    }
}
