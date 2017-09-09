
# coding=utf-8

# Local modules
from user import bibleuser_actions

from common.action import action_classes
from common.telegram import telegram_utils

CONFIRM = "Alright~"

class UserDoneAction(action_classes.Action):
    def identifier(self):
        return '/done'

    def name(self):
        return 'Done'

    def resolve(self, userObj, msg):
        telegram_utils.send_close_keyboard(CONFIRM, userObj.get_uid())
        userObj.set_state(None)

        return True

def get():
    return [
        UserDoneAction()
    ] + \
    bibleuser_actions.get()