
# Local modules
from common import debug
from common import database
from common import telegram
from common import telegram_utils
from bgw import bgw_utils
from tms import tms_utils

CMD_TMS = '/tms'

def cmds(user, cmd, msg):
    if user is None:
        return False
    
    debug.log('Running TMS commands')

    return (    \
    cmd_tms(user, cmd, msg) \
    )

def cmd_tms(user, cmd, msg):
    if cmd == CMD_TMS:
        debug.log('Command: ' + cmd)

        query = msg.get('text')
        query = query.replace(cmd, '').strip()
        query = query.split(' ')

        debug.log('Attempting to get ' + '|'.join(query))

        verse = None
        if len(query) == 2:
            verse = tms_utils.get_verse_by_pack(query[0], int(query[1]))

        if verse is not None:
            verse_text = bgw_utils.get_passage(verse.reference, user.version)
            verse_msg = tms_utils.format_verse(verse, verse_text)
            
            debug.log("Sending TMS verse: " + verse_msg)

            telegram.send_msg(verse_msg, user.get_uid())

            return True

    return False

