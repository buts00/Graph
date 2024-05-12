import cv2 as cv
import numpy as np

from Vertex import Vertex

COLOR_R_FACTOR: float = 0.4 
COLOR_THRESHOLD: float = 0.2

VERTEX_AREA_UPPER: float = 0.1
VERTEX_AREA_LOWER: float = 0.00015

K3 = np.ones((3, 3), dtype=np.uint8)
K5 = np.ones((5, 5), dtype=np.uint8)


def find_vertices(filled: np.ndarray, edgeless: np.ndarray) -> list:
    
    round_ratio = 1.75
    contours, hierarchy = cv.findContours(edgeless, cv.RETR_CCOMP, cv.CHAIN_APPROX_SIMPLE)
    vertices_list = []
    image_area = edgeless.shape[0] * edgeless.shape[1]

    for i, cnt in enumerate(contours):
        if hierarchy[0][i][3] == -1: 
            x, y, w, h = cv.boundingRect(cnt)
            if 1.0 / round_ratio <= h / w <= round_ratio:
                (x, y), r = cv.minEnclosingCircle(cnt)
                x, y, r = (int(x), int(y), int(r * 1.05))
                fill_ratio = circle_fill_ratio(edgeless, x, y, r)
                if fill_ratio >= 0.35 and image_area * VERTEX_AREA_UPPER >= cv.contourArea(cnt) >= image_area * VERTEX_AREA_LOWER:
                    is_filled = vertex_color_fill(cv.medianBlur(filled, 5), x, y, r, COLOR_R_FACTOR, COLOR_THRESHOLD)
                    vertices_list.append(Vertex(x, y, r, is_filled, (0,0,0)))  

    return vertices_list


def circle_fill_ratio(binary: np.ndarray, x: int, y: int, r: int) -> float:
    
    Y, X = np.ogrid[:binary.shape[0], :binary.shape[1]]
    mask = (X - x) ** 2 + (Y - y) ** 2 <= r ** 2
    circle_pixels = binary[mask]
    if circle_pixels.size > 0:
        return np.mean(circle_pixels) / 255
    return 0.0


def extract_circle_area(image: np.ndarray, x: int, y: int, r: int) -> np.ndarray:
    
    y_indices, x_indices = np.ogrid[:image.shape[0], :image.shape[1]]
    mask = (x_indices - x) ** 2 + (y_indices - y) ** 2 <= r ** 2
    return image[mask]


def vertex_color_fill(binary: np.ndarray, x: int, y: int, r: float, r_factor: float,
                      threshold: float) -> bool:
    
    fill_ratio = circle_fill_ratio(binary, x, y, int(r * r_factor))
    is_filled = fill_ratio >= threshold
    return is_filled