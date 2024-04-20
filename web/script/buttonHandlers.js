import {addEdge, deleteEdge} from "./edgeOperations.js";
import {getSelectedAlgorithm, processDataFromDijkstra} from "./graphFunctions.js";

export function setButtons() {
    const addEdgeButton = document.querySelector('.add-edge');
    addEdgeButton.addEventListener('click', addEdge)
    const removeEdgeButton = document.querySelector('.remove-edge');
    removeEdgeButton.addEventListener('click', deleteEdge)
    const selectAlgorithmButton = document.querySelector('.algorithm-button');
    selectAlgorithmButton.addEventListener('click', getSelectedAlgorithm)
    const confirmDijkstraButton = document.querySelector('.confirm-dijkstra');
    confirmDijkstraButton.addEventListener('click', processDataFromDijkstra)
}