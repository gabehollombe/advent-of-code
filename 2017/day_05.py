def step(offsets, pointer_index):
    new_offsets = offsets.copy()
    new_offsets[pointer_index] = offsets[pointer_index] + 1
    new_pointer_index = pointer_index + offsets[pointer_index]
    return new_offsets, new_pointer_index


def step_decrease_if_three_or_more(offsets, pointer_index):
    #new_offsets = offsets.copy()
    offset = offsets[pointer_index]
    new_pointer_index = pointer_index + offsets[pointer_index]
    offsets[pointer_index] = offset - 1 if offset >= 3 else offset + 1
    return offsets, new_pointer_index


def test_step():
    offsets = [0, 3, 0, 1, -3]
    pointer_index = 0

    new_offsets, new_pointer_index = step(offsets, pointer_index)
    assert new_offsets[0] == 1, "expected offset at pointer_index to increment"
    assert new_pointer_index == 0, "expected new_pointer_index not to change"

    offsets, pointer_index = new_offsets, new_pointer_index
    new_offsets, new_pointer_index = step(offsets, pointer_index)
    assert new_offsets[0] == 2, "expected offset at pointer_index to increment"
    assert new_pointer_index == 1, "expected new_pointer_index to advance to 1"

    offsets, pointer_index = new_offsets, new_pointer_index
    new_offsets, new_pointer_index = step(offsets, pointer_index)
    assert new_offsets[1] == 4, "expected offset at pointer_index to increment"
    assert new_pointer_index == 4, "expected new_pointer_index to advance to 4"

    offsets, pointer_index = new_offsets, new_pointer_index
    new_offsets, new_pointer_index = step(offsets, pointer_index)
    assert new_offsets[4] == -2, "expected offset at pointer_index to increment"
    assert new_pointer_index == 1, "expected new_pointer_index to go back to 1"

    offsets, pointer_index = new_offsets, new_pointer_index
    new_offsets, new_pointer_index = step(offsets, pointer_index)
    assert new_offsets[1] == 5, "expected offset at pointer_index to increment"
    assert new_pointer_index == 5, "expected new_pointer_index to 5, escaping"




def test_step_decrease_if_three_or_more():
    offsets = [0, 3, 0, 1, -3]
    pointer_index = 1
    new_offsets, new_pointer_index = step_decrease_if_three_or_more(offsets, pointer_index)
    assert new_offsets[1] == 2, "expected offset at pointer_index to increment"
    assert new_pointer_index == 4, "expected new_pointer_index to advance to 4"


def count_steps_to_exit(offsets):
    steps = 0
    pointer_index = 0
    while pointer_index < len(offsets):
        offsets, pointer_index = step(offsets, pointer_index)
        steps = steps + 1
    return steps

def count_steps_to_exit_decrease_if_three_or_more(offsets):
    steps = 0
    pointer_index = 0
    while pointer_index < len(offsets):
        offsets, pointer_index = step_decrease_if_three_or_more(offsets, pointer_index)
        steps = steps + 1
    return steps


def test_count_steps_to_exit():
    offsets = [0, 3, 0, 1, -3]
    assert count_steps_to_exit(offsets) == 5

def test_count_steps_to_exit_decrease_if_three_or_more():
    offsets = [0, 3, 0, 1, -3]
    assert count_steps_to_exit_decrease_if_three_or_more(offsets) == 10


offsets = [int(offset) for offset in open("inputs/day_05_input.txt").read().splitlines()]
print(f"Part one: {count_steps_to_exit(offsets)}")
print(f"Part two: {count_steps_to_exit_decrease_if_three_or_more(offsets)}")
