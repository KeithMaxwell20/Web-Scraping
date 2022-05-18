const puppeteer = require('puppeteer'); //Se llama a la librería puppeteer
const fs = require('fs'); //Se llama al módulo para interactuar con los archivos
const { match } = require('assert');
const path = require("path");
const { Console } = require('console');

//url de las rutas para trabajar con archivos
const urlGitHub = "https://github.com/topics/";
const urlTiobeList = "../../tiobe-list.csv";
const urlResult = "../../Resultado/";

//Array para los nombres de los arrays de objetos
const headerTiobeList = ['NOMBRE_LENGUAJE', 'TOPIC'];
const headerRating = ['NOMBRE_LENGUAJE', 'RATING_GITHUB', 'NRO_APARICIONES'];

main(); // Funcion que inicia el proceso

async function main() {
    // Se lee el archivos de la lista de tiobe, recibe un string
    var data = fs.readFileSync(urlTiobeList, 'utf8');

    // Primero se procesa el string leido del archivos, se retorna como array
    // Luego se realiza el scraping con el array obtenido
    // Por ultimo scraping retorna un array con el nombre y apariciones de los topicos
    var arrayTopicsMatching = await scraping(strToArray(data));

    // Se escribe en un archivos los resultados
    // Se verifica que exista el directorio y que sino se cree
    if (stateDirectory()) {
        fs.writeFileSync(urlResult + "ResultadosJavaScriptTema1.csv", arrayToStr(arrayTopicsMatching));
    }

    // Se obtiene el valor min y max para el calculo del Rating
    var min = Math.min.apply(Math, arrayTopicsMatching.map(function (e) { return e.NRO_APARICIONES; }));
    var max = Math.max.apply(Math, arrayTopicsMatching.map(function (e) { return e.NRO_APARICIONES; }));

    // Se obtiene un array ordenado segiun el rating
    var arrayRatingGitHub = getRatingGitHub(min, max, arrayTopicsMatching);

    // Se imprime por consola en rating en forma de tabla
    console.table(arrayRatingGitHub);

    // Se obtiene todos los nombres de los lenguajes de array ordenado por Rating
    let NOMBRE_LENGUAJE = arrayRatingGitHub.map(function (obj) {
        return obj.NOMBRE_LENGUAJE;
    });

    // Se obtiene todas las apariciones de los lenguajes de array ordenado por Rating
    let NRO_APARICIONES = arrayRatingGitHub.map(function (obj) {
        return obj.NRO_APARICIONES;
    });

    // Se eliminan los ultimos 10, ya que se tienen 20 y solo pide 10
    for (let i = 0; i < 10; i++) {
        NOMBRE_LENGUAJE.pop();
        NRO_APARICIONES.pop();
    }

    // Se crea un objeto que posee dos array para guardar en el JSON, esto es para la gráfica
    data = {
        'NOMBRE_LENGUAJE': NOMBRE_LENGUAJE,
        'NRO_APARICIONES': NRO_APARICIONES
    };

    // Se escribe el objeto en un archivo JSON en formato JSON para poder recuperar luego
    fs.writeFileSync("data.json", JSON.stringify(data));
}


async function scraping(arrayTopics) {
    const browser = await puppeteer.launch(); // Lanzamos un nuevo navegador
    const arrayTopicsMatching = []; // Array que contendrá las apariciones por lenguaje

    // Se recorre los topics para usar como URL del navegador
    for (var index = 0; index < arrayTopics.length; index++) {
        let element = arrayTopics[index];

        const page = await browser.newPage(); // Abrimos una nueva página
        await page.goto(urlGitHub + element.TOPIC); // Vamos a la URL
        await page.waitForTimeout(3000);

        // Se obtiene el numero de apariciones para el topic
        const topicMatching = await page.evaluate(() => {
            const tmp = document.querySelector('.col-md-8 h2').innerText;
            return tmp.split(" ")[2].replace(",", "");
        });

        // Se almacena en un objeto el nombre y nro de apariciones
        let topic = {};
        topic[headerRating[0]] = element.NOMBRE_LENGUAJE;
        topic[headerRating[2]] = topicMatching;

        arrayTopicsMatching.push(topic); // Se agrega al array

        await page.close(); // Cerramos la página

        //Mensajes del procedimiento
        console.log("Procedimiento exitoso: " + urlGitHub + element.TOPIC);
        console.log("NAME: " + element.NOMBRE_LENGUAJE + " | MATCHING: " + topicMatching + "\n");
    }
    await browser.close(); // Cerramos el navegador

    return arrayTopicsMatching;
}

//Convierte a un array el string obtenido del archivo de la lista de tiobe 
function strToArray(str, delimiter = ",") {
    //Se omite la primela linea del archivo que contiene solo titulos
    let deleteOneLine = str.replace("\r", "").slice(0, str.indexOf("\n")).split(delimiter);

    // Se corta desde \n índice + 1 hasta el final del texto
    // Split para crear una matriz de cada fila de valor csv
    str.replace("\r", "");
    const rows = str.slice(str.indexOf("\n") + 1).split("\n");

    //  Crear una matriz de objetos a partir de la matriz de valores csv
    let arrayTopic = [];
    rows.forEach(row => {
        let values = row.split(delimiter); // Se divide en dos el string para el nombre y el topic

        // Se guarda en un objeto para tener el array de objetos
        let element = {};
        element[headerTiobeList[0]] = values[0];
        element[headerTiobeList[1]] = values[1];

        // Se agrega el objeto al array
        arrayTopic.push(element);
    });

    arrayTopic.pop(); //Se elimina la ultima linea vacia del archivo

    return arrayTopic;
}

// Función que convierte un array a string para guardar en un archivo
function arrayToStr(array, delimiter = ',') {
    var str = "";

    array.forEach(function (element) {
        str = str + element.NOMBRE_LENGUAJE + delimiter + element.NRO_APARICIONES + "\n";
    });

    return str;
}

// Funcion que verifica si exsite el directorio de donde se creara el archivo con los Resultados
// En caso que no exista crea el directorio
function stateDirectory() {
    let state = true;

    if (fs.existsSync(urlResult)) {
        console.log("El directorio ya ha sido creado");
    } else {
        fs.mkdir(urlResult, (error) => {
            if (error) {
                console.log(error.message);
                state = false;
            }
            console.log("Directorio creado con exito");
        });
    }

    return state;
}

// Funcion que calcula el rating de apariciones
function getRatingGitHub(min, max, array) {
    // Array que contendra el rating
    let rating = [];

    // Se recorre el array para tener las apariciones
    for (let index = 0; index < array.length; index++) {
        //Se guardar el nombre, nro apariciones y el rating por eso se usa un objeto
        let element = {};

        element[headerRating[0]] = array[index].NOMBRE_LENGUAJE; // Se guarda el nombre del lenguaje
        let number = (array[index].NRO_APARICIONES - min) / (max - min) * 100 // Se calcula su rating
        element[headerRating[1]] = number.toFixed(2);
        element[headerRating[2]] = array[index].NRO_APARICIONES; // Se guarda su numero de apariciones

        rating.push(element); // Se agrega el objeto al array
    }

    // Se ordena el array de forma descendente
    rating.sort(function (a, b) {
        return b.RATING_GITHUB - a.RATING_GITHUB;
    });

    return rating;
}
