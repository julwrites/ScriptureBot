# coding=utf-8

# Python modules
import string
import re

# Local modules
# from common import debug # This cannot be imported if we want debug to use text_utils


def to_utf8(value):
    return value.decode("utf-8")


def stringify(value):
    if not isinstance(value, str) and not isinstance(value, unicode):
        value = unicode(value)

    value = bytearray(value, "utf-8")
    return value.decode("utf-8")


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


def segment_text(text, delimiters):
    blocks = []

    index = 0
    while index <= len(text):
        disp = -1

        delim_disp = [text[index:].find(d) for d in delimiters]
        delim_disp.sort()
        for d in delim_disp:
            if d != -1:
                disp = d
                break

        if disp == -1:
            break

        pos = index + disp + 1
        try_block = text[index:pos].strip()

        if try_block:
            blocks.append([index, pos])

        index = pos + 1

    return blocks


def detect_paragraphs(text):
    return segment_text(text, ["\n"])


def detect_sentences(text):
    return segment_text(text, [".", "!", "?"])