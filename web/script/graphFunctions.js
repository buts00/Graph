// import {getMst, sendStartPointToServer} from "./serverCommunication.js";
import {clearMst} from "./visualization.js";
import {dijkstraPath, mstPath, nodes} from "./main.js";

export function isValidInput(value) {
    return /^[1-9]\d*$/.test(value);
}


export function  getDataFromInputs() {
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


export function processDataFromDijkstra() {
    let vertexValue = document.getElementById("vertexInput").value;
    let node = nodes.get(parseInt(vertexValue));
    if (!isValidInput(vertexValue) || node == null) {
        alert("Enter valid vertex");
        return
    }
    document.getElementById("vertexInput").value = "";
    sendStartPointToServer(parseInt(vertexValue),dijkstraPath)
}

export function getSelectedAlgorithm() {
    let selectElement = document.getElementById("algorithmSelect").value;
    let inputContainer = document.getElementById("input-node-container");
    if (selectElement === "mst") {
        getMst(mstPath)
        inputContainer.style.display = "none";
    } else if (selectElement === "dijkstra") {
        clearMst()
        inputContainer.style.display = "flex";
    } else {
        clearMst()
        inputContainer.style.display = "none";
    }
}