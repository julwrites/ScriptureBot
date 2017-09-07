
# Local modules
from common import debug, text_utils
from common.telegram import telegram_utils
from common.action import action_classes

from devo import cac_hooks, mcheyne_hooks

PROMPT = "Please select a devotional of your choosing\n\
(if unsure, always go with the one you are comfortable with!)"
BADQUERY = "I don't have this subscription!"
CONFIRM = "I\'ve set up your subscription to {}!"

class DevoSubscriptionAction(action_classes.Action):
    def identifier(self):
        return '/version'

    def name(self):
        return 'Bible Version'

    def resolve(self, userObj, msg):
        query = telegram_utils.strip_command(msg, self.identifier())
        devos = \
        cac_hooks.get() + \
        mcheyne_hooks.get()

        if text_utils.is_valid(query):

            for devo in devos:

                if text_utils.text_compare(query, devo.name()):
                    userObj.set_subscription(devo)
                    userObj.set_state(None)

                    telegram_utils.send_close_keyboard(\
                    CONFIRM.format(devo.name()), userObj.get_uid())
                    break
            else:
                telegram_utils.send_msg(BADQUERY, userObj.get_uid())

        else:
            telegram_utils.send_msg_keyboard(\
            PROMPT, userObj.get_uid(), [devo.name() for devo in devos])

            userObj.set_state(self.identifier())

        return True

def get():
    return [
        DevoSubscriptionAction()
    ]