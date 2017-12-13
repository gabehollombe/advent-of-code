from copy import copy


def highest_value_and_index(l):
    index = 0
    highest_val_so_far = l[0]
    for i in range(1, len(l)):
        val = l[i]
        if val > highest_val_so_far:
            highest_val_so_far = val
            index = i
    return highest_val_so_far, index

def reallocate(banks):
    blocks_remaining, index_to_reallocate= highest_value_and_index(banks)
    banks[index_to_reallocate] = 0
    next_index = index_to_reallocate + 1
    while blocks_remaining > 0:
        wrapped_index = next_index % len(banks)
        banks[wrapped_index] += 1
        blocks_remaining -= 1
        next_index += 1
    return banks


def test_index_of_highest_val():
    assert highest_value_and_index([22,33,11,33]) == (33, 1)
    assert highest_value_and_index([22,33,44,11]) == (44, 2)


def test_reallocate():
    assert reallocate([0, 2, 7, 0]) == [2, 4, 1, 2]
    assert reallocate([2, 4, 1, 2]) == [3, 1, 2, 3]
    assert reallocate([3, 1, 2, 3]) == [0, 2, 3, 4]
    assert reallocate([0, 2, 3, 4]) == [1, 3, 4, 1]
    assert reallocate([1, 3, 4, 1]) == [2, 4, 1, 2]


def hash_list(l):
    return ",".join([str(i) for i in l])


def reallocate_until_duplicate(banks):
    previous_hashes = dict()
    reallocated = reallocate(banks)
    count = 1
    while hash_list(reallocated) not in previous_hashes:
        previous_hashes[hash_list(reallocated)] = count
        reallocated = reallocate(banks)
        count += 1
    length_of_loop = count - previous_hashes[hash_list(reallocated)]
    return banks, count, length_of_loop


def test_hash_list():
    assert hash_list([3,2,1]) != hash_list([1,2,3])
    assert hash_list([3,2,1]) == hash_list([3,2,1])

def test_count_reallocations_until_duplicate():
    assert reallocate_until_duplicate([0, 2, 7, 0]) == ([2, 4, 1, 2], 5, 4)

input = "5	1	10	0	1	7	13	14	3	12	8	10	7	12	0	6"
banks = [int(i) for i in input.split("\t")]

_, reallocations_count, length_of_loop = reallocate_until_duplicate(banks)
print(f"Part one: {reallocations_count}")
print(f"Part two: {length_of_loop}")
