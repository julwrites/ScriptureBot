
# coding=utf-8
            bible_actions.get() + \

# Local modules
from common import debug, text_utils
from common.action import action_classes
from common.telegram import telegram_utils

from bible import bible_utils

PROMPT = "Give me a Bible reference"
BADQUERY = "Sorry, I can't find this reference"


class BiblePassageAction(action_classes.Action):
    def identifier(self):
        return "/passage"

    def name(self):
        return "Bible Passage"

    def description(self):
        return "Search for a passage of Scripture"

    def match(self, msg):
        return msg is not None

    def resolve(self, userObj, msg):
        query = telegram_utils.strip_command(msg, self.identifier())

        if text_utils.is_valid(query):
            passage = bible_utils.get_passage(query, userObj.get_version())

            if passage is not None:
                telegram_utils.send_msg(passage, userObj.get_uid())
                userObj.set_state(None)
            else if self.waiting(self, userObj):
                telegram_utils.send_msg(BADQUERY, userObj.get_uid())
            else:
                return False
        else:
            telegram_utils.send_msg(PROMPT, userObj.get_uid())

            userObj.set_state(self.identifier())

        return True

def get():
    return [
        BiblePassageAction()
    ]