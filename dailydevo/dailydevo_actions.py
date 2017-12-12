# coding=utf-8

# Python modules
import random

# Local modules
from dailydevo import dailydevo_modules

from common import text_utils, debug
from common.action import action_classes
from common.telegram import telegram_utils

from user import user_actions

PROMPT = "Choose a Daily-Devo to read!"
BADQUERY = "I don't have this devotional!"
CONFIRM = [
    "Give me a moment to get it {}~!",
    "Hold on, {}",
    "I'll be right back, {}",
    "Got it! Wait for a bit {}",
    "I need to step away for a moment to get it, {}",
    "Let me get back to you with that, {}~",
]


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
                    choose = random.randint(0, len(CONFIRM) - 1)
                    confirmString = CONFIRM[choose].format(
                        userObj.get_name_string())

                    telegram_utils.send_reply(
                        user=userObj.get_uid(),
                        text=confirmString,
                        reply=telegram_utils.make_close_keyboard())

                    userObj.set_state(None)

                    hook.resolve(userObj)

                    break
            else:
                telegram_utils.send_msg(user=userObj.get_uid(), text=BADQUERY)

        else:
            options = [
                telegram_utils.make_reply_button(text=hook.name())
                for hook in hooks
            ]

            telegram_utils.send_reply(
                user=userObj.get_uid(),
                text=PROMPT,
                reply=telegram_utils.make_reply_keyboard(
                    buttons=options, width=1))

            userObj.set_state(self.identifier())

        return True


def get():
    return [
        DailyDevoAction()
    ] + \
    dailydevo_modules.get_actions()
