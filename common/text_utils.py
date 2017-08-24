
# Python modules
import string

# Local modules

def stringify(value):
    if value is None:
        return ''

    return str(value)

def fuzzify(s):
    return ''.join(s.upper().strip().replace('-', '').replace(',', '').split())

def fuzzy_compare(lhs, rhs):
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