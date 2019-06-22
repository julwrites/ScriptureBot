# coding=utf8

# Local modules
from common.utils import debug
from common.action import action_classes
from common.telegram import telegram_utils

from dailydevo import odb_hooks


class ODBDevoAction(action_classes.Action):
    def identifier(self):
        return "/odb"

    def name(self):
        return "Our Daily Bread"

    def description(self):
        return "Articles from Our Daily Bread"

    def resolve(self, userObj, msg):
        odb_hooks.ODBDevoHook().resolve(userObj)

        return True


def get():
    return [ODBDevoAction()]
