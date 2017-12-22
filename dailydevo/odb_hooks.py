# coding=utf-8

# Local modules
from common import debug, text_utils
from common.action import hook_classes
from common.telegram import telegram_utils

from dailydevo import odb_utils
from user import user_actions

PROMPT = "Here are today's articles from Our Daily Breadk!\nTap on any one to get the article!"


class ODBDevoHook(hook_classes.Hook):
    def identifier(self):
        return "/odb"

    def name(self):
        return "Our Daily Bread"

    def description(self):
        return "Articles from Our Daily Bread"

    def resolve(self, userObj):
        refs = odb_utils.get_odb()

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
    return [ODBDevoHook()]
