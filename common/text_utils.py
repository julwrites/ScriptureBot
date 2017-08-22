
# Python modules
import string

def stringify(value):
    if value is None:
        return ''

    return str(value)

def fuzzy_compare(lhs, rhs):
    lhs = lhs.upper().strip().split()
    lhs = ''.join(lhs)

    rhs = lhs.upper().strip().split()
    rhs = ''.join(lhs)

    return lhs == rhs