
# Local modules
from dailydevo import mcheyne_actions, mcheyne_hooks
from dailydevo import cac_hooks

def get_actions():
    return mcheyne_actions.get()

def get_hooks():
    return mcheyne_hooks.get()