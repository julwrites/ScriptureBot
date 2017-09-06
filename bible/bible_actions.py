
# Local modules
from common import debug, text_utils
from common.action import action_classes
from common.telegram import telegram_utils

from bible import bible_utils

CMD_PASSAGE = "/passage"
CMD_PASSAGE_PROMPT = "Give me a Bible reference"
CMD_PASSAGE_BADQUERY = "Sorry, I can't find this reference"


class BiblePassageAction(action_classes.Action):
    def identifier(self):
        return '/passage'

    def resolve(self, user_obj, msg):
        query = telegram_utils.strip_command(msg, self.identifier())

        if text_utils.is_valid(query):
            passage = bible_utils.get_passage(query, user_obj.get_version())

            if passage is not None:
                telegram_utils.send_msg(passage, user_obj.get_uid())
                user_obj.set_state(None)
            else:
                telegram_utils.send_msg(CMD_PASSAGE_BADQUERY, user_obj.get_uid())
        else:
            telegram_utils.send_msg(CMD_PASSAGE_PROMPT, user_obj.get_uid())

            user_obj.set_state(self.identifier())

        return True

def get():
    return [
        BiblePassageAction()
    ]