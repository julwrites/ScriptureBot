
# coding=utf-8

# Local modules
from common import debug
from common.telegram import telegram_utils
from common.action import hook_classes

from dailydevo import mcheyne_utils
from user import user_actions

PROMPT = "Here are today's M'Cheyne verses!\n{}\nTap on any one to get the passage!"

class McheyneDailyHook(hook_classes.Hook):
    def identifier(self):
        return "/mcheynedaily"

    def name(self):
        return "M'cheyne Bible Reading Plan"

    def description(self):
        return "M'cheyne Bible Reading Plan (1 Year)"

    def resolve(self, userObj):
        refs = mcheyne_utils.get_mcheyne()

        if refs is not None:
            refString = "\n".join(refs)
            refs.append(user_actions.UserDoneAction().name())
            options = [telegram_utils.make_button(text=ref) for ref in refs]

            telegram_utils.send_msg_keyboard(PROMPT.format(refString), userObj.get_uid(), options, 1)

def get():
    return [
        McheyneDailyHook()
    ]