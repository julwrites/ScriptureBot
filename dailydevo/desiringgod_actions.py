# coding=utf8

# Local modules
from common import debug
from common.action import action_classes
from common.telegram import telegram_utils

from dailydevo import desiringgod_utils
from user import user_actions

class DGDevoAction(action_classes.Action):
    def identifier(self):
        return "/desiringgod"

    def name(self):
        return "Desiring God Articles"

    def description(self):
        return "Articles from DesiringGod.org"

    def resolve(self, userObj, msg):
        query = telegram_utils.strip_command(msg, self.identifier())

        doneAction = user_actions.UserDoneAction()
        if doneAction.try_execute(userObj, msg):
            return True

        refs = desiringgod_utils.get_desiringgod()

        if refs is not None:
            options = refs
            options.append({"text":doneAction.name(), "url":None})

            telegram_utils.send_url_keyboard("", userObj.get_uid(), options, 1)
            userObj.set_state(self.identifier())
        
        return True


def get():
    return [DGDevoAction()]

