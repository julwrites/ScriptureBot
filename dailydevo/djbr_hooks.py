# coding=utf-8

# Local modules
from common import debug
from common.telegram import telegram_utils
from common.action import hook_classes

from dailydevo import djbr_utils
from user import user_actions

PROMPT = "Here are today's Bible Reading passages!\n{}\nTap on any one to get the passage!"


class DJBRDailyHook(hook_classes.Hook):
    def identifier(self):
        return "/djbr"

    def name(self):
        return "Discipleship Journal Bible Reading Plan"

    def description(self):
        return "Discipleship Journal 1-Year Bible Reading Plan"

    def resolve(self, userObj):
        debug.log("Resolving DJBR hook")

        refs = djbr_utils.get_djbr()

        if refs is not None:
            refString = "\n".join(refs)
            refs.append(user_actions.UserDoneAction().name())
            options = [
                telegram_utils.make_reply_button(text=ref) for ref in refs
            ]

            telegram_utils.send_reply(
                user=userObj.get_uid(),
                text=PROMPT.format(refString),
                reply=telegram_utils.make_reply_keyboard(
                    buttons=options, width=1))

            userObj.set_state(self.identifier())


def get():
    return [DJBRDailyHook()]
