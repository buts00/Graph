import {getGraph} from "./serverCommunication.js";
import {setButtons} from "./buttonHandlers.js";

export let nodes = new vis.DataSet();
export let edges = new vis.DataSet();
export const generalPath = "http://localhost:8080/";
export const graphPath = generalPath + "graph/";
export const mstPath = graphPath + "MST";
export const dijkstraPath = graphPath + "dijkstra";

function main() {
    setButtons()
    getGraph()
}

const imgPreviewNode = document.querySelector('.img-preview')
const imgInput = document.querySelector('.img-input')
const formNode = document.querySelector('.form')

imgInput.addEventListener('change', () => {
    const file = imgInput.files[0]
    if (file) {
        imgPreviewNode.style.display = 'block'
        imgPreviewNode.src = URL.createObjectURL(file)
    }
})

formNode.addEventListener('submit', async (e) => {
    console.log('in submit form')
    e.preventDefault()
    try {
        const formData = new FormData(formNode)
        if (!imgInput.files[0]) {
            throw new Error('You must upload an image')
        }
        try {
            const res = await fetch(generalPath + "graph/image", {
                method: 'POST',
                body: formData
            })
        } catch {
            throw new Error('There is no valid graph in the image')
        }
        // const data = await res.json()
        await getGraph()
        imgPreviewNode.style.display = 'none'
        imgInput.value = ''

    } catch (e) {
        alert(e.message)
    }

})


main()











