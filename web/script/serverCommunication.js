import {createDistanceGraph, createGraph, createMst} from "./visualization.js";
import {dijkstraPath, graphPath, mstPath} from "./main.js";

export async function getGraph() {
    await fetch(graphPath, {
        method: "GET"
    })
        .then(handleResponse)
        .then(data => {
            createGraph(data);
        })
        .catch(handleError);
}


export async function getMst() {
    try {
        const response = await fetch(mstPath)
        if (response.ok) {
            const data = await response.json()
            createMst(data)
        }
    } catch (e) {
        handleError(e)
    }
}

export const getDijkstra = async (nodeFrom, nodeTo) => {
    try {
        const response = await fetch(dijkstraPath + `?s=${nodeFrom}&d=${nodeTo}`, {
            method: "POST"
        })
        const data = await response.json()
        if (response.ok) {
            createDistanceGraph(data.path, data.distance)
        }
        const {distance} = data
        if (!distance) {
            document.querySelector('.distance-block').innerHTML = 'There is no such path'
        }
    } catch (e) {
        handleError(e)
    }
};

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

export async function removeEdgeFromServer(edges) {
    fetch(graphPath, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(edges),
    })
        .then(handleResponse)
        .then(data => {
            getGraph();
        })
        .catch(err => {
            handleError(err)
        });
}

export function handleResponse(response) {
    // if (!response.ok) {
    //     throw new Error('Network response was not ok');
    // }
    return response.json();
}

export function handleError(error) {
    // alert("There was with the fetch operation")
    console.error('There was a problem with the fetch operation:', error);
}
