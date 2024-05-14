import {addEdge, deleteEdge} from "./edgeOperations.js";
import {getSelectedAlgorithm} from "./graphFunctions.js";

export function setButtons() {
    const addEdgeButton = document.querySelector('.add-edge');
    addEdgeButton.addEventListener('click', addEdge)
    const removeEdgeButton = document.querySelector('.remove-edge');
    removeEdgeButton.addEventListener('click', deleteEdge)

    const selectElement = document.getElementById("algorithmSelect");
    selectElement.addEventListener('change', (event) => {
        let selectElementValue = event.target.value;
        let inputContainer = document.getElementById("input-node-container");
        if (selectElementValue === "mst") {
            inputContainer.style.display = "none";
        } else if (selectElementValue === "dijkstra") {
            inputContainer.style.display = "block";
        } else {
            inputContainer.style.display = "none";
        }
    })

    const selectAlgorithmButton = document.querySelector('.algorithm-button');
    selectAlgorithmButton.addEventListener('click', getSelectedAlgorithm)
    // const confirmDijkstraButton = document.querySelector('.confirm-dijkstra');
    // confirmDijkstraButton.addEventListener('click', processDataFromDijkstra)
}
