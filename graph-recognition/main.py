import numpy as np
from flask import Flask, request, jsonify
import cv2 as cv
from preprocessing import preprocess
from segmentation import find_vertices
from filler import fill_vertices
from topology_recognition import recognize_topology
from postprocessing import postprocess

app = Flask(__name__)


@app.route('/process_image', methods=['POST'])
def process_image():
    #
    image_data = request.files['image']
    image = cv.imdecode(np.fromstring(image_data.read(), np.uint8), cv.IMREAD_COLOR)


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

    # Після обробки
    json_string = postprocess(vertices_list)

    return jsonify({'result': json_string}), 200


if __name__ == "__main__":
    app.run(debug=True)
