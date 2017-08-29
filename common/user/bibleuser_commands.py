
# Local modules
from common import debug
from common import database

from common.telegram import telegram_utils

from common.constants import *

from common.user.bibleuser_utils import *

from common.classes import action_class


SUPPORTED_VERSIONS = ["NIV", "ESV", "KJV", "NASB", "NLT", "AMP"]

CMD_VERSION = "/version"
CMD_VERSION_PROMPT = "Please select a version of your choosing\n\
(if unsure, always go with the one you are comfortable with!)"
CMD_VERSION_BADQUERY = "I don't have this version!"

STATE_WAIT_VERSION = "Waiting for version"
STATE_VERSION_PROMPT = "I\"ve changed your version to {}!"

class BibleUserAction(action_class.Action):
    def identifier(self):
        return '/version'

    def resolve(self, user, msg):
        debug.log('Action being executed: ' + self.identifier())
        text = msg.get('text').strip()
        query = text.replace(self.identifier(), '')

        if text_utils.is_valid(query):

            for ver in SUPPORTED_VERSIONS:

                if text_utils.text_compare(query, ver):
                    user.set_version(ver)
                    user.set_state(None)

                    telegram_utils.send_close_keyboard(\
                    STATE_VERSION_PROMPT.format(ver), user.get_uid())
                    break
            else:
                telegram_utils.send_msg(CMD_VERSION_BADQUERY, user.get_uid())

        else:
            telegram_utils.send_msg_keyboard(\
            CMD_VERSION_PROMPT, user.get_uid(), SUPPORTED_VERSIONS)

            user.set_state(self.identifier())

        return True

ACTION = BibleUserAction()
def get_action():
    return ACTION


def cmds(user, cmd, msg):
    if user is None:
        return False

    debug.log("Running user settings commands")

    return (    \
    cmd_version(user, cmd, msg) \
    )

def states(user, msg):
    if user is None:
        return False

    debug.log("Running user settings states")

    return ( \
    state_version(user, msg)       \
    )

def resolve_version(user, query):
    if user is not None:
        if text_utils.is_valid(query):
            for ver in SUPPORTED_VERSIONS:
                if text_utils.text_compare(query, ver):
                    user.set_version(ver)
                    user.set_state(None)
                    telegram_utils.send_close_keyboard(STATE_VERSION_PROMPT.format(ver), user.get_uid())
                    break
            else:
                telegram_utils.send_msg(CMD_VERSION_BADQUERY, user.get_uid())
        else:
            telegram_utils.send_msg_keyboard(CMD_VERSION_PROMPT, user.get_uid(), SUPPORTED_VERSIONS)
            user.set_state(STATE_WAIT_VERSION)
        return True
    return False

def cmd_version(user, cmd, msg):
    if cmd == CMD_VERSION:
        debug.log_cmd(cmd)
        query = msg.get("text").strip()
        query = query.replace(cmd, '')

        return resolve_version(user, query)

def state_version(user, msg):
    if user.get_state() == STATE_WAIT_VERSION:
        debug.log_state(STATE_WAIT_VERSION)
        query = msg.get("text").strip()

        return resolve_version(user, query)
    return False