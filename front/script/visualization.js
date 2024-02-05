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

export function createDistanceGraph(Distance) {
    console.log(Distance)
}

export function createMst(data) {
    let delay = 400;
    function updateColor(index) {
        if (index < data.length) {
            let edgeIndex = data[index];
            let edge = edges.get(edgeIndex);
            edge.color = '#FF9843';
            edges.update(edge);
            setTimeout(() => {
                updateColor(index + 1);
            }, delay);
        }
    }
    updateColor(0);
}

export function clearMst() {
    edges.forEach(edge => {
        let curEdge = edges.get(edge.id);
        curEdge.color = '#6895D2';
        edges.update(curEdge);
    })
}