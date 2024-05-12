import cv2 as cv
import numpy as np

from math import ceil, floor
from typing import Tuple

MAX_R_FACTOR: float = 0.04
MIN_R_FACTOR: float = 0.005
DIST_FACTOR: float = 0.06
INNER_CANNY: int = 200

VERTEX_AREA_UPPER: float = 0.2
VERTEX_AREA_LOWER: float = 0.00015

ROUND_RATIO: float = 3.0

K3 = np.ones((3, 3), dtype=np.uint8)
K5 = np.ones((5, 5), dtype=np.uint8)

def fill_vertices(image: np.ndarray) -> Tuple[np.ndarray,np.ndarray]:

    image = fill_elliptical_contours(image)
    image = fill_circular_shapes(image)
  
    image = cv.morphologyEx(image, cv.MORPH_CLOSE, K3, iterations=1)
    image = fill_small_contours(image)

    edgeless = remove_edges(image.copy())
    
    return image, edgeless


def fill_elliptical_contours(image: np.ndarray) -> np.ndarray:

    processed = cv.morphologyEx(image, cv.MORPH_CLOSE, K3, iterations=4)
    contours, hierarchy = cv.findContours(processed, cv.RETR_CCOMP, cv.CHAIN_APPROX_SIMPLE)
    img_area = image.shape[0] * image.shape[1]

    def is_valid_contour(contour, parent_idx):
        contour_area = cv.contourArea(contour)
        return (hierarchy[0][parent_idx][3] != -1 and
                img_area * VERTEX_AREA_UPPER >= contour_area >= img_area * VERTEX_AREA_LOWER)

    for i, contour in enumerate(contours):
        if is_valid_contour(contour, i):
            (x, y), (a, b), angle = cv.minAreaRect(contour)
            if 2 >= a / b >= 1.0 / 2:
                ellipse_cnt = cv.ellipse2Poly((int(x), int(y)), (int(a / 2.0), int(b / 2.0)), int(angle), 0, 360, 1)
                if contours_overlap_level(ellipse_cnt, contour) >= 0.3:
                    cv.drawContours(image, contours, i, 255, thickness=cv.FILLED)
            else:
                cv.drawContours(image, contours, i, 0, thickness=6)

    return image

def fill_circular_shapes(image: np.ndarray) -> np.ndarray:
    
    r_min, r_max, min_dist = get_hough_param(image)

    circles = cv.HoughCircles(
        image, cv.HOUGH_GRADIENT, 1,
        minDist=min_dist,
        param1=INNER_CANNY,
        param2=13,
        minRadius=r_min,
        maxRadius=r_max
    )
    if circles is not None:
        for circle in circles[0]:
            x, y, r = circle
            cv.circle(image, (int(x), int(y)), int(r), 255, thickness=cv.FILLED)

    return image


def fill_small_contours(image: np.ndarray) -> np.ndarray:
    
    contours, hierarchy = cv.findContours(image, cv.RETR_CCOMP, cv.CHAIN_APPROX_SIMPLE)
    max_area = image.shape[0] * image.shape[1] * 0.001
    for i in range(0, len(contours)):
        if hierarchy[0][i][3] != -1 and cv.contourArea(contours[i]) <= max_area:
            cv.drawContours(image, contours, i, 255, thickness=cv.FILLED)
    return image


def contours_overlap_level(contour1, contour2):
    
    dist_limit = 1 if cv.contourArea(contour1) >= 150 else 0
    overlapping_pixels = 0
    for pt in contour1:
        x, y = pt.ravel()  
        dist = abs(cv.pointPolygonTest(contour2, (int(x), int(y)), True))  
        if dist <= dist_limit: 
            overlapping_pixels += 1
    overlay_level = overlapping_pixels / float(len(contour1))

    return overlay_level


def get_hough_param(image: np.ndarray) -> Tuple[int, int, int,np.ndarray]:
    
    edgeless = remove_edges(image.copy())
    contours, hierarchy = cv.findContours(edgeless, cv.RETR_CCOMP, cv.CHAIN_APPROX_SIMPLE)
    radius_list = [cv.minEnclosingCircle(contour)[1] for contour, hier in zip(contours, hierarchy[0])
                   if hier[3] == -1 and 1.0 / 2 <= cv.boundingRect(contour)[3] / cv.boundingRect(contour)[2] <= 2]
    if radius_list:
        r_avg = np.average(radius_list)
        return floor(r_avg * 0.5), ceil(r_avg * 1.2), r_avg * 3
    else:
        width = image.shape[1]
    
    return floor(width * MIN_R_FACTOR), ceil(width * MAX_R_FACTOR), floor(width * DIST_FACTOR),edgeless


def remove_edges(image: np.ndarray) -> np.ndarray:

    eroded_contours = image.copy()
    contours_list = []

    while True:
        contours, _ = cv.findContours(eroded_contours, cv.RETR_LIST, cv.CHAIN_APPROX_SIMPLE)
        if not contours:
            break
        contours_list.append(len(contours))
        eroded_contours = cv.erode(eroded_contours, K3, iterations=1)

    max_length = 0
    current_length = 0
    best_position = 0
    last_count = contours_list[0] if contours_list else 0

    for i, count in enumerate(contours_list):
        if abs(last_count - count) <= 1:
            current_length += 1
        else:
            if current_length > max_length:
                max_length = current_length
                best_position = i - current_length
            current_length = 0
        last_count = count

    if current_length > max_length:
        best_position = len(contours_list) - current_length

    position_max = max(best_position, 0)

    eroded = cv.erode(image, K3, iterations=position_max)
    dilated = cv.dilate(eroded, K3, iterations=position_max)

    edges_removed = cv.morphologyEx(dilated, cv.MORPH_CLOSE, K5, iterations=1)
    return edges_removed