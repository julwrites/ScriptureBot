
# Local modules
from subscribe import subscription_actions
from subscribe import mcheyne_actions, mcheyne_hooks
from subscribe import cac_hooks

def get_actions():
    return subscription_actions.get() + \
            mcheyne_actions.get()

def get_hooks():
    return mcheyne_hooks.get()