# coding=utf8

# Local modules
from common.utils import debug
from common.action import action_classes
from common.telegram import telegram_utils

from dailydevo import desiringgod_hooks


class DGDevoAction(action_classes.Action):
    def identifier(self):
        return "/desiringgod"

    def name(self):
        return "Desiring God Articles"

    def description(self):
        return "Articles from DesiringGod.org"

    def resolve(self, userObj, msg):
        desiringgod_hooks.DGDevoHook().resolve(userObj)

        return True


def get():
    return [DGDevoAction()]
