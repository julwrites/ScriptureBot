# coding=utf8

# Local modules
from common import debug
from common.action import action_classes
from common.telegram import telegram_utils

from dailydevo import desiringgod_utils


class DGDevoAction(action_classes.Action):
    def identifier(self):
        return "/desiringgod"

    def name(self):
        return "Desiring God Articles"

    def resolve(self, userObj, msg):
        passage = desiringgod_utils.get_desiringgoddevo(userObj.get_version())

        if passage is not None:
            telegram_utils.send_msg(passage, userObj.get_uid())

        return True


def get():
    return [DGDevoAction()]

