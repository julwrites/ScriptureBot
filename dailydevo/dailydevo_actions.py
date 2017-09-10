
# coding=utf-8

# Local modules
from dailydevo import dailydevo_modules

from common import text_utils, debug
from common.action import action_classes
from common.telegram import telegram_utils

PROMPT = "Choose a Daily-Devo to read!"
BADQUERY = "I don\'t have this devotional!"
CONFIRM = "Give me a moment to get it~!"

class DailyDevoAction(action_classes.Action):
    def identifier(self):
        return "/dailydevo"

    def name(self):
        return "Daily-Devo"

    def description(self):
        return "Get reading material right now"

    def is_command(self):
        return True

    def resolve(self, userObj, msg):
        query = telegram_utils.strip_command(msg, self.identifier())
        hooks = dailydevo_modules.get_hooks()

        if text_utils.is_valid(query):

            for hook in hooks:

                if text_utils.text_compare(query, hook.name()):

                    telegram_utils.send_close_keyboard(CONFIRM, userObj.get_uid())
                    userObj.set_state(None)

                    hook.resolve(userObj)

                    break
            else:
                telegram_utils.send_msg(BADQUERY, userObj.get_uid())

        else:
            options = [hook.name() for hook in hooks]

            telegram_utils.send_msg_keyboard(\
            PROMPT, userObj.get_uid(), options, 1)

            userObj.set_state(self.identifier())

        return True

def get():
    return [
        DailyDevoAction()
    ] + \
    dailydevo_modules.get_actions()
