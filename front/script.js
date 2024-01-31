

function getData() {
    fetch('http://localhost:8080/array')
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {

            console.log(data);
        })
        .catch(error => {

            console.error('There was a problem with the fetch operation:', error);
        });
}

getData()



let nodes = new vis.DataSet([
    { id: 1, label: '1' },
    { id: 2, label: '2' },
    { id: 3, label: '3' },
]);

let edges = new vis.DataSet([
    { from: 1, to: 2, label: '5' },
    { from: 1, to: 3, label: '3' },
    { from: 2, to: 3,  label: '7' },
]);

let options = {
    nodes: {
        font: {
            size: 25,
            align: 'center',
        },
        shape: 'circle',
        color: '#6fd7ed',


    },
    edges: {
        width: 3,
    },
};

let container = document.getElementById('network');
let data = { nodes: nodes, edges: edges };
let network = new vis.Network(container, data, options);


