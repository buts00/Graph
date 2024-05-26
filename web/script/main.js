let nodes = new vis.DataSet();
let edges = new vis.DataSet();
const generalPath = "http://localhost:8080/";
const graphPath = generalPath + "graph/";
const mstPath = graphPath + "MST";
const dijkstraPath = graphPath + "dijkstra";

function setButtons() {
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
    const clearButton = document.querySelector('.clear-button');
    clearButton.addEventListener('click', clearGraph)
    const hideButton = document.querySelector('.hide-button');
    hideButton.addEventListener('click', () => {
        const showHideImage = document.querySelector('.show-hide-img')
        const mainContainer = document.querySelector('.main-container')
        if (showHideImage.src.includes('hide')) {
            showHideImage.src = 'img/show.svg'
            mainContainer.classList.add('hidden')
        } else {
            showHideImage.src = 'img/hide.svg'
            mainContainer.classList.remove('hidden')
        }
    })
}

function addEdge() {
    const textarea = document.querySelector('textarea')
    const textareaContent = textarea.value.trim().split('\n')
    const edges = []
    try {
        const errorContainer = document.querySelector('.error-container')
        errorContainer.innerHTML = ''
        textareaContent.forEach((str, i) => {
            const elements = str.split(' ').map(el => Number.parseInt(el))
            if (elements.length === 3) {
                const [Source, Destination, Weight] = elements;
                if (isValidInput(Source) && isValidInput(Destination) && isValidInput(Weight)) {
                    edges.push({Source, Destination, Weight})
                } else {
                    errorContainer.innerHTML += `<p>there is an error in line ${i + 1}</p>`
                }
            } else if (elements.length === 2) {
                const [Source, Destination] = elements;
                if (isValidInput(Source) && isValidInput(Destination)) {
                    edges.push({Source, Destination, Weight: 1})
                } else {
                    errorContainer.innerHTML += `<p>there is an error in line ${i + 1}</p>`
                }
            } else {
                errorContainer.innerHTML += `<p>there is an error in line ${i + 1}</p>`
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
        textarea.value = ''
        textarea.style.borderColor = 'red'
    }
}

function deleteEdge() {
    const textarea = document.querySelector('textarea')
    const textareaContent = textarea.value.trim().split('\n')
    const edges = []
    try {
        const errorContainer = document.querySelector('.error-container')
        textareaContent.forEach(str => {
            const elements = str.split(' ').map(el => Number.parseInt(el))
            if (elements.length === 3) {
                const [Source, Destination, Weight] = elements;
                if (isValidInput(Source) && isValidInput(Destination) && isValidInput(Weight)) {
                    edges.push({Source, Destination, Weight})
                } else {
                    errorContainer.innerHTML += `<p>there is an error in line ${i + 1}</p>`
                }
            } else if (elements.length === 2) {
                const [Source, Destination] = elements;
                if (isValidInput(Source) && isValidInput(Destination)) {
                    edges.push({Source, Destination, Weight: 1})
                } else {
                    errorContainer.innerHTML += `<p>there is an error in line ${i + 1}</p>`
                }
            } else {
                errorContainer.innerHTML += `<p>there is an error in line ${i + 1}</p>`
            }
        })
        if (edges.length === 0) {
            textarea.value = ''
            textarea.style.borderColor = 'red'
            throw new Error('It must be at least one correct edge')
        }
        removeEdgeFromServer(edges)
    } catch (err) {
        textarea.value = ''
        textarea.style.borderColor = 'red'
    }
}

async function getGraph() {
    await fetch(graphPath, {
        method: "GET"
    })
        .then(handleResponse)
        .then(data => {
            createGraph(data);
        })
        .catch(handleError);
}

function isValidInput(value) {
    return /^\d+$/.test(value);
}

const clearGraph = async () => {
    await fetch(graphPath, {
            method: "DELETE",
            body: JSON.stringify([])
        }
    )
    await getGraph()
}

async function getSelectedAlgorithm() {
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

async function getMst() {
    try {
        const response = await fetch(mstPath)
        if (response.ok) {
            const data = await response.json()
            createMst(data)
        }
    } catch (e) {
        handleError(e)
    }
}

const getDijkstra = async (nodeFrom, nodeTo) => {
    try {
        const response = await fetch(dijkstraPath + `?s=${nodeFrom}&d=${nodeTo}`, {
            method: "POST"
        })
        const data = await response.json()
        if (response.ok) {
            createDistanceGraph(data.path, data.distance)
        }
        const {distance} = data
        if (!distance) {
            document.querySelector('.distance-block').innerHTML = 'There is no such path'
        }
    } catch (e) {
        handleError(e)
    }
};

function sendEdgeDataToServer(edgeData) {
    fetch(graphPath, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(edgeData),
    })
        .then(handleResponse)
        .then(data => {
            getGraph();
        })
        .catch(err => {
            handleError(err)
        });
}

async function removeEdgeFromServer(edges) {
    fetch(graphPath, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(edges),
    })
        .then(handleResponse)
        .then(data => {
            getGraph();
        })
        .catch(err => {
            handleError(err)
        });
}

function handleResponse(response) {
    // if (!response.ok) {
    //     throw new Error('Network response was not ok');
    // }
    return response.json();
}

function handleError(error) {
    // alert("There was with the fetch operation")
    console.error('There was a problem with the fetch operation:', error);
}

function createGraph(data) {
    edges.clear();
    nodes.clear();
    data.forEach(edge => {
        if (!nodes.get(edge.Source)) {
            nodes.add({id: edge.Source, label: edge.Source.toString()});
        }
        if (!nodes.get(edge.Destination)) {
            nodes.add({id: edge.Destination, label: edge.Destination.toString()});
        }

        edges.add({
            id: edge.Id,
            from: edge.Source,
            to: edge.Destination,
            label: edge.Weight.toString(),
            color: '#6895D2'
        });
    });

    let options = {
        nodes: {
            shadow: {
                enabled: true,
            },
            font: {
                size: 25,
            },
            shape: 'circle',
            color: '#86bbf8',

        },

        edges: {
            width: 3,
            shadow: {
                enabled: true,
            },
            smooth: false
        },

        physics: {
            barnesHut: {
                centralGravity: 0.0,
                gravitationalConstant: -1000,
            },

        },

    };

    let container = document.getElementById('network');
    let graph = {nodes: nodes, edges: edges};
    let network = new vis.Network(container, graph, options);

}

function createDistanceGraph(path, distance) {
    let delay = 500; // delay in milliseconds
    restoreGraph()

    document.querySelector('.distance-block').innerHTML = distance === -1 ? 'There is no such path' : `Distance: ${distance}`
    path.reverse()

    function addEdgeWithDelay(index) {
        if (index < path.length) {
            let dist = path[index];
            if (!nodes.get(dist.Source)) {
                nodes.add({id: dist.Source, label: dist.Source.toString()});
            }
            if (!nodes.get(dist.Destination)) {
                nodes.add({id: dist.Destination, label: dist.Destination.toString()});
            }
            setTimeout(() => {
                edges.forEach(edge => {
                    if (edge.from === dist.Source && edge.to === dist.Destination) {
                        edges.remove(edge.id)
                    }
                })
                edges.add({
                    from: dist.Source,
                    to: dist.Destination,
                    label: dist.Weight.toString(),
                    arrows: 'to', // This will add an arrow to the edge, making it directed
                    color: 'red'
                });

                addEdgeWithDelay(index + 1);
            }, delay);

        }
    }

    addEdgeWithDelay(0);


    let options = {
        nodes: {
            shadow: {
                enabled: true,
            },
            font: {
                size: 25,
            },
            shape: 'circle',
            color: '#86bbf8',
        },
        edges: {
            width: 3,
            shadow: {
                enabled: true,
            },
            smooth: false
        },
        physics: {
            barnesHut: {
                centralGravity: 0.0,
                gravitationalConstant: -1000,
            },
        },
    };

    let container = document.getElementById('network');
    let graph = {nodes: nodes, edges: edges};
    let network = new vis.Network(container, graph, options);
}

function createMst(data) {
    let delay = 300;
    restoreGraph()
    edges.forEach(e => console.log(e))

    function updateColor(index) {
        if (index < data.length) {
            let edgeIndex = data[index];
            let edge = edges.get(edgeIndex);
            console.log(edgeIndex, edge)
            edge.color = '#FF9843';
            edges.update(edge);
            setTimeout(() => {
                updateColor(index + 1);
            }, delay);
        }
    }

    updateColor(0);
}

function restoreGraph() {
    edges.forEach(edge => {
        let curEdge = edges.get(edge.id);
        curEdge.color = '#6895D2';
        curEdge.arrows = 'none';
        edges.update(curEdge);
    })
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

function main() {
    setButtons()
    getGraph()
}

main()
