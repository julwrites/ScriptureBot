
# Local modules
from common import debug, text_utils
from common.telegram import telegram_utils
from common.action import action_classes


SUPPORTED_VERSIONS = ["NIV", "ESV", "KJV", "NASB", "NLT", "AMP"]

PROMPT = "Please select a version of your choosing\n\
(if unsure, always go with the one you are comfortable with!)"
BADQUERY = "I don't have this version!"

STATE_VERSION_PROMPT = "I\'ve changed your version to {}!"

class BibleUserVersionAction(action_classes.Action):
    def identifier(self):
        return '/version'

    def name(self):
        return 'Bible Version'

    def resolve(self, userObj, msg):
        query = telegram_utils.strip_command(msg, self.identifier())

        if text_utils.is_valid(query):

            for ver in SUPPORTED_VERSIONS:

                if text_utils.text_compare(query, ver):
                    userObj.set_version(ver)
                    userObj.set_state(None)

                    telegram_utils.send_close_keyboard(\
                    STATE_VERSION_PROMPT.format(ver), userObj.get_uid())
                    break
            else:
                telegram_utils.send_msg(BADQUERY, userObj.get_uid())

        else:
            telegram_utils.send_msg_keyboard(\
            PROMPT, userObj.get_uid(), SUPPORTED_VERSIONS)

            userObj.set_state(self.identifier())

        return True

def get():
    return [
        BibleUserVersionAction()
    ]