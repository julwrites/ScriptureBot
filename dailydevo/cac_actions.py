# coding=utf8

# Local modules
from common import debug
from common.action import action_classes
from common.telegram import telegram_utils

from dailydevo import cac_utils

class CACDevoAction(action_classes.Action):
    def identifier(self):
        return "/cacdevo"

    def name(self):
        return "Center for Action and Contemplation Devotional"

    def resolve(self, userObj):
        debug.log("Resolving CAC Devo Action")
        passage = cac_utils.get_cacdevo(userObj.get_version())

        debug.log("Got CAC Devo Passage")

        if passage is not None:
            telegram_utils.send_msg(passage, userObj.get_uid())
            debug.log("Sent message")

        return True

def get():
    return [
        CACDevoAction()
    ]