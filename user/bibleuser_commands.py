
# Local modules
from common import debug
from common import database
from common import telegram
from common import telegram_utils

from common.constants import *

from user.bibleuser_utils import *


SUPPORTED_VERSIONS = ['NIV', 'ESV', 'KJV', 'NASB', 'NLT', 'AMP']

CMD_VERSION = '/version'
CMD_VERSION_PROMPT = 'Please select a version of your choosing \n (if unsure, always go with the one you are comfortable with!)'

STATE_WAIT_VERSION = 'Waiting for version'
STATE_VERSION_PROMPT = 'I\'ve changed your version to {}!'

def cmds(user, cmd, msg):
    if user is None:
        return False

    debug.log('Running user settings commands')

    return (    \
    cmd_version(user, cmd, msg) \
    )

def states(user, msg):
    if user is None:
        return False

    debug.log('Running user settings states')

    return ( \
    state_version(user, msg)       \
    )

def cmd_version(user, cmd, msg):
    if cmd == CMD_VERSION:
        debug.log('Command: ' + cmd)

        telegram.send_msg_keyboard(CMD_VERSION_PROMPT, user.get_uid(), SUPPORTED_VERSIONS)
        user.set_state(STATE_WAIT_VERSION)

        return True
    return False

def state_version(user, msg):
    if user.get_state() == STATE_WAIT_VERSION:
        debug.log('State: ' + STATE_WAIT_VERSION)

        version_found = False
        version = msg.get('text')
        for ver in SUPPORTED_VERSIONS:
            if text_utils.fuzzy_compare(version, ver):
                version_found = True
                user.set_version(ver)
                telegram.send_close_keyboard(user.get_uid())
                telegram.send_msg(STATE_VERSION_PROMPT.format(ver), user.get_uid())
                user.set_state(None)

        if not version_found:
            telegram.send_msg('That is not a version!', user.get_uid())

        return True
    return False