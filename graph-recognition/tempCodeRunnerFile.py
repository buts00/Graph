import cv2 as cv
import numpy as np
from preprocessing import preprocess
from segmentation import find_vertices
from filler import fill_vertices
from topology_recognition import recognize_topology

def main():
    source = 'test_images\\graph_test_hand_5.jpg'
    debug_mode = True
    image = cv.imread(source)
    if image is None:
        print("Error opening image!")
        return -1
    
    # Preprocessing
    source, preprocessed = preprocess(image)

    #filler
    filled_image, edgeless = fill_vertices(preprocessed)
    '''if debug_mode == True:
        cv.imshow("filled image", filled_image)    
        cv.waitKey(0)'''

    # vertices finder
    vertices_list = find_vertices(filled_image, edgeless)
    if not vertices_list:
        print("No vertices found")
        return -1

    # Topology Recognition
    vertices_list = recognize_topology(vertices_list, filled_image, source)
    
    for vertex in vertices_list:
        print(f"Vertex ({vertex.x}, {vertex.y}) with radius {vertex.r}")
        print("Adjacency List:", vertex.adjacency_list)

    print("Process completed successfully.")
    return 0

if __name__ == "__main__":
    main()
