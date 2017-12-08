
# coding=utf-8

# Local modules
from common import debug, text_utils
from common.action import hook_classes
from common.telegram import telegram_utils

from dailydevo import desiringgod_actions
from user import user_actions

class DGDevoHook(hook_classes.Hook):
    def identifier(self):
        return "/desiringgod"

    def name(self):
        return "Desiring God Articles"

    def description(self):
        return "Articles from DesiringGod.org"

    def resolve(self, userObj):
        desiringgod_actions.DGDevoAction().resolve(userObj, "")

def get():
    return [
        DGDevoHook()
    ]
