
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
    debug.log('lhs stripped to: ' + lhs)

    rhs = lhs.upper().strip().split()
    rhs = ''.join(rhs)
    debug.log('rhs stripped to: ' + rhs)

    return lhs == rhs