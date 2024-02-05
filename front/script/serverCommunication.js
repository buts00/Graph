import {createDistanceGraph, createGraph, createMst} from "./visualization";
import {dijkstraPath, graphPath, mstPath} from "./main";

export function getGraph() {
    fetch(graphPath)
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
        .catch(handleError);
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
    alert("There was with the fetch operation")
    console.error('There was a problem with the fetch operation:', error);
}