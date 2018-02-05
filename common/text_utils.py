# coding=utf-8

# Python modules
import string
import re

# Local modules
from common import debug


def utf8(value):
    return bytearray(value, "utf-8").decode("utf-8")


def stringify(value):
    return utf8(unicode(value))


def is_valid(s):
    if s is not None:
        if s.strip():
            return True
    return False


def fuzzify_raw(s):
    return s.upper().replace("-", " ").replace(",", " ").strip().split()


def fuzzify_join(parts):
    return "".join(parts)


def fuzzify(s):
    return fuzzify_join(fuzzify_raw(s))


def overlap(lhsSub, rhsSub):
    for lhs in lhsSub:
        for rhs in rhsSub:
            if lhs == rhs:
                return True
    return False


def find_alpha(blocks):
    for i in range(len(blocks)):
        if blocks[i].isalpha():
            return i
    return -1


def fuzzy_compare(lhs, rhs):
    lhsParts = fuzzify_raw(lhs)
    rhsParts = fuzzify_raw(rhs)
    lhs = fuzzify_join(lhsParts)
    rhs = fuzzify_join(rhsParts)

    return (lhs == rhs or overlap(lhsParts, rhsParts))


def overlap_compare(lhs, rhs):
    lhsParts = fuzzify_raw(lhs)
    rhsParts = fuzzify_raw(rhs)

    return overlap(lhsParts, rhsParts)


def text_compare(lhs, rhs):
    return fuzzify(lhs) == fuzzify(rhs)


def strip_whitespace(s):
    s = s.strip().split()
    return " ".join(s)


def strip_numbers(s):
    result = "".join([c for c in s if not c.isdigit()])
    return result if len(result) > 0 else None


def strip_alpha(s):
    result = "".join([c for c in s if not c.isalpha()])
    return result if len(result) > 0 else None


def strip_block(s, start, end):
    start = s.find(start)
    end = s.find(end)
    start_block = s[:start].strip()
    end_block = s[end + 1:].strip()
    return start_block + end_block


def replace(s, sub, new):
    return s.replace(sub, new)


def replace_newline(s):
    return "\n"