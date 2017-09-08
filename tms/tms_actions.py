
# coding=utf-8

# Local modules
from common import debug, text_utils
from common.telegram import telegram_utils
from common.action import action_classes

from bible import bible_utils

from tms import tms_utils


PROMPT = "Give me a Verse reference, or Pack and Verse number\n(P.S. you can even try giving me a topic)"
BADQUERY = "I can't find anything related to this, try another one?"

class TMSAction(action_classes.Action):
    def identifier(self):
        return '/tms'

    def name(self):
        return 'Topical Memory System'

    def resolve(self, userObj, msg):
        query = telegram_utils.strip_command(msg, self.identifier())

        if text_utils.is_valid(query): 
            verse = None

            verseReference = bible_utils.get_reference(query)
            if text_utils.is_valid(verseReference):
                verse = tms_utils.query_verse_by_reference(verseReference)
            
            if verse is None:
                verse = tms_utils.query_verse_by_pack_pos(query)

            if verse is None:
                verse = tms_utils.query_verse_by_topic(query)

            if verse is not None:
                passage = bible_utils.get_passage_raw(verse.reference, userObj.get_version())
                verseMsg = tms_utils.format_verse(verse, passage)

                telegram_utils.send_msg(verseMsg, userObj.get_uid())
                userObj.set_state(None)
            else:
                telegram_utils.send_msg(BADQUERY, userObj.get_uid())
        else:
            telegram_utils.send_msg_keyboard(PROMPT, userObj.get_uid())

            userObj.set_state(self.identifier())

        return True


def get():
    return [
        TMSAction()
    ]
