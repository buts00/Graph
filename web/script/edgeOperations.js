import { clearInputFields, getDataFromInputs, isValidInput } from "./graphFunctions.js";
import { sendEdgeDataToServer } from "./serverCommunication.js";
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

export function parseTextArea() {
    const textarea = document.querySelector('textarea')
    const textareaContent = textarea.value.trim().split('\n')
    const edges = []
    textareaContent.forEach(str => {
        const elements = str.split(' ').map(el => Number.parseInt(el))
        if (elements.length === 3) {
            const [Source, Destination, Weight] = elements;
            edges.push({ Source, Destination, Weight })
        } else if (elements.length === 2) {
            const [Source, Destination] = elements;
            edges.push({ Source, Destination, Weight: 1 })
        } else {
            textarea.value = ''
            throw new Error('Invalid input')
        }
    })

    return edges
}

export function addEdge() {

    const textarea = document.querySelector('textarea')
    const textareaContent = textarea.value.trim().split('\n')
    const edges = []
    try {
        textareaContent.forEach(str => {
            const elements = str.split(' ').map(el => Number.parseInt(el))
            if (elements.length === 3) {
                const [Source, Destination, Weight] = elements;
                if (!isValidInput(Source) || !isValidInput(Destination) || !isValidInput(Weight)) {
                    throw new Erorr("Please enter valid numeric values.");
                }
                edges.push({ Source, Destination, Weight })
            } else if (elements.length === 2) {
                const [Source, Destination] = elements;
                if (!isValidInput(Source) || !isValidInput(Destination)) {
                    throw new Erorr("Please enter valid numeric values.");
                }
                edges.push({ Source, Destination, Weight: 1 })
            } else {
                throw new Error('Invalid input')
            }
        })
        console.log('finish')
        textarea.style.borderColor = 'black'
        sendEdgeDataToServer(edges)
    } catch (err) {
        textarea.value = ''
        textarea.style.borderColor = 'red'
        return
    }
}

export function deleteEdge() {
    const textarea = document.querySelector('textarea')
    const textareaContent = textarea.value.trim().split('\n')
    const edges = []
    try {
        textareaContent.forEach(str => {
            const elements = str.split(' ').map(el => Number.parseInt(el))
            if (elements.length === 3) {
                const [Source, Destination, Weight] = elements;
                if (!isValidInput(Source) || !isValidInput(Destination) || !isValidInput(Weight)) {
                    throw new Erorr("Please enter valid numeric values.");
                }
                if (!isEdgeAlreadyExists(Source, Destination, Weight) && !isEdgeAlreadyExists(Destination, Source, Weight)) {
                    throw new Erorr("Edge does not exist");
                }
                edges.push({ Source, Destination, Weight })
            } else {
                throw new Error('Invalid input. Must be 3 ints')
            }
            removeEdgeFromServer(edgeData)
        })
    } catch (err) {
        textarea.value = ''
        textarea.style.borderColor = 'red'
        return
    }

}

