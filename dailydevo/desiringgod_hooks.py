
# coding=utf-8

# Local modules
from common import debug, text_utils
from common.action import hook_classes
from common.telegram import telegram_utils

from dailydevo import desiringgod_utils
from user import user_actions

PROMPT = "Here are today's articles from desiringgod.org!\n{}\nTap on any one to get the article!"

class DGDevoHook(hook_classes.Hook):
    def identifier(self):
        return "/desiringgod"

    def name(self):
        return "Desiring God Articles"

    def description(self):
        return "Articles from DesiringGod.org"

    def resolve(self, userObj):
        refs = desiringgod_utils.get_desiringgod_raw()

        if refs is not None:
            options = refs
            refString = "\n".join(options)
            options.append(user_actions.UserDoneAction().name())

            telegram_utils.send_url_keyboard(PROMPT.format(refString), userObj.get_uid(), options, urls, 1)
            userObj.set_state(self.identifier())


def get():
    return [
        DGDevoHook()
    ]
