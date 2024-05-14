import {createDistanceGraph, createGraph, createMst} from "./visualization.js";
import {dijkstraPath, graphPath, mstPath} from "./main.js";

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

export async function getMst() {
    try {
        const response = await fetch(mstPath)
        if (response.ok) {
            const data = await response.json()
            console.log(data)
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
        if (response.ok) {
            const data = await response.json()
            createDistanceGraph(data.path, data.distance)
        }

        // const distance = [
        //     {
        //         Source: 534,
        //         Destination: 634,
        //         Weight: 3,
        //     },
        //     {
        //         Source: 634,
        //         Destination: 3,
        //         Weight: 3,
        //     },
        //     {
        //         Source: 3,
        //         Destination: 4,
        //         Weight: 2,
        //     }
        // ]
        // createDistanceGraph(distance)
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
    console.log("in fn array edges: ", edges)
    console.log("remove from server string edges", JSON.stringify(edges))
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
    console.log(response)
    // if (!response.ok) {
    //     throw new Error('Network response was not ok');
    // }
    console.log('response is ok')
    return response.json();
}

export function handleError(error) {
    // alert("There was with the fetch operation")
    console.error('There was a problem with the fetch operation:', error);
}
