
# Local modules
from common import debug
from common.telegram import telegram_utils
from common import text_utils

from bible import bible_utils

CMD_PASSAGE = "/passage"
CMD_PASSAGE_PROMPT = "Give me a Bible reference"
CMD_PASSAGE_BADQUERY = "Sorry, I can't find this reference"
STATE_WAIT_PASSAGE = "Waiting for Bible reference"

def cmds(user, cmd, msg):
    if user is None:
        return False

    debug.log("Running Bible Query commands")

    try:
        result = (    \
        cmd_passage(user, cmd, msg) \
        )
    except:
        debug.log("Exception in Bible Query Commands")
        return False
    return result

def states(user, msg):
    if user is None:
        return False

    debug.log("Running Bible Query states")
    
    try:
        result = ( \
        state_passage(user, msg)       \
        )
    except:
        debug.log("Exception in Bible Query States")
        return False
    return result


def resolve_passage_query(user, query):
    if user is not None:
        if text_utils.is_valid(query):
            passage = bible_utils.get_passage(query, user.get_version())

            if passage is not None:
                telegram_utils.send_msg(passage, user.get_uid())
                user.set_state(None)
            else:
                telegram_utils.send_msg(CMD_PASSAGE_BADQUERY, user.get_uid())
        else:
            telegram_utils.send_msg(CMD_PASSAGE_PROMPT, user.get_uid())
            user.set_state(STATE_WAIT_PASSAGE)

        return True
    return False

def cmd_passage(user, cmd, msg):
    if user is not None:
        if cmd == CMD_PASSAGE:
            debug.log_cmd(cmd)

            query = msg.get("text").strip()
            query = query.replace(cmd, "")

            return resolve_passage_query(user, query)
    return False

def state_passage(user, msg):
    if user is not None and user.get_state() == STATE_WAIT_PASSAGE:
        debug.log_state(STATE_WAIT_PASSAGE)
        query = msg.get("text").strip()

        return resolve_passage_query(user, query)
    return False

