# coding=utf-8

# Python modules
import random

# Local modules
from user import bibleuser_actions

from common.action import action_classes
from common.telegram import telegram_utils
from common import debug, text_utils

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
        telegram_utils.send_reply(
            user=userObj.get_uid(),
            text=userObj.get_reply_string(CONFIRM),
            reply=telegram_utils.make_close_keyboard())

        userObj.set_state(None)

        return True


def get():
    return [
        UserDoneAction()
    ] + \
    bibleuser_actions.get()