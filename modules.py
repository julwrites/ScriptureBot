
# Local modules
from bible import bible_actions, bible_hooks
from tms import tms_actions, tms_hooks
from admin import admin_actions
from user import user_actions
from subscribe import subscribe_actions, subscribe_hooks

def get_actions():
    return tms_actions.get() + \
            bible_actions.get() + \
            user_actions.get() + \
            subscribe_actions.get() + \
            admin_actions.get()

def get_hooks():
    return tms_hooks.get() + \
            bible_hooks.get() + \
            subscribe_hooks.get()
