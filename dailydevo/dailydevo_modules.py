# coding=utf-8

# Local modules
from dailydevo import djbr_hooks, djbr_actions
from dailydevo import mcheyne_hooks, mcheyne_actions
from dailydevo import cac_hooks, cac_actions
from dailydevo import desiringgod_hooks, desiringgod_actions
from dailydevo import odb_hooks, odb_actions


def get_actions():
    return mcheyne_actions.get() + \
            djbr_actions.get() + \
            desiringgod_actions.get() + \
            cac_actions.get() + \
            odb_actions.get()


def get_hooks():
    return mcheyne_hooks.get() + \
            djbr_hooks.get() + \
            desiringgod_hooks.get() + \
            odb_hooks.get()
