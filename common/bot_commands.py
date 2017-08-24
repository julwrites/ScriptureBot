
# Local modules
from common import debug
from common import database
from common import telegram
from common import telegram_utils
from common import admin_commands
from bgw import bgw_commands

from user.bibleuser_utils import *

CMD_STORE = '/store'
CMD_RETRIEVE = '/retrieve'

STATE_WAIT_STORE = 'Waiting For Store'

def cmds(user, cmd, msg):
    if user is None:
        return False

    debug.log('Running bot commands')

    try:
        result = ( \
        cmd_store(user, cmd, msg)           \
        or cmd_retrieve(user, cmd, msg)     \
        )
    except:
        debug.log('Exception in Bot Commands')
        return False
    return result

def states(user, msg):
    if user is None:
        return False

    debug.log('Running bot states')

    try:
        result = ( \
        state_store(user, msg)       \
        )
    except:
        debug.log('Exception in Bot States')
        return False
    return result

def cmd_store(user, cmd, msg):
    if user is not None:
        if cmd == CMD_STORE:
            debug.log_cmd(cmd)
        
            if database.has_data(user.get_uid()):
                debug.log('Prompting data store for User: ' + user.get_name_string())

                telegram.send_msg('Please send the data to be stored', user.get_uid())
                user.set_state(STATE_WAIT_STORE)

            return True
    return False

def state_store(user, msg):
    if user is not None and user.get_state() == STATE_WAIT_STORE:
        debug.log_state(STATE_WAIT_STORE)

        payload = telegram_utils.parse_payload(msg)
        database.set_data(user.get_uid(), payload)

        user.set_state(None)

        debug.log('Stored Data to User: ' + user.get_name_string())

        return True
    return False

def cmd_retrieve(user, cmd, msg):
    if user is not None:
        if cmd == CMD_RETRIEVE:
            debug.log_cmd(cmd)

            if database.has_data(user.get_uid()):
                debug.log(user.get('username'))
        
                telegram.send_msg('Retrieving Data of User: ' + user.get_name_string(), user.get_uid())
                telegram.send_msg(database.get_data(user.get_uid()), user.get_uid())

            return True
    return False
