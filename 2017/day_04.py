from collections import Counter


def is_valid_passphrase(passphrase):
    words = passphrase.split(' ')
    return no_repeats(words)


def no_repeats(words):
    counts = Counter(words)
    _, num_most_common = counts.most_common(1)[0]
    return num_most_common == 1

def is_valid_passphrase_no_anagrams(passphrase):
    words = passphrase.split(' ')
    sorted_words = ["".join(sorted(w)) for w in words]
    return no_repeats(sorted_words)

def valid_passphrases(passphrases):
    return list(filter(is_valid_passphrase, passphrases))

def valid_passphrases_no_anagrams(passphrases):
    return list(filter(is_valid_passphrase_no_anagrams, passphrases))

def test_valid_passphrase():
    assert is_valid_passphrase("aa bb cc dd ee")
    assert not is_valid_passphrase("aa bb cc dd aa")
    assert is_valid_passphrase("aa bb cc dd aaa")

def test_valid_passphrase_no_anagrams():
    assert is_valid_passphrase_no_anagrams("abcde fghij")
    assert not is_valid_passphrase_no_anagrams("abcde xyz ecdab")
    assert is_valid_passphrase_no_anagrams("a ab abc abd abf abj")
    assert is_valid_passphrase_no_anagrams("iiii oiii ooii oooi oooo")
    assert not is_valid_passphrase_no_anagrams("oiii ioii iioi iiio")


passphrases = open("inputs/day_04_input.txt").read().splitlines()
print(f"Part one: {len(valid_passphrases(passphrases))}")
print(f"Part two: {len(valid_passphrases_no_anagrams(passphrases))}")
