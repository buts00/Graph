class Vertex:
    id = -1
    x = -1
    y = -1
    r = -1.0  
    adjacency_list = []

    def __init__(self, x, y, r):
        self.x = x
        self.y = y
        self.r = r
        self.adjacency_list = []