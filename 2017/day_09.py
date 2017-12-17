import pytest
import re


class StreamProcessor():
    def __init__(self):
        self.garbage_count = 0
        self.currently_in_garbage = False
        self.group_count = 0
        self.ignore_next_char = False
        self.score = 0

    @property
    def active_groups(self):
        return self.group_count

    def ingest(self, str):
        for c in str:
            self.ingest_char(c)

        return self

    def ingest_char(self, char):
        if self.ignore_next_char:
            self.ignore_next_char = False

        elif char == '!':
            self.ignore_next_char = True

        elif char == '<' and not self.currently_in_garbage:
            self.currently_in_garbage = True

        elif self.currently_in_garbage:
            if char == '>':
                self.currently_in_garbage = False
            else:
                self.garbage_count += 1

        elif char == '{':
            self.increment_group_count()

        elif char == '}':
            self.score += self.group_count
            self.decrement_group_count()

    def decrement_group_count(self):
        self.group_count -= 1

    def increment_group_count(self):
        self.group_count += 1


def test_tracks_braces():
    proc = StreamProcessor()
    assert proc.active_groups == 0

    proc.ingest('{')
    assert proc.active_groups == 1

    proc.ingest('{')
    assert proc.active_groups == 2

    proc.ingest('}')
    assert proc.active_groups == 1

    proc.ingest('}')
    assert proc.active_groups == 0

def test_ignores_groups_inside_garbage():
    proc = StreamProcessor()

    proc.ingest('{<{')
    assert proc.active_groups == 1

def test_resumes_group_tracking_after_garbage_ends():
    proc = StreamProcessor()

    proc.ingest('{<{<>{')
    assert proc.active_groups == 2

def test_ignores_char_following_bang():
    proc = StreamProcessor()

    proc.ingest('!{')
    assert proc.active_groups == 0

    proc.ingest('!!{')
    assert proc.active_groups == 1

def test_keeps_score():
    assert StreamProcessor().ingest('{}').score == 1
    assert StreamProcessor().ingest('{{{}}}').score == 6
    assert StreamProcessor().ingest('{{},{}}').score == 5
    assert StreamProcessor().ingest('{{{},{},{{}}}}').score == 16
    assert StreamProcessor().ingest('{<a>,<a>,<a>,<a>}').score == 1
    assert StreamProcessor().ingest('{{<ab>},{<ab>},{<ab>},{<ab>}}').score == 9
    assert StreamProcessor().ingest('{{<!!>},{<!!>},{<!!>},{<!!>}}').score == 9
    assert StreamProcessor().ingest('{{<a!>},{<a!>},{<a!>},{<ab>}}').score == 3

def test_counts_garbage():
    assert StreamProcessor().ingest('<>').garbage_count == 0
    assert StreamProcessor().ingest('<random characters>').garbage_count == 17
    assert StreamProcessor().ingest('<<<<>').garbage_count == 3
    assert StreamProcessor().ingest('<{!>}>').garbage_count == 2
    assert StreamProcessor().ingest('!!').garbage_count == 0
    assert StreamProcessor().ingest('<!!!>>').garbage_count == 0
    assert StreamProcessor().ingest('<{o"i!a,<{i<a>').garbage_count == 10


input = open('inputs/day_09_input.txt').readline()
proc = StreamProcessor().ingest(input)
print(f"Part one: {proc.score}")
print(f"Part two: {proc.garbage_count}")


pytest.main(['--color=yes', '--tb=short', __file__])
