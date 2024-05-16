import {edges, nodes} from "./main.js";

export function createGraph(data) {
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

export function createDistanceGraph(path, distance) {
    let delay = 500; // delay in milliseconds
    restoreGraph()
    document.querySelector('.distance-span').innerHTML = distance === '-1' ? 'There is no such path' : distance
    function addEdgeWithDelay(index) {
        if (index < path.length) {
            let dist = path[index];
            console.log(dist)
            edges.forEach(edge => {
                if (edge.from === dist.Source && edge.to === dist.Destination) {
                    edges.remove(edge.id)
                }
            })
            if (!nodes.get(dist.Source)) {
                nodes.add({id: dist.Source, label: dist.Source.toString()});
            }
            if (!nodes.get(dist.Destination)) {
                nodes.add({id: dist.Destination, label: dist.Destination.toString()});
            }

            setTimeout(() => {
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

export function createMst(data) {
    let delay = 400;
    restoreGraph()
    function updateColor(index) {
        if (index < data.length) {
            let edgeIndex = data[index];
            let edge = edges.get(edgeIndex);
            console.log(edge)
            edge.color = '#FF9843';
            edges.update(edge);
            setTimeout(() => {
                updateColor(index + 1);
            }, delay);
        }
    }

    updateColor(0);
}

export function restoreGraph() {
    edges.forEach(edge => {
        let curEdge = edges.get(edge.id);
        curEdge.color = '#6895D2';
        curEdge.arrows = 'none';
        edges.update(curEdge);
    })
}
