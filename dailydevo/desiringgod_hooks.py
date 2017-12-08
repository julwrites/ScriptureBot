
# coding=utf-8

# Local modules
from common import debug, text_utils
from common.action import hook_classes
from common.telegram import telegram_utils

from dailydevo import desiringgod_utils
from user import user_actions

PROMPT = "Here are today's articles from desiringgod.org!\nTap on any one to get the article!"

class DGDevoHook(hook_classes.Hook):
    def identifier(self):
        return "/desiringgod"

    def name(self):
        return "Desiring God Articles"

    def description(self):
        return "Articles from DesiringGod.org"

    def resolve(self, userObj):
        refs = desiringgod_utils.get_desiringgod()

        if refs is not None:
            refs.append({"title":user_actions.UserDoneAction().name(), "link":""})
            options = [telegram_utils.make_button(text=text_utils.stringify(ref["title"]), fields={"url":text_utils.stringify(ref["link"])}) for ref in refs]

            telegram_utils.send_url_keyboard(PROMPT, userObj.get_uid(), options, 1)

def get():
    return [
        DGDevoHook()
    ]
