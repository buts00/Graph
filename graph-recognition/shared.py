from typing import Tuple
from math import sqrt

def distance_L2(point1: Tuple[float, float], point2: Tuple[float, float]) -> float:
    return sqrt((point1[0] - point2[0])**2 + (point1[1] - point2[1])**2)