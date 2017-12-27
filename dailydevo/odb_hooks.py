# coding=utf-8

# Local modules
from common import debug, text_utils
from common.action import hook_classes
from common.telegram import telegram_utils

from dailydevo import odb_utils
from user import user_actions


class ODBDevoHook(hook_classes.Hook):
    def identifier(self):
        return "/odb"

    def name(self):
        return "Our Daily Bread"

    def description(self):
        return "Articles from Our Daily Bread"

    def resolve(self, userObj):
        passage = odb_utils.get_odb()

        if passage is not None:
            telegram_utils.send_msg(user=userObj.get_uid(), text=passage)


def get():
    return [ODBDevoHook()]
