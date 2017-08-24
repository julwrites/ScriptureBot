
# Local modules
from common import debug
from common import database
from common import telegram
from common import telegram_utils

from bgw import bgw_utils

CMD_PASSAGE = '/passage'
CMD_PASSAGE_PROMPT = 'Give me a Bible reference!'
STATE_WAIT_PASSAGE = 'Waiting for Bible reference'

def cmds(user, cmd, msg):
    if user is None:
        return False

    debug.log('Running BGW commands')

    try:
        result = (    \
        cmd_passage(user, cmd, msg) \
        )
    except:
        debug.log('Exception in BGW Commands')
        return False
    return result

def states(user, msg):
    if user is None:
        return False

    debug.log('Running BGW states')
    
    try:
        result = ( \
        state_passage(user, msg)       \
        )
    except:
        debug.log('Exception in BGW States')
        return False
    return result


def resolve_passage_query(user, query):
    if user is not None and text_utils.is_valid(query):
        passage = bgw_utils.get_passage(query, user.get_version())

        if text_utils.is_valid(query):
            telegram.send_msg(passage, user.get_uid())
        else:
            telegram.send_msg(CMD_PASSAGE_PROMPT, user.get_uid())
            user.set_state(STATE_WAIT_PASSAGE)

        return True
    return False

def cmd_passage(user, cmd, msg):
    if user is not None:
        if cmd == CMD_PASSAGE:
            debug.log_cmd(cmd)

            query = msg.get('text')
            query = query.replace(cmd, '')

            return resolve_passage_query(user, query)
    return False

def state_passage(user, msg):
    if user is not None and user.get_state() is STATE_WAIT_PASSAGE:
        debug.log_state(STATE_WAIT_PASSAGE)
        query = msg.get('text')

        if resolve_passage_query(user, query):
            user.set_state(None)
            return True
    return False

