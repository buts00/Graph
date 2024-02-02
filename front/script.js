function getData() {
    fetch('http://localhost:8080/graph')
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            processData(data)
        })
        .catch(error => {
            console.error('There was a problem with the fetch operation:', error);
        });
}

getData()

let nodes = new vis.DataSet();
let edges = new vis.DataSet();
let network;

function processData(data) {
    data.Edges.forEach(edge => {
        let color = '#6895D2'
        if (data.InMst.includes(edge.Id)) color = '#FF9843'
        if (!nodes.get(edge.Source)) {
            nodes.add({id: edge.Source, label: edge.Source.toString()});
        }
        if (!nodes.get(edge.Destination)) {
            nodes.add({id: edge.Destination, label: edge.Destination.toString()});
        }

        edges.add({
            from: edge.Source,
            to: edge.Destination,
            label: edge.Weight.toString(),
            color: color

        });
    });

    let options = {
        nodes: {
            shadow: {
                enabled: true,
            },
            font: {
                size: 25,
            },
            shape: 'circle',
            color: '#86bbf8',


        },

        edges: {
            width: 3,
            shadow: {
                enabled: true,
            },
            smooth: false,

        },
        physics: {

            barnesHut: {
                centralGravity: 0.0,
                gravitationalConstant: -1000
            },

        }
    };

    let container = document.getElementById('network');
    let graph = {nodes: nodes, edges: edges};
    network = new vis.Network(container, graph, options);

}

function isValidInput(value) {
    return /^[1-9]\d*$/.test(value);
}


function addNode() {
    var fromValue = document.getElementById("from").value;
    var toValue = document.getElementById("to").value;
    var weightValue = document.getElementById("weight").value;
    document.getElementById("from").value = "";
    document.getElementById("to").value = "";
    document.getElementById("weight").value = "";

    if (!isValidInput(fromValue) || !isValidInput(toValue) || !isValidInput(weightValue)) {
        alert("Please enter valid numeric values.");
        return;
    }

    var edgeData = {
        Source: parseInt(fromValue),
        Destination: parseInt(toValue),
        Weight: parseInt(weightValue)
    };


    fetch('http://localhost:8080/graph', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(edgeData),
    })
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            console.log('Data sent successfully:', data);
            getData();
        })
        .catch(error => {
            console.error('There was a problem with the fetch operation:', error);
        });



}







