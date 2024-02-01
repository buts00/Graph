function getData() {
    fetch('http://localhost:8080/graph')
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            processData(data)
        })
        .catch(error => {
            console.error('There was a problem with the fetch operation:', error);
        });
}

getData()

function processData(data) {
    let nodes = new vis.DataSet();
    let edges = new vis.DataSet();


    data.Edges.forEach(edge => {
        let color = '#6fd7ed'
        if (data.InMst.includes(edge.Id)) color = 'orange'
        if (!nodes.get(edge.Source)) {
            nodes.add({id: edge.Source, label: edge.Source.toString()});
        }
        if (!nodes.get(edge.Destination)) {
            nodes.add({id: edge.Destination, label: edge.Destination.toString()});
        }

        edges.add({
            from: edge.Source,
            to: edge.Destination,
            label: edge.Weight.toString(),
            color: color

        });
    });

    let options = {
        nodes: {
            shadow: {
                enabled: true,
            },
            font: {
                size: 25,
                align: 'center',
            },
            shape: 'circle',


        },

        edges: {
            width: 3,
            shadow: {
                enabled: true,
            },
            smooth: false,
        },
    };

    let container = document.getElementById('network');
    let graph = {nodes: nodes, edges: edges};
    let network = new vis.Network(container, graph, options);


}











