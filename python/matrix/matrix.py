class Matrix:
    def __init__(self, matrix_string):
        self.str = matrix_string
        self.rows = [list(map(int, x.split(" ")))
                     for x in self.str.split("\n")]

    def row(self, index):
        return self.rows[0] if len(self.rows) == 1 else self.rows[index-1]

    def column(self, index):
        return self.rows[0] if len(self.rows) == 1 else list(map(list, zip(*self.rows)))[index-1]
