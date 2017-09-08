
# Local modules
from dailydevo import mcheyne_hooks
from dailydevo import cac_hooks

def get_hooks():
    return mcheyne_hooks.get() + \
            cac_hooks.get()