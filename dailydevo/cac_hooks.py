
# coding=utf-8

# Local modules
from common import debug, text_utils
from common.action import hook_classes
from common.telegram import telegram_utils

from dailydevo import cac_utils

class CACDevoHook(hook_classes.Hook):
    def identifier(self):
        return "/cacdevo"

    def name(self):
        return "Center for Action and Contemplation Devotional"

    def resolve(self, userObj):
        passage = cac_utils.get_cacdevo(userObj.get_version())

        if passage is not None:
            telegram_utils.send_msg(userObj.get_uid(), passage)


def get():
    return [
        CACDevoHook()
    ]