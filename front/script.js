let nodes = new vis.DataSet();
let edges = new vis.DataSet();
let network;
const generalPath = "http://localhost:8080";
const graphPath = generalPath + "/graph";
const mstPath = graphPath + "/MST";


function getGraph() {
    fetch(graphPath)
        .then(handleResponse)
        .then(data => {
            createGraph(data);
        })
        .catch(handleError);
}

getGraph()

function getMst() {
    fetch(mstPath)
        .then(handleResponse)
        .then(data => {
            createMst(data);
        })
        .catch(handleError);
}



function createGraph(data) {
    edges.clear();
    nodes.clear();
    data.forEach(edge => {
        if (!nodes.get(edge.Source)) {
            nodes.add({id: edge.Source, label: edge.Source.toString()});
        }
        if (!nodes.get(edge.Destination)) {
            nodes.add({id: edge.Destination, label: edge.Destination.toString()});
        }

        edges.add({
            id: edge.Id,
            from: edge.Source,
            to: edge.Destination,
            label: edge.Weight.toString(),
            color: '#6895D2'

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

function isEdgeAlreadyExists(startNode, endNode, weight) {
    const existingEdge = edges.get({
        filter: item => {
            return (
                item.from === startNode &&
                item.to === endNode &&
                item.label === weight.toString()
            );
        }
    });
    return existingEdge.length > 0;
}


function addEdge() {
    let edgeData = getDataFromInputs()
    const { Source, Destination, Weight } = edgeData;
    if (!isValidInput(Source) || !isValidInput(Destination) || !isValidInput(Weight) ) {
        alert("Please enter valid numeric values.");
        return;
    }
    clearInputFields()
    if (isEdgeAlreadyExists(Source,Destination,Weight) || isEdgeAlreadyExists(Destination, Source, Weight)) {
        alert("This edge is already exist");
        return;
    }
    sendDataToServer(edgeData)
}

function deleteEdge() {
    let edgeData = getDataFromInputs()
    const { Source, Destination, Weight } = edgeData;
    if (!isValidInput(Source) || !isValidInput(Destination) || !isValidInput(Weight) ) {
        alert("Please enter valid numeric values.");
        return;
    }
    clearInputFields()
    if (!isEdgeAlreadyExists(Source,Destination,Weight) && !isEdgeAlreadyExists(Destination, Source, Weight)) {
        alert("There is not such edge");
        return;
    }

    sendDataToServer(edgeData)
}

function  getDataFromInputs() {
    let fromValue = document.getElementById("from").value;
    let toValue = document.getElementById("to").value;
    let weightValue = document.getElementById("weight").value;
    edgeData = {
        Source: parseInt(fromValue),
        Destination: parseInt(toValue),
        Weight: parseInt(weightValue)
    };
    return edgeData
}

function sendDataToServer(edgeData) {
    fetch(graphPath, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(edgeData),
    })
        .then(handleResponse)
        .then(data => {
                getGraph();
        })
        .catch(handleError);
}

function handleResponse(response) {
    if (!response.ok) {
        throw new Error('Network response was not ok');
    }
    return response.json();
}

function handleError(error) {
    console.error('There was a problem with the fetch operation:', error);
}

function clearInputFields() {
    document.getElementById("from").value = "";
    document.getElementById("to").value = "";
    document.getElementById("weight").value = "";
}

function createMst(data) {
    data.forEach(edgeIndex => {
        let edge = edges.get(edgeIndex);
        edge.color = '#FF9843';
        edges.update(edge);
    });
}

function clearMst() {
    edges.forEach(edge => {
        let curEdge = edges.get(edge.id);
        curEdge.color = '#6895D2';
        edges.update(curEdge);
    })
}


function getSelectedAlgorithm() {
    let selectElement = document.getElementById("algorithmSelect").value;
    if (selectElement === "mst") {
        getMst()
    } else {
        clearMst()
    }
}







