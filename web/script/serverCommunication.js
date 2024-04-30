import { createDistanceGraph, createGraph, createMst } from "./visualization.js";
import { dijkstraPath, graphPath, mstPath } from "./main.js";

export function getGraph() {
    fetch(graphPath, {
        method: "GET"
    })
        .then(handleResponse)
        .then(data => {
            createGraph(data);
        })
        .catch(handleError);
}

export function getMst() {
    fetch(mstPath)
        .then(handleResponse)
        .then(data => {
            createMst(data);
        })
        .catch(handleError);
}

export function sendEdgeDataToServer(edgeData) {
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
        .catch(err => {
            handleError(err)
        });
}

export function removeEdgeFromServer(edgeData) {
    console.log("here5")
    fetch(graphPath, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(edgeData),
    })
        .then(handleResponse)
        .then(data => {
            getGraph();
        })
        .catch(err => {
            handleError(err)
        });
}

export function sendStartPointToServer(startPoint) {
    fetch(dijkstraPath, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(startPoint),
    })
        .then(handleResponse)
        .then(data => {
            createDistanceGraph(data)
        })
        .catch(handleError);
}


export function handleResponse(response) {
    if (!response.ok) {
        throw new Error('Network response was not ok');
    }
    return response.json();
}

export function handleError(error) {
    // alert("There was with the fetch operation")
    console.error('There was a problem with the fetch operation:', error);
}