
# Local modules
from common import debug, text_utils, telegram
from common.classes import action

import bible

CMD_PASSAGE = "/passage"
CMD_PASSAGE_PROMPT = "Give me a Bible reference"
CMD_PASSAGE_BADQUERY = "Sorry, I can't find this reference"


class BiblePassageAction(action.Action):
    def identifier(self):
        return '/passage'

    def resolve(self, user, msg):
        query = telegram.utils.strip_command(msg, self.identifier())

        if text_utils.is_valid(query):
            passage = bible.utils.get_passage(query, user.get_version())

            if passage is not None:
                telegram.utils.send_msg(passage, user.get_uid())
                user.set_state(None)
            else:
                telegram.utils.send_msg(CMD_PASSAGE_BADQUERY, user.get_uid())
        else:
            telegram.utils.send_msg(CMD_PASSAGE_PROMPT, user.get_uid())

            user.set_state(self.identifier())

        return True

def get():
    return [
        BiblePassageAction()
    ]