import json

import numpy as np
import requests

from flask import Flask, request, jsonify
import cv2 as cv
from preprocessing import preprocess
from segmentation import find_vertices
from filler import fill_vertices
from topology_recognition import recognize_topology

app = Flask(__name__)


@app.route('/', methods=['POST'])
def process_image():
    if 'file' not in request.files:
        return jsonify({'error': 'No file part in the request'}), 400

    image_data = request.files['file']
    image_buffer = image_data.read()
    image = cv.imdecode(np.frombuffer(image_buffer, np.uint8), cv.IMREAD_COLOR)

    if image is None:
        return jsonify({'error': 'Error opening image'}), 400

    # Попередня обробка
    source, preprocessed = preprocess(image)

    # Заповнення
    filled_image, edgeless = fill_vertices(preprocessed)

    # Знаходження вершин
    vertices_list = find_vertices(filled_image, edgeless)
    if not vertices_list:
        return jsonify({'error': 'No vertices found'}), 400

    # Визначення топології
    vertices_list = recognize_topology(vertices_list, filled_image, source)

    json_data = []
    added_edges = set()

    for vertex_id, vertex in enumerate(vertices_list):
        for adj_vertex_id in vertex.adjacency_list:
            if (vertex_id, adj_vertex_id) not in added_edges and (adj_vertex_id, vertex_id) not in added_edges:
                json_data.append({
                    "Source": vertex_id,
                    "Destination": adj_vertex_id,
                    "Weight": 1
                })
                added_edges.add((vertex_id, adj_vertex_id))

    return jsonify(json_data), 200


if __name__ == "__main__":
    app.run(debug=True, port=14880)
