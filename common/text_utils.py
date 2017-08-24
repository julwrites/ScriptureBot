
# Python modules
import string

# Local modules
import debug

def stringify(value):
    if value is None:
        return ''

    return str(value)

def is_valid(s):
    if s is not None:
        if len(s.strip()) > 0:
            return True
    return False

def fuzzify_raw(s):
    return s.upper().replace('-', ' ').replace(',', ' ').strip().split()

def fuzzify_join(parts):
    return ''.join(parts)

def fuzzify(s):
    return fuzzify_join(fuzzify_raw(s))

def overlap(lhs_sub, rhs_sub):
    for lhs in lhs_sub:
        for rhs in rhs_sub:
            if lhs == rhs:
                return True
    return False

def fuzzy_compare(lhs, rhs):
    lhs_parts = fuzzify_raw(lhs)
    rhs_parts = fuzzify_raw(rhs)
    lhs = fuzzify_join(lhs_parts)
    rhs = fuzzify_join(rhs_parts)

    return ( lhs == rhs or overlap(lhs_parts, rhs_parts) )

def text_compare(lhs, rhs):
    return fuzzify(lhs) == fuzzify(rhs)

def strip_whitespace(s):
    s = s.strip().split()
    return ' '.join(s)

def strip_numbers(s):
    result = ''.join([c for c in s if not c.isdigit()])
    return result if len(result) > 0 else None

def strip_alpha(s):
    result = ''.join([c for c in s if not c.isalpha()])
    return result if len(result) > 0 else None