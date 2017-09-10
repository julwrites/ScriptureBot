
# coding=utf-8

# Local modules
from common import debug

def execute(actions, userObj, msg):
    try:
        debug.log("Trying actions: " + "|".join([action.identifier() for action in actions]))
        waiting = [action for action in actions if action.waiting(userObj)]

        if len(waiting) == 1:
            action = waiting[0]
            debug.log_action(action.identifier())
            return action.resolve(userObj, msg)

        matched = [action for action in actions if action.match(msg)]

        if len(matched) == 1:
            action = matched[0]
            debug.log_action(action.identifier())
            return action.resolve(userObj, msg)
        return False
    except:
        debug.log("Execute failed!")
    return False
