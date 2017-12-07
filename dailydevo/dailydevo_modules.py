# coding=utf-8

# Local modules
from dailydevo import mcheyne_hooks, mcheyne_actions
from dailydevo import cac_hooks, cac_actions
from dailydevo import desiringgod_hooks, desiringgod_actions


def get_actions():
    return mcheyne_actions.get() + \
            desiringgod_actions.get() + \
            cac_actions.get()


def get_hooks():
    return mcheyne_hooks.get() + \
            desiringgod_hooks.get()