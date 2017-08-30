
# Local modules
from common import debug, text_utils
from common.telegram import telegram_utils
from common.action import action_class

from bible import bible_utils

from tms import tms_utils


CMD_TMS = "/tms"
CMD_TMS_PROMPT = "Give me a Verse reference, or Pack and Verse number\n(P.S. you can even try giving me a topic)"
CMD_TMS_BADQUERY = "I can't find anything related to this, try another one?"

STATE_WAIT_TMS = "Waiting for TMS query"

class TMSAction(action_class.Action):
    def identifier(self):
        return '/tms'

    def resolve(self, user, msg):
        query = telegram_utils.strip_command(msg, self.identifier())

        if text_utils.is_valid(query): 
            debug.log("Resolving TMS Query")
            verse = None

            verse_reference = bible_utils.get_reference(query)
            if text_utils.is_valid(verse_reference):
                verse = tms_utils.query_verse_by_reference(verse_reference)
            
            if verse is None:
                verse = tms_utils.query_verse_by_pack_pos(query)

            if verse is None:
                verse = tms_utils.query_verse_by_topic(query)

            if verse is not None:
                passage = bible_utils.get_passage_raw(verse.reference, user.get_version())
                verse_msg = tms_utils.format_verse(verse, passage)

                telegram_utils.send_msg(verse_msg, user.get_uid())
                user.set_state(None)
            else:
                telegram_utils.send_msg(CMD_TMS_BADQUERY, user.get_uid())
        else:
            telegram_utils.send_msg_keyboard(CMD_TMS_PROMPT, user.get_uid())
            user.set_state(STATE_WAIT_TMS)

        return True


ACTION = TMSAction()
def get_action():
    return ACTION
