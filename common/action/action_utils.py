
# coding=utf-8

# Local modules
from common import debug

def clear():
    debug.log("Clearing the user's state and keyboard")

    telegram_utils.send_close_keyboard(confirmString, userObj.get_uid())
    userObj.set_state(None)


def execute(actions, userObj, msg):
    try:
        debug.log("Trying actions: " + "|".join([action.identifier() for action in actions]))

        matched = [action for action in actions if action.match(msg)]

        if len(matched) > 0:
            clear()

        for action in matched:
            debug.log_action(action.identifier())
            if action.resolve(userObj, msg):
                return True

        waiting = [action for action in actions if action.waiting(userObj)]

        for action in waiting:
            debug.log_action(action.identifier())
            if action.resolve(userObj, msg):
                return True

        return False
    except Exception as e:
        debug.log("Execute failed!")
        debug.err(e)
    return False
