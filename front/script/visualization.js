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

export function createDistanceGraph(distance) {

    let newNodes = new vis.DataSet();
    let newEdges = new vis.DataSet();

    for (let i = 0; i < distance.length - 1; i++) {
        let sourceNode = distance[i].Node;
        let targetNode = distance[i + 1].Node;

        if (!newNodes.get(-sourceNode)) {
            newNodes.add({id: -sourceNode, label: sourceNode.toString()});
        }
        if (!newNodes.get(-targetNode)) {
            newNodes.add({id: -targetNode, label: targetNode.toString()});
        }


        newEdges.add({
            from: -sourceNode,
            to: -targetNode,
            label: distance[i + 1].Weight.toString(),
            arrows: 'to'
        });
    }



    let newOptions = {
        nodes: {
            shadow: { enabled: true },
            font: { size: 25 },
            shape: 'circle',

        },
        edges: { width: 3, shadow: { enabled: true }, smooth: false},

    };

    let container = document.getElementById('additional-network');
    let graph = {nodes: newNodes, edges: newEdges};
    let network = new vis.Network(container, graph, newOptions);


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