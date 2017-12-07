
# coding=utf-8

# Local modules
from common import debug, text_utils
from common.action import hook_classes
from common.telegram import telegram_utils

from dailydevo import desiringgod_utils

class DGDevoHook(hook_classes.Hook):
    def identifier(self):
        return "/desiringgod"

    def name(self):
        return "Desiring God Articles"

    def resolve(self, userObj):
        passage = desiringgod_utils.get_desiringgoddevo(userObj.get_version())

        if passage is not None:
            telegram_utils.send_msg(passage, userObj.get_uid())


def get():
    return [
        DGDevoHook()
    ]
