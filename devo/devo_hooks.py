
# Local modules
from devo import cac_hooks, mcheyne_hooks


def get():
    return cac_hooks.get() + mcheyne_hooks.get()