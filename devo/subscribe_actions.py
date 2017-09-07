
# Local modules
from common import debug, text_utils
from common.telegram import telegram_utils
from common.action import action_classes

from devo import devo_modules

PROMPT = "Please select a devotional of your choosing\n\
(if unsure, always go with the one you are comfortable with!)"
BADQUERY = "I don't have this subscription!"
CONFIRM_SUBSCRIBE = "I\'ve set up your subscription to {}!"
CONFIRM_UNSUBSCRIBE = "I\'ve unsubscribed you from {}!"


class DevoSubscriptionAction(action_classes.Action):
    def identifier(self):
        return '/devo'

    def name(self):
        return 'Devotionals'

    def resolve(self, userObj, msg):
        query = telegram_utils.strip_command(msg, self.identifier())
        devos = devo_modules.get_hooks()

        if text_utils.is_valid(query):

            for devo in devos:

                if text_utils.fuzzy_compare(query, devo.name()):

                    if userObj.has_subscription(devo.identifier()):
                        userObj.remove_subscription(devo.identifier())

                        telegram_utils.send_close_keyboard(\
                        CONFIRM_UNSUBSCRIBE.format(devo.name()), userObj.get_uid())

                    else:
                        userObj.add_subscription(devo.identifier())

                        telegram_utils.send_close_keyboard(\
                        CONFIRM_SUBSCRIBE.format(devo.name()), userObj.get_uid())

                    userObj.set_state(None)
                    break
            else:
                telegram_utils.send_msg(BADQUERY, userObj.get_uid())

        else:
            devoList = [devo.name() for devo in devos]

            for i in range(len(devoList)):

                if userObj.has_subscription(devos[i].identifier()):
                    devoList[i] = devoList[i] + ' ' + telegram_utils.tick()

            telegram_utils.send_msg_keyboard(\
            PROMPT, userObj.get_uid(), devoList, 1)

            userObj.set_state(self.identifier())

        return True

def get():
    return [
        DevoSubscriptionAction()
    ]