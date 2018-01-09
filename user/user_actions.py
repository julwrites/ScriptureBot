# coding=utf-8

# Python modules
import random

# Local modules
from user import bibleuser_actions

from common.action import action_classes
from common.telegram import telegram_utils
from common import text_utils

CONFIRM = [
    "Alright {}~",
    "Okay {}~",
    "Got it, {}~",
    "Yes, yes, {}~",
    "I understand, {}~",
    "If you say so {}~",
    "Done, {}~",
]


class UserDoneAction(action_classes.Action):
    def identifier(self):
        return "/done"

    def name(self):
        return "Done"

    def resolve(self, userObj, msg):
        choose = random.randint(0, len(CONFIRM) - 1)
        confirmString = CONFIRM[choose].format(
            text_utils.stringify(userObj.get_name_string()))

        telegram_utils.send_reply(
            user=userObj.get_uid(),
            text=confirmString,
            reply=telegram_utils.make_close_keyboard())
        userObj.set_state(None)

        return True


def get():
    return [
        UserDoneAction()
    ] + \
    bibleuser_actions.get()