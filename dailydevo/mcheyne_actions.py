# coding=utf8

# Local modules
from common import debug
from common.action import action_classes
from common.telegram import telegram_utils

from dailydevo import mcheyne_utils, mcheyne_hooks
from bible import bible_utils
from user import user_actions


class McheyneDailyAction(action_classes.Action):
    def identifier(self):
        return "/mcheyne"

    def name(self):
        return "M'cheyne Bible Reading Plan"

    def description(self):
        return "M'cheyne Bible Reading Plan (1 Year)"

    def resolve(self, userObj, msg):
        query = telegram_utils.strip_command(msg, self.identifier())

        done = user_actions.UserDoneAction()
        if done.try_execute(userObj, msg):
            return True

        passage = bible_utils.get_passage(query, userObj.get_version())
        if passage is not None:
            debug.log("Sending passage " + passage)
            telegram_utils.send_msg(user=userObj.get_uid(), text=passage)

        mcheyne_hooks.McheyneDailyHook().resolve(userObj)

        return True


def get():
    return [McheyneDailyAction()]
