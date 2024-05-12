from math import sqrt
import cv2 as cv
import numpy as np
import functools
from Vertex import Vertex
from typing import List, Tuple

MIN_EDGE_LEN: int = 10
VERTEX_AREA_FACTOR: float = 1.3
SEARCH_RADIUS_CONST = 15


def recognize_topology(vertices_list: list, filled_image: np.ndarray, visualised: np.ndarray) -> list:
    filled_image = remove_vertices(vertices_list, filled_image, VERTEX_AREA_FACTOR)
    lines_list, backend = lines_from_contours(filled_image, visualised.copy())

    search_radius = int(np.average(np.array([v.r for v in vertices_list])))
    search_radius = search_radius if search_radius > SEARCH_RADIUS_CONST else SEARCH_RADIUS_CONST
    linked_lines, backend = link_nearby_endpoints(lines_list, backend, 1.5 * search_radius, 20)
    vertices_list, backend, visualised = edges_from_lines(linked_lines, vertices_list, backend, visualised, 3.1)

    debug = False
    if debug:
        cv.imshow("removed vertices and lines intersections", filled_image)
        cv.imshow("\"backend\" - colors description in topology.py", backend)
        cv.imshow("final results - green: vertices, red: edges", visualised)
        cv.waitKey(0)

    return vertices_list


def remove_vertices(vertices_list: list, filled_image: np.ndarray, vertex_area_factor: float) -> np.ndarray:
    for vertex in vertices_list:
        cv.circle(filled_image, (vertex.x, vertex.y), round(vertex.r * vertex_area_factor), 0, cv.FILLED)
    return filled_image


def lines_from_contours(filled_image: np.ndarray, backend: np.ndarray, min_line_length: float = 10) \
        -> Tuple[list, np.ndarray]:
    lines_list = []
    contours, hierarchy = cv.findContours(filled_image, cv.RETR_CCOMP, cv.CHAIN_APPROX_SIMPLE)
    cv.drawContours(backend, contours, -1, (0, 255, 255), 1)
    for i in range(0, len(contours)):
        if hierarchy[0][i][3] == -1:
            cnt = contours[i]
            pt1, pt2 = fit_line(cnt)
            if pt1 is not None and pt2 is not None and distance_L2(pt1, pt2) >= min_line_length:
                cv.circle(backend, (pt1[0], pt1[1]), 4, (255, 255, 0), cv.FILLED)
                cv.circle(backend, (pt2[0], pt2[1]), 4, (255, 255, 0), cv.FILLED)
                lines_list.append([pt1, pt2])
    return lines_list, backend


def fit_line(edge_contour: list, epsilon: float = 0.01, delta: float = 0.01) -> Tuple[Tuple[int, int], Tuple[int, int]]:
    perimeter = cv.arcLength(edge_contour, True)
    while True:
        approx = cv.approxPolyDP(edge_contour, perimeter * epsilon, True)
        if len(approx) <= 2:
            break
        epsilon += delta
    if len(approx) == 2:
        return approx[0][0].astype(int), approx[1][0].astype(int)
    return None, None


def link_nearby_endpoints(lines_list: list, backend: np.ndarray, search_radius: float, angle_threshold: float) \
        -> Tuple[list, np.ndarray]:
    lines_list = sorted(lines_list, key=functools.cmp_to_key(lines_lengths_compare), reverse=True)
    for i, line in enumerate(lines_list):
        if line is None:
            continue
        line_len = distance_L2(line[0], line[1])
        search_radius = min(search_radius, line_len * 1.2)
        for j in range(2):
            main_point = line[j]
            other_point = line[(j + 1) % 2]
            cv.circle(backend, tuple(main_point), int(search_radius), (255, 0, 255), 1)
            cv.line(backend, (main_point[0] - int(search_radius), main_point[1]),
                    (main_point[0] + int(search_radius), main_point[1]), (255, 0, 255), 1)
            if main_point is None or other_point is None:
                break
            else:
                main_angle = vector_angle(other_point, main_point)
                in_area_list = find_endpoints_in_area(lines_list, i + 1, main_point[0], main_point[1], search_radius,
                                                      main_angle)
                if in_area_list is not None:
                    min_index = np.argmin(in_area_list[:, 2])
                    min_delta = in_area_list[min_index, 2]
                    if min_delta <= angle_threshold:
                        k, l = int(in_area_list[min_index, 0]), int(in_area_list[min_index, 1])
                        lines_list[i][j] = lines_list[k][(l + 1) % 2]
                        lines_list[k] = None
    final_lines_list = [line for line in lines_list if line is not None]
    for line in final_lines_list:
        pt1, pt2 = line
        cv.line(backend, tuple(pt1), tuple(pt2), (0, 140, 255), 2)

    return final_lines_list, backend


