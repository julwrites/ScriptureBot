
# Local modules
from common import debug
from common import database
from common.telegram import telegram_utils
from common import admin_commands
from bible import bible_commands
import hooks

from common.bible_user import *

CMD_STORE = '/store'
CMD_RETRIEVE = '/retrieve'

BOT_STATE_WAIT_STORE = 'Waiting For Store'

def cmds(user, cmd, msg):
    debug.log('Running bot commands')

    return ( \
    cmd_store(user, cmd, msg)           \
    or cmd_retrieve(user, cmd, msg)     \
    )

def states(msg):
    # Read the user to echo back
    uid = get_uid(msg.get('from').get('id'))
    user = get_user(uid)

    return ( \
    state_store(user, msg)       \
    )

def cmd_store(user, cmd, msg):
    if user is not None:
        if cmd == CMD_STORE:
            debug.log_cmd(cmd)
       
            if database.has_data(user.get_uid()):
                debug.log('Prompting data store for User: ' + user.get_name_string())

                telegram_utils.send_msg('Please send the data to be stored', user.get_uid())
                user.set_state(BOT_STATE_WAIT_STORE)

                return True
    return False

def state_store(user, msg):
    if user is not None and user.get_state() == BOT_STATE_WAIT_STORE:
        debug.log_state(BOT_STATE_WAIT_STORE

        payload = telegram_utils.parse_payload(msg)
        database.set_data(user.get_uid(), payload)

        debug.log('Stored Data to User: ' + user.get_name_string())

        return True
    return False

def cmd_retrieve(user, cmd, msg):
    if user is not None:
        if cmd == CMD_RETRIEVE:
            debug.log_cmd(cmd)

            if database.has_data(user.get_uid()):
                debug.log(user.get('username'))
        
                telegram_utils.send_msg('Retrieving Data of User: ' + user.get_name_string(), user.get_uid())
                telegram_utils.send_msg(database.get_data(user.get_uid()), user.get_uid())

                return True
    return False
