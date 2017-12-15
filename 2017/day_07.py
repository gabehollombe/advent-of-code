from collections import namedtuple
from functools import reduce
from typing import List
import pytest

Report = namedtuple('Report', ['name', 'weight', 'names_of_children'])

class Program():
    def __init__(self, name, weight, names_of_children=None):
        self.name = name
        self.weight = weight
        self.names_of_children = names_of_children or []
        self.children = []

    def add_child(self, child):
        self.children.append(child)

    def __eq__(self, other):
        return (
                self.name == other.name
                and
                self.weight == other.weight
                and
                self.names_of_children == other.names_of_children
        )


class Towers():
    def __init__(self, reports: List[Report]):
        self.programs = dict()
        # Build programs and track names of children
        for report in reports:
            program = Program(report.name, report.weight, {child_name: None for child_name in report.names_of_children})
            self.programs[program.name] = program

        # Update children name references into actual programs
        for program in self.programs.values():
            for child_name in program.names_of_children:
                program.add_child(self.programs[child_name])

        # Figure out which node(s) are root nodes
        child_names = set([name for p in self.programs.values() for name in p.names_of_children])
        all_names = set(self.programs.keys())
        self.roots = {name: self.programs[name] for name in (all_names - child_names)}

    def recursive_weight(self, program):
        return program.weight + self.weight_of_all_decendants(program)

    def weight_of_all_decendants(self, program):
        return sum([self.recursive_weight(c) for c in program.children])

    def find(self, name):
        return self.programs[name]

    def is_unbalanced(self, program):
        return len(set([self.recursive_weight(p) for p in program.children])) > 1

    def find_deepest_unbalanced_node(self, node):
        unbalanced_children = list(filter(lambda n: self.is_unbalanced(n), node.children))
        if len(unbalanced_children) == 0:
            return node
        else:
            unbalanced_child = unbalanced_children[0]
            return self.find_deepest_unbalanced_node(unbalanced_child)

    def find_wrong_weight(self):
        program = self.find_deepest_unbalanced_node(self.root)
        children = program.children
        total_weights = set([self.recursive_weight(p) for p in children])
        difference = abs(reduce(lambda diff, w: diff - w, total_weights))
        adjustments = [difference, -difference]

        for child in children:
            for adjustment in adjustments:
                original_weight = child.weight
                adjusted_weight = original_weight + adjustment
                self.programs[child.name].weight = adjusted_weight
                if not self.is_unbalanced(program):
                    return child.name, child.weight
                else:
                    self.programs[child.name].weight = original_weight


    @property
    def root(self) -> Program:
        return list(self.roots.values())[0]


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
        Report('fwft', 72, {'ktlj', 'cntj', 'xhth'}),
        Report('qoyq', 66, set()),
        Report('padx', 45, {'pbga', 'havc', 'qoyq'}),
        Report('tknk', 41, {'ugml', 'padx', 'fwft'}),
        Report('jptl', 61, set()),
        Report('ugml', 68, {'gyxo', 'ebii', 'jptl'}),
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
        Report('fwft', 72, {'ktlj', 'cntj', 'xhth'}),
        Report('qoyq', 66, set()),
        Report('padx', 45, {'pbga', 'havc', 'qoyq'}),
        Report('tknk', 41, {'ugml', 'padx', 'fwft'}),
        Report('jptl', 61, set()),
        Report('ugml', 68, {'gyxo', 'ebii', 'jptl'}),
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
    assert 'b' in towers.roots['a'].names_of_children
    assert 'c' in towers.roots['a'].names_of_children
    assert Program('b', 1) in towers.roots['a'].children
    assert Program('c', 2) in towers.roots['a'].children


def test_find_wrong_weight():
    reports = [
        Report('pbga', 66, set()),
        Report('xhth', 57, set()),
        Report('ebii', 61, set()),
        Report('havc', 66, set()),
        Report('ktlj', 57, set()),
        Report('fwft', 72, {'ktlj', 'cntj', 'xhth'}),
        Report('qoyq', 66, set()),
        Report('padx', 45, {'pbga', 'havc', 'qoyq'}),
        Report('tknk', 41, {'ugml', 'padx', 'fwft'}),
        Report('jptl', 61, set()),
        Report('ugml', 68, {'gyxo', 'ebii', 'jptl'}),
        Report('gyxo', 61, set()),
        Report('cntj', 57, set()),
    ]
    towers = Towers(reports)

    assert towers.is_unbalanced(towers.find('tknk'))
    assert not towers.is_unbalanced(towers.find('padx'))




lines = open("inputs/day_07_input.txt").read().splitlines()
reports = parse_lines(lines)
towers = Towers(reports)
print(f"Part one: {find_root_name(reports)}")
print(f"Part two: {towers.find_wrong_weight()}")

pytest.main([__file__])
