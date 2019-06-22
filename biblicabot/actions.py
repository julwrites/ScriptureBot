# coding=utf-8

# Python std modules
import json

# Local modules
from .common.utils import debug_utils, text_utils
from .common.action import action_classes
from .common.telegram import telegram_utils

from .user import user_utils

from . import modules

INTRO_MSG = "Hello {}, I'm Biblica Bot! Give me a Bible reference and I'll give you the passage!"
COMMAND_MSG = "Here are some other things I can do:"
COMMAND_LIST = "\n".join([(action.identifier() + " - " + action.description())
                          for action in modules.get_actions()
                          if action.is_command()])
HELP_MSG = INTRO_MSG + "\n\n" + COMMAND_MSG + "\n" + COMMAND_LIST


class StartAction(action_classes.Action):
    def identifier(self):
        return "/start"

    def name(self):
        return "Start Bot"

    def description(self):
        return "Start the Bot"

    def match_command(self, msg):
        return msg is not None

    def resolve(self, userObj, msg):
        # Register User
        userJson = msg.get("from")
        userId = user_utils.get_uid(userJson.get("id"))

        # This runs to update the user"s info, or register
        if userJson is not None:
            debug_utils.log("Updating user info")
            user_utils.set_profile(
                userJson.get("id"), userJson.get("username"),
                userJson.get("first_name"), userJson.get("last_name"))
            debug_utils.log("Finished updating user info")

        # If this is the user"s first time registering
        if userObj is None:
            userObj = user_utils.get_user(userId)
            debug_utils.log("Sending user help")

            HelpAction().resolve(userObj, msg)

            return True
        return False


class HelpAction(action_classes.Action):
    def identifier(self):
        return "/help"

    def name(self):
        return "Help"

    def description(self):
        return "Get the Bot's Help menu"

    def match(self, msg):
        return msg is not None

    def resolve(self, userObj, msg):
        telegram_utils.send_msg(
            user=userObj.get_uid(),
            text=HELP_MSG,
            args=[userObj.get_name_string()])
        return True


def get():
    return [
        StartAction(),
    ] + \
    modules.get_actions() + \
    [
        HelpAction(),
    ]