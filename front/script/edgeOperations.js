import {clearInputFields, edges, getDataFromInputs, isValidInput} from "./main";
import {sendEdgeDataToServer} from "./serverCommunication";

export function isEdgeAlreadyExists(startNode, endNode, weight, edges) {
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
    if (!isEdgeAlreadyExists(Source,Destination,Weight,edges) && !isEdgeAlreadyExists(Destination, Source, Weight,edges)) {
        alert("There is not such edge");
        return;
    }

    sendEdgeDataToServer(edgeData)
}

