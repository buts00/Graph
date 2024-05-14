import {getDijkstra, getMst} from "./serverCommunication.js";
import {restoreGraph} from "./visualization.js";
import {mstPath} from "./main.js";


export function isValidInput(value) {
    return /^[1-9]\d*$/.test(value);
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


// export function processDataFromDijkstra() {
//     let vertexValue = document.getElementById("vertexInput").value;
//     let node = nodes.get(parseInt(vertexValue));
//     if (!isValidInput(vertexValue) || node == null) {
//         // alert("Enter valid vertex");
//         return
//     }
//     document.getElementById("vertexInput").value = "";
//     sendStartPointToServer(parseInt(vertexValue),dijkstraPath)
// }

export async function getSelectedAlgorithm() {
    const selectElementValue = document.getElementById("algorithmSelect").value;
    console.log(selectElementValue)
    if (selectElementValue === "mst") {
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
