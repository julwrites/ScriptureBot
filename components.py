
# Local modules
from bible import bible_actions, bible_hooks
from tms import tms_actions, tms_hooks
from admin import admin_actions
from user import user_actions
from devo import devo_hooks

def actions():
    return tms_actions.get() + \
            bible_actions.get() + \
            user_actions.get() + \
            admin_actions.get()

def hooks():
    return tms_hooks.get() + \
            bible_hooks.get() + \
            devo_hooks.get()

def subscriptions():
    subList = []

    for hook in hooks():
        subList.append(hook.name())

    return subList

