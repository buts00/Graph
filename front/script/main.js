import {getGraph} from "./serverCommunication.js";
import {setButtons} from "./buttonHandlers.js";

export let nodes = new vis.DataSet();
export let edges = new vis.DataSet();
export const generalPath = "http://localhost:8080";
export const graphPath = generalPath + "/graph";
export const mstPath = graphPath + "/MST";
export const dijkstraPath = graphPath + "/dijkstra";

function main() {
    setButtons()
    getGraph()
}

main()











