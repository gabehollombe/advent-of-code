from collections import namedtuple, defaultdict
from copy import copy

import pytest

Instruction = namedtuple('Instruction', 'register operation operation_value comparison_register comparison comparison_value')
COMPARISONS = {
    '>': lambda a, b: a > b,
    '<': lambda a, b: a < b,
    '>=': lambda a, b: a >= b,
    '<=': lambda a, b: a <= b,
    '==': lambda a, b: a == b,
    '!=': lambda a, b: a != b,
}

OPERATIONS = {
    'inc': lambda n, amount: n + amount,
    'dec': lambda n, amount: n - amount,
}

def parse_instruction(instruction_string):
    tokens = instruction_string.split(' ')
    register, operation, operation_value, _, comparison_register, comparison, comparison_value = tokens
    operation_value = int(operation_value)
    comparison_value = int(comparison_value)
    return Instruction(register, operation, operation_value, comparison_register, comparison, comparison_value)


def execute_instruction(registers, instruction):
    comparison = COMPARISONS[instruction.comparison]
    a = registers[instruction.comparison_register]
    b = instruction.comparison_value
    if not comparison(a, b):
        return registers

    registers_after_execution = copy(registers)
    operation = OPERATIONS[instruction.operation]
    register_value = registers[instruction.register]
    registers_after_execution[instruction.register] = operation(register_value, instruction.operation_value)

    if registers_after_execution[instruction.register] > registers['__highest_value_seen__']:
        registers_after_execution['__highest_value_seen__'] = registers_after_execution[instruction.register]

    return registers_after_execution


def test_execute_instruction():
    registers = defaultdict(int)

    assert registers['b'] == 0
    registers = execute_instruction(registers, parse_instruction('b inc 5 if a > 1'))
    assert registers['b'] == 0

    assert registers['a'] == 0
    registers = execute_instruction(registers, parse_instruction('a inc 1 if b < 5'))
    assert registers['a'] == 1

    assert registers['c'] == 0
    registers = execute_instruction(registers, parse_instruction('c dec -10 if a >= 1'))
    assert registers['c'] == 10

    registers = execute_instruction(registers, parse_instruction('c inc -20 if c == 10'))
    assert registers['c'] == -10

    assert registers['d'] == 0
    registers = execute_instruction(registers, parse_instruction('d inc 2 if c <= 10'))
    assert registers['d'] == 2

    registers = execute_instruction(registers, parse_instruction('d inc 2 if d != 0'))
    assert registers['d'] == 4

    assert registers['__highest_value_seen__'] == 10


def test_parse_instruction():
    assert parse_instruction("a inc 123 if b < 456") == Instruction('a', 'inc', 123, 'b', '<', 456)


lines = open("inputs/day_08_input.txt").read().splitlines()

instructions = map(parse_instruction, lines)
registers = defaultdict(int)
for instruction in instructions:
    registers = execute_instruction(registers, instruction)

print(f"Part one: {max(registers.values())}")
print(f"Part two: {registers['__highest_value_seen__']}")

pytest.main(['--color=yes', '--tb=short', __file__])