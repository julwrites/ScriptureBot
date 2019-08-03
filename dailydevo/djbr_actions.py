# coding=utf8

# Local modules
from common.utils import debug_utils
from common.action import action_classes
from common.telegram import telegram_utils

from dailydevo import djbr_utils, djbr_hooks
from bible import bible_utils
from user import user_actions


class DJBRDailyAction(action_classes.Action):
    def identifier(self):
        return "/djbr"

    def name(self):
        return "Discipleship Journal Bible Reading Plan"

    def description(self):
        return "Discipleship Journal 1-Year Bible Reading Plan"

    def resolve(self, userObj, msg):
        debug_utils.log("Handling DJBR action")

        query = telegram_utils.strip_command(msg, self.identifier())

        done = user_actions.UserDoneAction()
        if done.try_execute(userObj, msg):
            return True

        passage = bible_utils.get_passage(query, userObj.get_version())
        if passage is not None:
            debug_utils.log("Sending passage {}", [passage])
            telegram_utils.send_msg(user=userObj.get_uid(), text=passage)
        else:
            djbr_hooks.DJBRDailyHook().resolve(userObj)

        return True


def get():
    return [DJBRDailyAction()]
