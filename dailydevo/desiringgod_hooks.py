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
            options = [
                telegram_utils.make_inline_button(
                    text=ref["title"], url=ref["link"]) for ref in refs
            ]

            telegram_utils.send_reply(
                user=userObj.get_uid(),
                text=PROMPT,
                reply=telegram_utils.make_inline_keyboard(
                    buttons=options, width=1))


def get():
    return [
        DGDevoHook()
    ]
