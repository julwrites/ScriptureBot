
# Local modules
from common import debug
from common import database
from common import telegram
from common import telegram_utils
from bgw import bgw_utils
from tms import tms_utils

CMD_TMS = '/tms'
CMD_TMS_PROMPT = 'Please enter /tms followed by the Verse reference, or Pack number and Verse number'

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

        verse = None
        verse_reference = bgw_utils.get_reference(query)

        if verse_reference is not None:
            debug.log('Attempting to get ' + verse_reference)
            verse = tms_utils.get_verse_by_reference(verse_reference)
        else:
            debug.log('Attempting to get ' + query)
            pack_pos = tms_utils.find_pack_pos(query)

            if pack_pos is not None:
                verse = tms_utils.get_verse_by_pack(pack_pos[0], pack_pos[1])

        if verse is not None:
            verse_text = bgw_utils.get_passage(verse.reference, user.version)
            verse_msg = tms_utils.format_verse(verse, verse_text)
            
            debug.log("Sending TMS verse: " + verse_msg)

            telegram.send_msg(verse_msg, user.get_uid())
        else:
            telegram.send_msg(CMD_TMS_PROMPT, user.get_uid())

        return True

    return False

