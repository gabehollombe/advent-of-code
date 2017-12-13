from itertools import count, cycle

def flatten(l):
    return (j for i in l for j in i)


class Cell():
    def __init__(self, grid, x, y, value=0):
        self.grid = grid
        self.x = x
        self.y = y
        self.value = value

    @property
    def coords(self):
        return self.x, self.y

    @property
    def neighbors(self):
        return {
            self.grid.cell(self.x + 1, self.y + 0),
            self.grid.cell(self.x + 1, self.y + 1),
            self.grid.cell(self.x + 0, self.y + 1),
            self.grid.cell(self.x - 1, self.y + 1),
            self.grid.cell(self.x - 1, self.y + 0),
            self.grid.cell(self.x - 1, self.y - 1),
            self.grid.cell(self.x + 0, self.y - 1),
            self.grid.cell(self.x + 1, self.y - 1),
        }

    def __repr__(self):
        return f"({self.x}, {self.y})"


class Grid():
    DIRECTION_OFFSETS = {
        'e': (1, 0),
        'n': (0, 1),
        'w': (-1, 0),
        's': (0, -1),
    }

    def __init__(self):
        self.cells_by_coords = dict()
        self.cells = list()
        self._steps = self._steps_per_direction()
        self._head_cell = self.cell(0, 0, 1)

    def cell(self, x, y, value=0):
        if not self.cells_by_coords.get((x, y)):
            self.cells_by_coords[(x, y)] = Cell(self, x, y, value)
        return self.cells_by_coords[(x, y)]

    @property
    def head_cell(self):
        return self._head_cell

    def populate_until_value(self, value):
        while self._head_cell.value < value:
            self.add_next_cell()

    def populate_until_length(self, length):
        while len(self.cells) < length:
            self.add_next_cell()

    def add_next_cell(self):
        self.advance_head_cell()
        neighbors = list(self._head_cell.neighbors)
        sum_neighbors = sum([c.value for c in neighbors])
        self._head_cell.value = sum_neighbors

    def nth_cell(self, n):
        return list(self.cells)[n]

    def _steps_per_direction(self):
        dimensions = flatten(zip(count(1), count(1)))
        directions = cycle(['e', 'n', 'w', 's'])

        direction_dimension_pairs = zip(dimensions, directions)
        return flatten(count * [dir] for count, dir in direction_dimension_pairs)

    def advance_head_cell(self):
        self.cells.append(self._head_cell)
        coords = self._coords_after_move(self._head_cell, next(self._steps))
        self._head_cell = self.cell(*coords)

    def _coords_after_move(self, from_cell, direction):
        x, y = from_cell.coords
        dx, dy = self.DIRECTION_OFFSETS[direction]
        return x+dx, y+dy


def test_neighbors():
    grid = Grid()
    c = grid.cell(0, 0)
    assert c.neighbors == {
        grid.cell(1, 0),
        grid.cell(1, 1),
        grid.cell(0, 1),
        grid.cell(-1, 1),
        grid.cell(-1, 0),
        grid.cell(-1, -1),
        grid.cell(0, -1),
        grid.cell(1, -1),
    }


def manhattan_distance_to_center_for_square(i):
    grid = Grid()
    grid.populate_until_length(i)
    x, y = grid.head_cell.coords
    return abs(x) + abs(y)


def test_manhattan_distance_to_center_for_square():
    assert manhattan_distance_to_center_for_square(0) == 0
    assert manhattan_distance_to_center_for_square(11) == 3
    assert manhattan_distance_to_center_for_square(22) == 2
    assert manhattan_distance_to_center_for_square(1023) == 31


def test_sum_diagonals_for_square():
    grid = Grid()
    grid.populate_until_value(55)
    assert grid.nth_cell(0).value == 1
    assert grid.nth_cell(1).value == 1
    assert grid.nth_cell(2).value == 2
    assert grid.nth_cell(3).value == 4
    assert grid.nth_cell(4).value == 5
    assert grid.nth_cell(10).value == 54


print(manhattan_distance_to_center_for_square(361527))

grid = Grid()
grid.populate_until_value(361527)
print(grid.head_cell.value)

