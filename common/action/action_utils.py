
# coding=utf-8

# Local modules
from common import debug

def execute(actions, userObj, msg):
    if userObj is not None:
        waiting = [action for action in actions if action.waiting(userObj)]

        if len(waiting) == 1:
            debug.log_action(waiting.identifier())
            return waiting.resolve(userObj, msg)

        matched = [action for action in actions if action.match(msg)]

        if len(matched) == 1:
            debug.log_action(matched.identifier())
            return waiting.resolve(userObj, msg)
        return False

    try:
        if userObj is not None:
            waiting = [action for action in actions if action.waiting(userObj)]

            if len(waiting) == 1:
                debug.log_action(waiting.identifier())
                return waiting.resolve(userObj, msg)

            matched = [action for action in actions if action.match(msg)]

            if len(matched) == 1:
                debug.log_action(matched.identifier())
                return waiting.resolve(userObj, msg)
            return False
    except:
        debug.log('Execute failed!')
    return False
