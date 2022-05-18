package tema2;


import com.gargoylesoftware.htmlunit.WebClient;
import com.gargoylesoftware.htmlunit.html.DomNode;
import com.gargoylesoftware.htmlunit.html.DomNodeList;
import com.gargoylesoftware.htmlunit.html.HtmlPage;
import java.io.IOException;
import java.util.ArrayList;
import java.util.Collections;
import java.util.Comparator;
import java.util.logging.Level;
import java.util.logging.Logger;
import lib.Grafico;
import lib.Tabla;
import org.w3c.dom.Node;
import tema1.Tema1;

import java.time.LocalDateTime;



/**
 *
 *
 */
public class Tema2 {

    private class Resultados {
        String horaUltimaAct = "";
        ArrayList<String> listaTopics = new ArrayList(); 
    }
    
    private class Registro {
        String nombreTopic;
        Integer cantidad;
        
        Registro(String nombreTopic, Integer cantidad) {
            this.nombreTopic = nombreTopic;
            this.cantidad = cantidad;
        }
    }
    
    private int cantRepeticiones = 33;
    
    public void run(String nombreArchivo1, String directorioArchivo1, String nombreArchivo2, String directorioArchivo2) {
        System.out.println("Iniciando Programa 2...");
        String localFileName = "data.html";
        
        // Para analizar 
//        String link = "https://github.com/topics/nodejs?o=desc&s=updated";
        String link = "https://github.com/topics/python?o=desc&s=updated";
        System.out.println("Procesando.\nEste proceso puede tomar un tiempo...");
        
        // Generando el archivo local con los articulos
        ArrayList<Resultados> listaArticulos = procesar(link);

        System.out.println("Cantidad de Articulos: " + listaArticulos.size());

        // Creamos una lista de los topics y cantidad de apariciones
        ArrayList<Registro> listaResultados = cargarListaRegistro(listaArticulos);
        
        // Ordenamos la lista
        Collections.sort(listaResultados, new Comparator<Registro>() {
            @Override
            public int compare(Registro reg1, Registro reg2) {
            return reg1.cantidad > reg2.cantidad ? -1 : (reg1.cantidad < reg2.cantidad) ? 1 : 0;
        }});

        // Encabezados        
        String[] encabezado = {"TOPIC", "MENCIONES"};
        
        // Datos de la tabla
        String[][] matrizResultados = new String[listaResultados.size()][encabezado.length];        
	for (int i = 0; i < listaResultados.size(); i++) {
            matrizResultados[i][0] = listaResultados.get(i).nombreTopic;
            matrizResultados[i][1] = listaResultados.get(i).cantidad.toString();
	}
        Tabla t = new Tabla();
        t.generarTabla(encabezado, matrizResultados);
        
        // Generando los datos para el grafico de barras de los 20
        // temas con mayor cantidad de menciones.
        String[] info = {"Temas relacionados con Nodejs", "Los 20 topics mas mencionados", "TOPIC", "MENCIONES"};
        
        String[][] datos = new String[20][3];
        for(int i = 0; i < datos.length; i++) {
            datos[i][0] = listaResultados.get(i).cantidad.toString();
            datos[i][1] = listaResultados.get(i).nombreTopic;
            datos[i][2] = "";
        }
        Grafico g = new Grafico(info, datos);
        g.Run();
        
    }
    
    /***
     * Recibe una lista Resultados y retorna el nombre del topic y la cantidad
     * de apariciones en una lista de Registros
     */
    private ArrayList<Registro> cargarListaRegistro(ArrayList<Resultados> lista) {
        ArrayList<Registro> ret = new ArrayList<Registro>();
            
        for(int i = 0; i < lista.size(); i++) {
            for(int j = 0; j < lista.get(i).listaTopics.size(); j++) {
                boolean existe = false;
                int pos = -1;
                String auxTopic = lista.get(i).listaTopics.get(j);
                String aaaaa = auxTopic;
                for(int k = 0; k < ret.size(); k++) {
                    if(ret.get(k).nombreTopic.equals(auxTopic)) {
                        existe = true;
                        Integer cant = ret.get(k).cantidad + 1;
                        ret.set(k, new Registro(auxTopic, cant));
                        break;
                    }
                }
                if(existe==false) {
                    Registro a = new Registro(auxTopic, 1);
                    ret.add(a);
                }
            }
        }
        return ret;
    }
    
