import {getDijkstra, getGraph, getMst} from "./serverCommunication.js";
import {restoreGraph} from "./visualization.js";
import {graphPath, mstPath} from "./main.js";


export function isValidInput(value) {
    return /^\d+$/.test(value);
}

export const clearGraph = async () => {
    await fetch(graphPath, {
            method: "DELETE",
            body: JSON.stringify([])
        }
    )
    await getGraph()
}

export function getDataFromInputs() {
    let fromValue = document.getElementById("from").value;
    let toValue = document.getElementById("to").value;
    let weightValue = document.getElementById("weight").value;
    return {
        Source: parseInt(fromValue),
        Destination: parseInt(toValue),
        Weight: parseInt(weightValue)
    }
}

export function clearInputFields() {
    document.getElementById("from").value = "";
    document.getElementById("to").value = "";
    document.getElementById("weight").value = "";
}

export async function getSelectedAlgorithm() {
    const selectElementValue = document.getElementById("algorithmSelect").value;
    if (selectElementValue === "mst") {
        await getGraph()
        await getMst(mstPath)
    } else if (selectElementValue === "dijkstra") {
        restoreGraph()
        const nodeFrom = document.querySelector("#vertexInputFrom").value;
        const nodeTo = document.querySelector("#vertexInputTo").value;
        await getDijkstra(nodeFrom, nodeTo)
    } else {
        restoreGraph()
    }
}
