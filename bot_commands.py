
# Local modules
import debug
import database
import telegram
import telegram_utils
import admin_commands
import bgw_commands
import hooks

from bible_user import *

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
            debug.log('Command: ' + cmd)
       
            if database.has_data(user.get_uid()):
                debug.log('Prompting data store for User: ' + user.get_name_string())

                telegram.send_msg('Please send the data to be stored', user.get_uid())
                user.set_state(BOT_STATE_WAIT_STORE)

                return True

    return False

def state_store(user, msg):
    if user is not None:
        if user.get_state() == BOT_STATE_WAIT_STORE:
            debug.log('Handler: ' + user.get_state())

            payload = telegram_utils.parse_payload(msg)
            database.set_data(user.get_uid(), payload)

            debug.log('Stored Data to User: ' + user.get_name_string())

            return True

    return False

def cmd_retrieve(user, cmd, msg):
    if user is not None:
        if cmd == CMD_RETRIEVE:
            debug.log('Command: ' + cmd)

            if database.has_data(user.get_uid()):
                debug.log(user.get('username'))
        
                telegram.send_msg('Retrieving Data of User: ' + user.get_name_string(), user.get_uid())
                telegram.send_msg(database.get_data(user.get_uid()), user.get_uid())

                return True

    return False
