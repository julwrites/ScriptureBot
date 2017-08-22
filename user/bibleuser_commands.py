
# Local modules
from common import debug
from common import database
from common import telegram
from common import telegram_utils

from user import bibleuser_utils

CMD_VERSION = '/version'
CMD_VERSION_PROMPT = 'Please select a version of your choosing \n (if unsure, always go with the one you are comfortable with!)'

def cmds(user, cmd, msg):
    if user is None:
        return False

    debug.log('Running user settings commands')

    return (    \
    cmd_version(user, cmd, msg) \
    )

def cmd_version(user, cmd, msg):
    if cmd == CMD_VERSION:
        debug.log('Command: ' + cmd)

        telegram.send_msg_keyboard(CMD_VERSION_PROMPT, user.get_uid(), ['NIV', 'ESV', 'KJV', 'RSV', 'NASB'])

        return True

    return False
