# coding=utf8

# Local modules
from common import debug
from common.action import action_classes
from common.telegram import telegram_utils

from dailydevo import desiringgod_utils
from user import user_actions

PROMPT = "Here are today's articles from desiringgod.org!\nTap on any one to get the article!"

class DGDevoAction(action_classes.Action):
    def identifier(self):
        return "/desiringgod"

    def name(self):
        return "Desiring God Articles"

    def description(self):
        return "Articles from DesiringGod.org"

    def resolve(self, userObj, msg):
        refs = desiringgod_utils.get_desiringgod()

        if refs is not None:
            debug.log("Appending done")
            refs.append({"title":user_actions.UserDoneAction.name(), "link":""})
            debug.log("Converting to buttons")
            options = [telegram_utils.make_button(text=ref["title"], fields={"url":ref["link"]}) for ref in refs]
            debug.log("Got the buttons: " + options)

            telegram_utils.send_url_keyboard(PROMPT, userObj.get_uid(), options, 1)

        return True


def get():
    return [DGDevoAction()]

