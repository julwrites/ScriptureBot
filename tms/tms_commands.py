
# Local modules
from common import debug
from common import database
from common import telegram
from common import telegram_utils
from bgw import bgw_utils
from tms import tms_utils

CMD_TMS = '/tms'
CMD_TMS_PROMPT = 'Give me a Verse reference, or Pack and Verse number\n(P.S. you can even try giving me a topic)'

STATE_WAIT_TMS = 'Waiting for TMS query'

def cmds(user, cmd, msg):
    if user is None:
        return False
    
    debug.log('Running TMS commands')

    try:
        result = (\
        cmd_tms(user, cmd, msg) \
        )
    except:
        return False
    return result

def states(user, msg):
    if user is None:
        return False

    debug.log('Running TMS states')
    
    try:
        result = ( \
        state_tms(user, msg)       \
        )
    except:
        return False
    return result

def resolve_tms_query(user, query):
    if user is not None and query is not None:
        verse = None
        
        verse_reference = bgw_utils.get_reference(query)
        if verse_reference is not None:
            debug.log('Attempting to get by reference ' + verse_reference)
            verse = tms_utils.query_verse_by_reference(verse_reference)
        
        if verse is None:
            debug.log('Attempting to get by position ' + query)
            verse = tms_utils.query_verse_by_pack_pos(query)

        if verse is None:
            debug.log('Attempting to get by topic ' + query)
            verse = tms_utils.query_verse_by_topic(query)

        if verse is not None:
            passage = bgw_utils.get_passage_raw(verse.reference, user.get_version())
            verse_msg = tms_utils.format_verse(verse, passage)
            
            debug.log("Sending TMS verse: " + verse_msg)

            telegram.send_msg(verse_msg, user.get_uid())
        else:
            telegram.send_msg_keyboard(CMD_TMS_PROMPT, user.get_uid())
            user.set_state(STATE_WAIT_TMS)

        return True

    return False

def cmd_tms(user, cmd, msg):
    if cmd == CMD_TMS:
        debug.log_cmd(cmd)

        query = msg.get('text')
        query = query.replace(cmd, '').strip()

        return resolve_tms_query(user, query)
    return False

def state_tms(user, msg):
    if user is not None and user.get_state() is STATE_WAIT_TMS:
        query = msg.get('text')

        if resolve_tms_query(user, query):
            user.set_state(None)
            return True
    return False