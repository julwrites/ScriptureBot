# coding=utf8

# Local modules
from common.utils import debug_utils
from common.action import action_classes
from common.telegram import telegram_utils

from dailydevo import cac_hooks


class CACDevoAction(action_classes.Action):
    def identifier(self):
        return "/cacdevo"

    def name(self):
        return "Center for Action and Contemplation Devotional"

    def resolve(self, userObj, msg):
        cac_hooks.CACDevoHook().resolve(userObj)

        return True


def get():
    return [CACDevoAction()]
