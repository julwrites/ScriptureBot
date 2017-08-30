
# Local modules
from common import debug, text_utils
from common.telegram import telegram_utils
from common.action import action_class

from bible import bible_utils

CMD_PASSAGE = "/passage"
CMD_PASSAGE_PROMPT = "Give me a Bible reference"
CMD_PASSAGE_BADQUERY = "Sorry, I can't find this reference"


class BibleAction(action_class.Action):
    def identifier(self):
        return '/passage'

    def resolve(self, user, msg):
        query = telegram_utils.strip_command(msg, self.identifier())

        if text_utils.is_valid(query):
            passage = bible_utils.get_passage(query, user.get_version())

            if passage is not None:
                telegram_utils.send_msg(passage, user.get_uid())
                user.set_state(None)
            else:
                telegram_utils.send_msg(CMD_PASSAGE_BADQUERY, user.get_uid())
        else:
            telegram_utils.send_msg(CMD_PASSAGE_PROMPT, user.get_uid())

            user.set_state(self.identifier())

        return True

ACTION = BibleAction()
def get_action():
    return ACTION