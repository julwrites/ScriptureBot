
# Python modules

# Local modules
from common import debug
from common import database
from common import telegram
from common import telegram_utils

from tms import tms_utils
from bgw import bgw_utils

CMD_VERSE = '/verse'
CMD_VERSE_PROMPT = 'Please enter /verse followed by a bible reference or topic\n(I hope I know something about it)'

def cmds(user, cmd, msg):
    if user is None:
        return False

    debug.log('Running Bible commands')

    return (    \
    cmd_verse(user, cmd, msg) \
    )

def cmd_verse(user, cmd, msg):
    if cmd == CMD_VERSE:
        debug.log_cmd(cmd)

        query = msg.get('text')
        query = query.replace(cmd, '')

        verse = tms_utils.query_verse_by_topic(query)
        if verse is not None:
            passage = bgw_utils.get_passage_raw(verse.get_reference(), user.get_version())
            verse_msg = tms_utils.format_verse(verse, passage)

            debug.log("Sending TMS verse: " + verse_msg)

            telegram.send_msg(verse_msg, user.get_uid())
        else:
            telegram.send_msg(CMD_PASSAGE_PROMPT, user.get_uid())

        return True

    return False
