
# Local modules
from devo import subscribe_actions
from devo import mcheyne_actions, mcheyne_hooks
from devo import cac_hooks

def get_actions():
    return subscribe_actions.get() + \
            mcheyne_actions.get()

def get_hooks():
    return mcheyne_hooks.get()