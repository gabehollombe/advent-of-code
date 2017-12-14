from collections import defaultdict, namedtuple
from typing import List

import pytest

Report = namedtuple('Report', ['name', 'weight', 'names_of_children'])
Program = namedtuple('Report', ['name', 'weight', 'children'])

class Towers():
    def __init__(self, reports: List[Report]):
        self.roots = dict()
        self.parents = dict() # Dict[str, Program]
        for report in reports:
            program = Program(report.name, report.weight, {})
            if program.name in self.parents:
                self.parents[program.name].children[program.name] = program
            else:
                self.roots[report.name] = program
            for name in report.names_of_children:
                self.parents[name] = program
                if name in self.roots:
                    program.children[name] = self.roots.pop(name)


def find_root_name(reports):
    towers = Towers(reports)
    return list(towers.roots.keys())[0]


def parse_lines(lines):
    lines = [l.strip() for l in lines]
    reports = list()
    child_sigil = ' -> '
    for line in lines:
        if child_sigil in line:
            name_and_weight, children = line.split(child_sigil)
            children = set(children.split(", "))
        else:
            name_and_weight = line
            children = set()
        name, weight = name_and_weight.split(" ")
        weight = weight.replace('(', '')
        weight = weight.replace(')', '')
        weight = int(weight)
        reports.append(Report(name, weight, children))
    return reports


def test_parse_lines():
    input = """pbga (66)
    xhth (57)
    ebii (61)
    havc (66)
    ktlj (57)
    fwft (72) -> ktlj, cntj, xhth
    qoyq (66)
    padx (45) -> pbga, havc, qoyq
    tknk (41) -> ugml, padx, fwft
    jptl (61)
    ugml (68) -> gyxo, ebii, jptl
    gyxo (61)
    cntj (57)"""
    expected_parsed = [
        Report('pbga', 66, set()),
        Report('xhth', 57, set()),
        Report('ebii', 61, set()),
        Report('havc', 66, set()),
        Report('ktlj', 57, set()),
        Report('fwft', 72, set(['ktlj', 'cntj', 'xhth'])),
        Report('qoyq', 66, set()),
        Report('padx', 45, set(['pbga', 'havc', 'qoyq'])),
        Report('tknk', 41, set(['ugml', 'padx', 'fwft'])),
        Report('jptl', 61, set()),
        Report('ugml', 68, set(['gyxo', 'ebii', 'jptl'])),
        Report('gyxo', 61, set()),
        Report('cntj', 57, set()),
    ]
    assert parse_lines(input.splitlines()) == expected_parsed


def test_find_root():
    reports = [
        Report('pbga', 66, set()),
        Report('xhth', 57, set()),
        Report('ebii', 61, set()),
        Report('havc', 66, set()),
        Report('ktlj', 57, set()),
        Report('fwft', 72, set(['ktlj', 'cntj', 'xhth'])),
        Report('qoyq', 66, set()),
        Report('padx', 45, set(['pbga', 'havc', 'qoyq'])),
        Report('tknk', 41, set(['ugml', 'padx', 'fwft'])),
        Report('jptl', 61, set()),
        Report('ugml', 68, set(['gyxo', 'ebii', 'jptl'])),
        Report('gyxo', 61, set()),
        Report('cntj', 57, set()),
    ]
    assert find_root_name(reports) == 'tknk'


def test_towers():
    reports = [
        Report('b', 1, []),
        Report('a', 3, ['b', 'c']),
        Report('c', 2, []),
    ]

    towers = Towers(reports)
    assert len(towers.roots) == 1
    assert 'b' in towers.roots['a'].children
    assert 'c' in towers.roots['a'].children
    assert towers.roots['a'].children['b'] == Program('b', 1, {})
    assert towers.roots['a'].children['c'] == Program('c', 2, {})


lines = open("inputs/day_07_input.txt").read().splitlines()
reports = parse_lines(lines)
print(f"Part one: {find_root_name(reports)}")

#pytest.main(f"-v --tb=short {__file__}")
#pytest.main(__file__)
