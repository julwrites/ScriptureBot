
# coding=utf-8

# Python std modules
import json

# Local modules
from common import debug
from common.action import action_classes
from common.telegram import telegram_utils

from user import user_utils

import modules

INTRO_MSG = "Hello {}, I\'m Biblica Bot! I hope I will be helpful as a tool for you to handle the Bible!"
COMMAND_MSG = "Here are some things I can do:"
COMMAND_LIST = '\n'.join([(action.identifier() + ' - ' + action.description()) for action in modules.get_actions() if action.is_command()])
HELP_MSG = INTRO_MSG + '\n\n' + COMMAND_MSG + '\n' + COMMAND_LIST

class StartAction(action_classes.Action):
    def identifier(self):
        return "/start"

    def name(self):
        return "Start Bot"

    def description(self):
        return "Start the Bot"

    def execute(self, userObj, msg):
        return self.resolve(userObj, msg)

    def resolve(self, userObj, msg):
        # Register User
        userJson = msg.get('from')
        userId = user_utils.get_uid(userJson.get('id'))

        # This runs to update the user's info, or register
        if userJson is not None:
            debug.log('Updating user info')
            user_utils.set_profile(
                userJson.get('id'), 
                userJson.get('username'), 
                userJson.get('first_name'), 
                userJson.get('last_name'))

        # If this is the user's first time registering
        if userObj is None:
            debug.log_cmd('New user!')
            userObj = user_utils.get_user(userId)

            telegram_utils.send_msg(INTRO_MSG.format(userObj.get_name_string()), userObj.get_uid())
            debug.log('Registering ' + userObj.get_name_string())

            return True
        return False

class HelpAction(action_classes.Action):
    def identifier(self):
        return "/help"

    def name(self):
        return "Help"

    def description(self):
        return "Get the Bot's Help menu"

    def resolve(self, userObj, msg):
        telegram_utils.send_msg(HELP_MSG.format(userObj.get_name_string()), userObj.get_uid())
        return True

def get():
    return [
        StartAction(),
        HelpAction()
    ] + \
    modules.get_actions()