    /***
     * Visitando la web y extrayendo los datos en una lista de Resultados
     */
    private ArrayList<Resultados> procesar(String link) {
        ArrayList<Resultados> lista = new ArrayList<Resultados>();
        boolean continuar = true;
        
        // Creamos el objeto con el tiempo actual.
        LocalDateTime tiempoActual = LocalDateTime.now();
        // Definimos el tiempo que sirve de limite (hace 30 dias)
        LocalDateTime tiempoLimite = tiempoActual.plusDays(-30);
        
        continuar = procesarPagina(lista, link, tiempoLimite);
        for(int i = 2; i < cantRepeticiones && continuar; i++) {
            System.out.printf("... ");
            continuar = procesarPagina(lista, link + "&page=" + Integer.toString(i), tiempoLimite);
            try {
                Thread.sleep(10000);
            } catch (InterruptedException ex) {
                Logger.getLogger(Tema2.class.getName()).log(Level.SEVERE, null, ex);
            }
        }
        System.out.println("");
        return lista;
    }
    
    private boolean procesarPagina (ArrayList<Resultados> lista, String link, LocalDateTime tiempoLimite) {
        boolean continuar = true;

        // Iniciamos el cliente web
        WebClient webClient = new WebClient();
        
        // remove log
        java.util.logging.Logger.getLogger("com.gargoylesoftware").setLevel(Level.OFF); 
        System.setProperty("org.apache.commons.logging.Log", "org.apache.commons.logging.impl.NoOpLog");
        
        //configuring options    
        webClient.getOptions().setUseInsecureSSL(true);
        webClient.getOptions().setCssEnabled(true);
        webClient.getOptions().setThrowExceptionOnFailingStatusCode(false);
        webClient.getOptions().setThrowExceptionOnScriptError(false);
        webClient.getOptions().setJavaScriptEnabled(true);
        HtmlPage page = null;
        try {
            page = webClient.getPage(link);
            
        } catch (IOException ex) {
            Logger.getLogger(Tema1.class.getName()).log(Level.SEVERE, null, ex);
        }

        // Hacemos una lista de los nodos de articulos extraidos
        DomNodeList<DomNode> listaNodos = page.querySelectorAll("article.border.rounded.color-shadow-small.color-bg-subtle.my-4");
        
        // Extramos los datos (hora de actualizacion y topics) de cada
        // articulo extraido.
        for(int i = 0; i < listaNodos.size(); i++) {
            Resultados registro = new Resultados();            
            // Extraemos la hora en el formato RFC3339
            DomNode aux = listaNodos.get(i).querySelector("relative-time.no-wrap");
            Node time = aux.getAttributes().getNamedItem("datetime");

            // Guardamos el registro de la hora en que se actualiza.
            registro.horaUltimaAct = time.getNodeValue();
            // Extramos la lista de topics para guardar en un string[]
            DomNodeList<DomNode> auxTopics = listaNodos.get(i).querySelectorAll("a.topic-tag.topic-tag-link.f6.mb-2");

            
            // Aca debe estar el validador de fecha
            LocalDateTime tiempoArticulo = LocalDateTime.parse(time.getNodeValue().substring(0, time.getNodeValue().length()-1));
            // Restando 4 horas obtenemos la zona horaria paraguaya
            tiempoArticulo = tiempoArticulo.plusHours(-4);
            
            
            // Si la lista estaba vacia, significa que se ignora el articulo
            // porque el articulo contiene un Issue, no una actualizacion.
            if(auxTopics.size() != 0) {
                // Verificamos si el tiempo del articulo figura que sucedio 
                // antes que el tiempo limite
                if(tiempoArticulo.isBefore(tiempoLimite)) {
                    continuar = false;
                    break;
                }
                registro.listaTopics = new ArrayList<String>();
                for(int j = 0; j < auxTopics.size(); j++) {
                    DomNode aux2 = auxTopics.get(j);
                    String topic = aux2.getTextContent();
                    registro.listaTopics.add(topic.trim());
                }
                lista.add(registro);
            }
    
        }

        return continuar;
    }
    
    
    
}
