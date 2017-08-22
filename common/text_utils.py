
# Python modules
import string

# Local modules
from common import debug

def stringify(value):
    if value is None:
        return ''

    return str(value)

def fuzzy_compare(lhs, rhs):
    lhs = lhs.upper().strip().split()
    lhs = ''.join(lhs)

    rhs = rhs.upper().strip().split()
    rhs = ''.join(rhs)

    return lhs == rhs