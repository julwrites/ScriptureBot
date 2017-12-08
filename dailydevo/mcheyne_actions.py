
# coding=utf8

# Local modules
from common import debug
from common.action import action_classes
from common.telegram import telegram_utils

from dailydevo import mcheyne_utils
from bible import bible_utils
from user import user_actions

class McheyneDailyAction(action_classes.Action):
    def identifier(self):
        return "/mcheynedaily"

    def name(self):
        return "M'cheyne Bible Reading Plan"

    def description(self):
        return "M'cheyne Bible Reading Plan (1 Year)"

    def resolve(self, userObj, msg):
        refs = mcheyne_utils.get_mcheyne()

        if refs is not None:
            refs.append(user_actions.UserDoneAction().name())
            options=[telegram_utils.make_button(text=ref) for ref in refs]

            telegram_utils.send_msg_keyboard("", userObj.get_uid(), options, 1)

        return True

def get():
    return [
        McheyneDailyAction()
    ]