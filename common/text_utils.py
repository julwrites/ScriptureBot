
# Python modules
import string

# Local modules

def stringify(value):
    if value is None:
        return ''

    return str(value)

def fuzzify(s):
    return s.upper().strip().replace('-', ' ').replace(',', ' ').split()

def fuzzy_find(s, substrings):
    return len([ss for ss in substrings if s.find(ss) is not -1]) > 0

def fuzzy_compare(lhs, rhs):
    lhs_parts = fuzzify(lhs)
    rhs_parts = fuzzify(rhs)
    lhs = ''.join(lhs_parts)
    rhs = ''.join(rhs_parts)

    return ( lhs == rhs \
        or fuzzy_find(lhs, rhs_parts) \
        or fuzzy_find(rhs, lhs_parts) \
    )

def strip_whitespace(s):
    s = s.strip().split()
    return ' '.join(s)

def strip_numbers(s):
    result = ''.join([c for c in s if not c.isdigit()])
    return result if len(result) > 0 else None

def strip_alpha(s):
    result = ''.join([c for c in s if not c.isalpha()])
    return result if len(result) > 0 else None