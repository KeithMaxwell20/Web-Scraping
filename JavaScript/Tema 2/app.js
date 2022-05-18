const puppeteer = require('puppeteer'); //Se llama a la librería puppeteer
const fs = require('fs'); //Se llama al módulo para interactuar con los archivos


const url = "https://github.com/topics/c";

const NRO_ITERACIONES_BUTTON = 1;


main();

async function main() {
    let aux = await scraping();
    let arrayMatchesTopic = convertObjToArray(aux);

    console.table(arrayMatchesTopic);

    writeDataJSON(arrayMatchesTopic);
    
    // Se escribe en un archivos los resultados
    // Se verifica que exista el directorio y que sino se cree
    if (stateDirectory()) {
        fs.writeFileSync(urlResult + "ResultadosJavaScriptTema2.csv", arrayToStr(arrayMatchesTopic));
        console.log("El archivo con los resultados ha sido creado con exito");
    }
}

async function scraping() {
    const browser = await puppeteer.launch(); // Lanzamos un nuevo navegador
    const page = await browser.newPage(); // Abrimos una nueva página
    await page.goto(url); // Vamos a la URL

    // Mensaje
    console.log("Se esta obtiendo los datos...");
    console.log("Esto puede demorar unos minutos.. ");

    // Interactuamos con el bottom de 'load more', 30 veces 
    for (let i = 0; i < NRO_ITERACIONES_BUTTON; i++) {
        await page.click('.ajax-pagination-form button');
        await page.waitForTimeout(3000);
    }

    const objMatchesTopcis= await page.evaluate(() => {
        // Obtenemos todos los articulos de la pagina
        const articles = document.querySelectorAll('#js-pjax-container article');

        const topicsName = [];
        const matchesTopics = [];

        // Recorremos articulo por articulo
        for (let article of articles) {
            // Obtenemos el DOM de los topics del articulo
            let topcisForEachArticle = article.querySelectorAll('.topic-tag-link');


            // Obtenemos el DOM de la fecha de este articulo
            let date = article.querySelector('.mr-4 relative-time');

            // Se evalua que el articulo cumpla con los requisitos del problema
            if (date != null && topcisForEachArticle.length > 0) {
                if (Date.now() - Date.parse(date.title) < 2592000000) {
                    for (let topic of topcisForEachArticle) {
                        let topicName = topic.innerText; // De cada topic obtenemos solo el nombre 

                        // Se obtiene el index del topic
                        let index = topicsName.indexOf(topicName);

                        // En caso que existe el topic se suma al contador
                        if (index != -1) {
                            matchesTopics[index]++;
                        } else { // Se agrega si no existe
                            topicsName.push(topicName);
                            matchesTopics.push(1);
                        }
                    }
                }
            }
        }

        let obj = {
            TOPIC: topicsName,
            NRO_APARICIONES: matchesTopics,
        }

        return obj;
    });

    await browser.close(); // Cerramos el navegador

    return objMatchesTopcis;
}

function convertObjToArray (obj) {
    const array = [];
    for (let i=obj.TOPIC.length-1; i >= 0; i--) {
        array.push(
            {
                TOPIC: obj.TOPIC[i],
                NRO_APARICIONES: obj.NRO_APARICIONES[i],
            }
        );
    }

    array.sort(function (a, b) {
        return b.NRO_APARICIONES - a.NRO_APARICIONES;
    });

    return array;
}

async function writeDataJSON (arrayMatchesTopic) {
    // Se crea un objeto que posee dos array para guardar en el JSON, esto es para la gráfica
    const data = {
        'TOPIC': [],
        'NRO_APARICIONES': []
    };

    // Se obtiene los primeros 10
    for (let i = 0; i < 10; i++) {
        let e = arrayMatchesTopic[i];
        data.TOPIC.push(arrayMatchesTopic[i].TOPIC);
        data.NRO_APARICIONES.push(arrayMatchesTopic[i].NRO_APARICIONES);
    }
   
    // Se escribe el objeto en un archivo JSON en formato JSON para poder recuperar luego
    fs.writeFileSync("data.json", JSON.stringify(data));
}

// Función que convierte un array a string para guardar en un archivo
function arrayToStr(array, delimiter = ',') {
    var str = "";

    array.forEach(function (element) {
        str = str + element.TOPIC + delimiter + element.NRO_APARICIONES + "\n";
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
