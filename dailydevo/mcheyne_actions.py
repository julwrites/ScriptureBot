
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
        query = telegram_utils.strip_command(msg, self.identifier())

        doneAction = user_actions.UserDoneAction()
        if doneAction.try_execute(userObj, msg):
            return True

        passage = bible_utils.get_passage(query, userObj.get_version())
        if passage is not None:
            debug.log("Sending passage " + passage)
            telegram_utils.send_msg(passage, userObj.get_uid())

        refs = mcheyne_utils.get_mcheyne()

        if refs is not None:
            refs.append(doneAction.name())
            options=[telegram.make_button(text=ref) for ref in refs]

            telegram_utils.send_msg_keyboard("", userObj.get_uid(), options, 1)
            userObj.set_state(self.identifier())

        return True

def get():
    return [
        McheyneDailyAction()
    ]