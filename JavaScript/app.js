const puppeteer = require('puppeteer');
const fs = require('fs');
const { match } = require('assert');
const path = require("path");

const urlGitHub = "https://github.com/topics/";
const urlTiobeList = "../tiobe-list.csv";
const urlResult = "../Resultado/";

const headers = ['name', 'topic'];

(async () => {
    const browser = await puppeteer.launch();
    const page = await browser.newPage();

    // put in a array the dato of the csv file
    var data = fs.readFileSync(urlTiobeList, 'utf8');
    var array = csvToArray(data);
    array.pop(); // Eliminar la última linea innecesaria
        
    const topics = [];

    for (var index=0; index<array.length; index++) {
        let element = array[index]

        await page.goto(urlGitHub + element.topic);
        await page.waitForSelector('#js-pjax-container');

        const topicMatching = await page.evaluate(() => {
            const tmp = document.querySelector('.col-md-8 h2').innerText; 

            var matching = "";
            for (var index=9; tmp[index] != ' '; index++) {
                if (tmp[index] != ',')
                    matching = matching + tmp[index];   
            }

            return matching;
        });

        let topic = {};
        topic[headers[0]] = element.name;
        topic[headers[1]] = topicMatching;

        topics.push(topic);
    }
    await browser.close();

    var str = arrayToCsv(topics);

    if (fs.existsSync(urlResult)) {
        console.log("El directorio ya ha sido creado");

    } else {
        fs.mkdir(urlResult, (error) => {
            if (error) {
                console.log (error.message);
            }
            console.log("Directorio creado con exito");
        });
    }

    fs.writeFile(
        urlResult + "ResultadoJavaScript.txt",
        str,
        {
          encoding: "utf8",
          mode: 0o666,
          flag: "w", // falla si el archivo existe
        },
        err => {
          if (err) {
            throw err;
          }
          console.log("Archivo escrito con éxito!!!");
        }
      ); 

})();

function csvToArray(str, delimiter = ",") {
    
    const deleteOneLine = str.replace("\r", "").slice(0, str.indexOf("\n")).split(delimiter);

    // slice from \n index + 1 to the end of the text
    // use split to create an array of each csv value row
    str.replace("\r", "");

    const rows = str.slice(str.indexOf("\n") + 1).split("\n");
  
    // create an array of objects from the array of csv values
    let array = [];
    rows.forEach(row => {
        let values = row.split(delimiter);
        let element={};
        element[headers[0]] = values[0];
        element[headers[1]] = values[1];
        array.push(element);
    });

   return array;
}

function arrayToCsv (array, delimiter = ',') {
    var str = "";

    array.forEach(function(element) {
        str = str + element.name + "," + element.topic + "\n";
    });

    return str;
}
