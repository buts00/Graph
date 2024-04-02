import {clearInputFields, getDataFromInputs, isValidInput} from "./graphFunctions.js";
import {sendEdgeDataToServer} from "./serverCommunication.js";
import {edges} from "./main.js"

export function isEdgeAlreadyExists(startNode, endNode, weight) {
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


export function addEdge() {
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
    sendEdgeDataToServer(edgeData)
}

export function deleteEdge() {
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

    sendEdgeDataToServer(edgeData)
}

