# coding=utf-8

# Local modules
from common.utils import debug_utils
from common.telegram import telegram_utils
from common.action import hook_classes

from dailydevo import mcheyne_utils
from user import user_actions

PROMPT = "Here are today's M'Cheyne passages!\n{}\nTap on any one to get the passage!"


class McheyneDailyHook(hook_classes.Hook):
    def identifier(self):
        return "/mcheyne"

    def name(self):
        return "M'cheyne Bible Reading Plan"

    def description(self):
        return "M'cheyne Bible Reading Plan (1 Year)"

    def resolve(self, userObj):
        debug_utils.log("Resolving MCheyne hook")

        refs = mcheyne_utils.get_mcheyne()

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
    return [McheyneDailyHook()]
