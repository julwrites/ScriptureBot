# coding=utf-8

# Local modules
from common import debug


def execute(actions, userObj, msg):
    try:
        debug.log("Trying actions: " +
                  "|".join([action.identifier() for action in actions]))

        # Execute in order:
        # Commands
        # Waiting
        # Names

        commands = [action for action in actions if action.match_command(msg)]

        for action in commands:
            debug.log_action(action.identifier())
            if action.resolve(userObj, msg):
                return True

        waiting = [action for action in actions if action.waiting(userObj)]

        for action in waiting:
            debug.log_action(action.identifier())
            if action.resolve(userObj, msg):
                return True

        matched = [action for action in actions if action.match(msg)]

        for action in matched:
            debug.log_action(action.identifier())
            if action.resolve(userObj, msg):
                return True

        return False
    except Exception as e:
        debug.log("Execute failed!")
        debug.err(e)
    return False
