# coding=utf-8

# Local modules
from common.utils import debug_utils
from common.action import action_classes


def execute(actions, userObj, msg):
    try:
        debug_utils.log(
            "Trying actions: {}",
            ["|".join([action.identifier() for action in actions])])

        # Execute in order:
        # Commands
        # Waiting
        # Names

        commands = [action for action in actions if action.match_command(msg)]

        for action in commands:
            debug_utils.log_action(action.identifier())
            if action.resolve(userObj, msg):
                return True

        waiting = [action for action in actions if action.waiting(userObj)]

        for action in waiting:
            debug_utils.log_action(action.identifier())
            if action.resolve(userObj, msg):
                return True

        matched = [action for action in actions if action.match(msg)]

        for action in matched:
            debug_utils.log_action(action.identifier())
            if action.resolve(userObj, msg):
                return True

        return False
    except Exception as e:
        debug_utils.log("Execute failed!")
        debug_utils.err(e)
    return False
