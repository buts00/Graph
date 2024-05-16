import { clearInputFields, getDataFromInputs, isValidInput } from "./graphFunctions.js";
import { sendEdgeDataToServer, removeEdgeFromServer } from "./serverCommunication.js";
import { edges } from "./main.js"

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
    const textarea = document.querySelector('textarea')
    const textareaContent = textarea.value.trim().split('\n')
    const edges = []
    try {
        console.log(textareaContent)
        textareaContent.forEach(str => {
            const elements = str.split(' ').map(el => Number.parseInt(el))
            console.log(elements)
            if (elements.length === 3) {
                const [Source, Destination, Weight] = elements;
                if (isValidInput(Source) && isValidInput(Destination) && isValidInput(Weight)) {
                    edges.push({ Source, Destination, Weight })
                }
            } else if (elements.length === 2) {
                const [Source, Destination] = elements;
                if (isValidInput(Source) && isValidInput(Destination)) {
                    edges.push({ Source, Destination, Weight: 1 })
                }
            }
        })
        if (edges.length === 0) {
            textarea.value = ''
            textarea.style.borderColor = 'red'
            throw new Error('It must be at least one correct edge')
        }
        textarea.style.borderColor = 'black'
        sendEdgeDataToServer(edges)
    } catch (err) {
        console.log(err)
        textarea.value = ''
        textarea.style.borderColor = 'red'
    }
}

export function deleteEdge() {
    const textarea = document.querySelector('textarea')
    const textareaContent = textarea.value.trim().split('\n')
    const edges = []
    try {
        console.log(textareaContent)
        textareaContent.forEach(str => {
            const elements = str.split(' ').map(el => Number.parseInt(el))
            if (elements.length === 3) {
                const [Source, Destination, Weight] = elements;
                if (isValidInput(Source) && isValidInput(Destination) && isValidInput(Weight)) {
                    edges.push({Source, Destination, Weight})
                }
            } else if (elements.length === 2) {
                const [Source, Destination] = elements;
                if (isValidInput(Source) && isValidInput(Destination)) {
                    edges.push({ Source, Destination, Weight: 1 })
                }
            }
        })
        if (edges.length === 0) {
            textarea.value = ''
            textarea.style.borderColor = 'red'
            throw new Error('It must be at least one correct edge')
        }
        removeEdgeFromServer(edges)
    } catch (err) {
        console.log('error', err)
        textarea.value = ''
        textarea.style.borderColor = 'red'
    }
}