def lines_lengths_compare(line1: Tuple[Tuple[int, int], Tuple[int, int]],
                          line2: Tuple[Tuple[int, int], Tuple[int, int]]) -> int:
    len1 = distance_L2(line1[0], line1[1])
    len2 = distance_L2(line2[0], line2[1])
    if len1 > len2:
        return 1
    elif len1 == len2:
        return 0
    else:
        return -1


def vector_angle(start_pt: tuple[int, int], end_pt: Tuple[int, int]) -> float:
    tmp_vec = np.array(start_pt) - np.array(end_pt)
    angle = np.arctan2(tmp_vec[0], tmp_vec[1]) * 180 / np.pi + 180
    return angle


def find_endpoints_in_area(lines_list: list, start_index: int, x: int, y: int, radius: float, main_angle: float) \
        -> np.ndarray:
    in_area_list = []
    endpoint = np.array([x, y])
    for i in range(start_index, len(lines_list)):
        if lines_list[i] is None:
            continue
        for j in range(0, 2):
            tmp_endpoint = np.array([lines_list[i][j][0], lines_list[i][j][1]])
            if distance_L2(endpoint, tmp_endpoint) <= radius:  # calculate distance
                other_endpoint = np.array([lines_list[i][(j + 1) % 2][0], lines_list[i][(j + 1) % 2][1]])
                tmp_angle = vector_angle(tmp_endpoint, other_endpoint)
                diff = abs(main_angle - tmp_angle)
                delta = diff if diff <= 180 else 360 - diff
                in_area_list.append([i, j, delta])
    ret_arr = np.array(in_area_list) if in_area_list else None
    return ret_arr


def edges_from_lines(lines_list: List, vertices_list: List, backend: np.ndarray, final_results: np.ndarray,
                     within_r_factor: float) -> Tuple[List, np.ndarray, np.ndarray]:
    for pt1, pt2 in lines_list:
        index1 = find_nearest_vertex(pt1, vertices_list)
        index2 = find_nearest_vertex(pt2, vertices_list)
        v1, v2 = (vertices_list[index1], vertices_list[index2])
        cv.circle(backend, (v1.x, v1.y), round(v1.r * within_r_factor), (127, 0, 127), thickness=2)
        cv.circle(backend, (v2.x, v2.y), round(v2.r * within_r_factor), (127, 0, 127), thickness=2)
        if point_within_radius(pt1, v1, within_r_factor) and point_within_radius(pt2, v2, within_r_factor) \
                and index1 != index2 \
                and index2 not in v1.adjacency_list and index1 not in v2.adjacency_list:
            v1.adjacency_list.append(index2)
            v2.adjacency_list.append(index1)
            cv.line(final_results, (v1.x, v1.y), (v2.x, v2.y), (0, 0, 255), thickness=2)
            cv.circle(final_results, (v1.x, v1.y), 4, (0, 0, 0), cv.FILLED)
            cv.circle(final_results, (v2.x, v2.y), 4, (0, 0, 0), cv.FILLED)

    return vertices_list, backend, final_results


def find_nearest_vertex(point: np.ndarray, vertices_list: list) -> int:
    min_distance = np.inf
    nearest_index = 0
    for i, vertex in enumerate(vertices_list):
        current_center = np.array([vertex.x, vertex.y])
        distance = np.linalg.norm(point - current_center)
        if distance < min_distance:
            min_distance = distance
            nearest_index = i
    return nearest_index


def point_within_radius(point: np.ndarray, vertex: Vertex, radius_factor: float) -> bool:
    radius = vertex.r * radius_factor
    return True if distance_L2(point, [vertex.x, vertex.y]) <= radius else False


def distance_L2(point1: Tuple[float, float], point2: Tuple[float, float]) -> float:
    return sqrt((point1[0] - point2[0]) ** 2 + (point1[1] - point2[1]) ** 2)
