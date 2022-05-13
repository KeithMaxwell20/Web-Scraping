const puppeteer = require('puppeteer');
const fs = require('fs');
const { match } = require('assert');
const path = require("path");
const { Console } = require('console');

const urlGitHub = "https://github.com/topics/";
const urlTiobeList = "../tiobe-list.csv";
const urlResult = "../Resultado/";

const headerTiobeList = ['NOMBRE_LENGUAJE', 'TOPIC'];
const headerRating = ['NOMBRE_LENGUAJE', 'RATING_GITHUB', 'NRO_APARICIONES'];


async function main() {
    // put in a array the dato of the csv file
    var data = fs.readFileSync(urlTiobeList, 'utf8');

    var arrayTopicsMatching = await scraping(strToArray(data));
    if (stateDirectory()) {
        fs.writeFileSync(urlResult + "ResultadosJavaScript.csv", arrayToStr(arrayTopicsMatching));
    }

    var min = Math.min.apply(Math, arrayTopicsMatching.map(function (e) { return e.NRO_APARICIONES; }));
    var max = Math.max.apply(Math, arrayTopicsMatching.map(function (e) { return e.NRO_APARICIONES; }));

    var arrayRatingGitHub = getRatingGitHub(min, max, arrayTopicsMatching);

    console.table(arrayRatingGitHub);

    let NOMBRE_LENGUAJE = arrayRatingGitHub.map(function(obj) {
        return obj.NOMBRE_LENGUAJE;
    });

    let NRO_APARICIONES = arrayRatingGitHub.map(function(obj) {
        return obj.NRO_APARICIONES;
    });

    data = {
        'NOMBRE_LENGUAJE': NOMBRE_LENGUAJE,
        'NRO_APARICIONES': NRO_APARICIONES
    };

    fs.writeFileSync("data.json", JSON.stringify(data));
}


async function scraping(arrayTopics) {
    const browser = await puppeteer.launch(); // Lanzamos un nuevo navegador
    const arrayTopicsMatching = [];

    for (var index = 0; index < arrayTopics.length; index++) {
        let element = arrayTopics[index];

        const page = await browser.newPage(); // Abrimos una nueva página
        await page.goto(urlGitHub + element.TOPIC); // Vamos a la URL

        const topicMatching = await page.evaluate(() => {
            const tmp = document.querySelector('.col-md-8 h2').innerText;
            return tmp.split(" ")[2].replace(",", "");
        });

        let topic = {};
        topic[headerRating[0]] = element.NOMBRE_LENGUAJE;
        topic[headerRating[2]] = topicMatching;

        arrayTopicsMatching.push(topic);

        await page.close(); // Cerramos la página

        console.log("Procedimiento exitoso: " + urlGitHub + element.TOPIC);
        console.log("NAME: " + element.NOMBRE_LENGUAJE + " | MATCHING: " + topicMatching + "\n");
    }
    await browser.close(); // Cerramos el navegador

    return arrayTopicsMatching;
}

function strToArray(str, delimiter = ",") {
    let deleteOneLine = str.replace("\r", "").slice(0, str.indexOf("\n")).split(delimiter);

    // slice from \n index + 1 to the end of the text
    // use split to create an array of each csv value row
    str.replace("\r", "");

    const rows = str.slice(str.indexOf("\n") + 1).split("\n");

    // create an array of objects from the array of csv values
    let arrayTopic = [];
    rows.forEach(row => {
        let values = row.split(delimiter);
        let element = {};
        element[headerTiobeList[0]] = values[0];
        element[headerTiobeList[1]] = values[1];
        arrayTopic.push(element);
    });

    arrayTopic.pop(); //Se elimina la ultima linea vacia del archivo

    return arrayTopic;
}

function arrayToStr(array, delimiter = ',') {
    var str = "";

    array.forEach(function (element) {
        str = str + element.NOMBRE_LENGUAJE + delimiter + element.NRO_APARICIONES + "\n";
    });

    return str;
}


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

function getRatingGitHub(min, max, array) {
    let rating = [];
    for (let index = 0; index < array.length; index++) {
        let element = {};

        element[headerRating[0]] = array[index].NOMBRE_LENGUAJE;
        let number = (array[index].NRO_APARICIONES - min) / (max - min) * 100
        element[headerRating[1]] = number.toFixed(2);
        element[headerRating[2]] = array[index].NRO_APARICIONES;

        rating.push(element);
    }

    rating.sort(function (a, b) {
        return b.RATING_GITHUB - a.RATING_GITHUB;
    });

    return rating;
}

main();
