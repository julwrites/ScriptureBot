
# coding=utf-8

# Local modules
from dailydevo import mcheyne_hooks, mcheyne_actions

def get_actions():
    return mcheyne_actions.get()

def get_hooks():
    return mcheyne_hooks.get()