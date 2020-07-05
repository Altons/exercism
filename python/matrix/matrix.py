class Matrix:
    def __init__(self, matrix):
        self.matrix = [list(map(int, x.split(" ")))
                       for x in matrix.split("\n")]

    def row(self, index):
        return self.matrix[index-1]

    def column(self, index):
        return list(map(list, zip(*self.matrix)))[index-1]
