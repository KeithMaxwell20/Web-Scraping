
package main.java;

import java.util.Scanner;
import tema1.Tema1;
import tema2.Tema2;

/**
 *
 *
 */
public class Java {

    public static void main(String[] args) {
        
        System.out.println("Tema: ");
        Scanner sc = new Scanner(System.in);
        String op = sc.nextLine();
        //String op = "2";
        switch (op) {
            case "1":
                Tema1 a = new Tema1(); 
                a.run("ResultadosTema1Java.csv", "../Resultados", "GraficoTema1GO.html", "../Resultados");
                break;
            case "2":
                Tema2 b = new Tema2();
                b.run("ResultadosTema2Java.csv", "../Resultados", "GraficoTema2Java.html", "../Resultados");
                break;
            default:
                System.err.println("Opcion invalida");
                break;
        }
    }
}
