// Se obtiene el archivo JSON que posee la informacion para el grafica
const requestURL = 'data.json';
const request = new XMLHttpRequest();

request.open('GET', requestURL);
request.responseType = 'json';
request.send();

request.onload = function () {
    const dataJSON = request.response

    // Obtener una referencia al elemento canvas del DOM
    const $grafica = document.querySelector("#grafica");

    // Se realiza la grafica
    new Chart($grafica, {
        type: 'bar',// Tipo de gráfica
        data: {
            labels: dataJSON['NOMBRE_LENGUAJE'],
            datasets: [{
                backgroundColor: "#42C2FF",
                data: dataJSON['NRO_APARICIONES']
            }]
        },
        options: {
            scales: {
                yAxes: [{
                    ticks: {
                        beginAtZero: true
                    }
                }],
            },
            legend: { display: false },
            title: {
                display: true,
                text: "Gráficas de barras",
                fontSize: 24
            }
        }
    });

}
