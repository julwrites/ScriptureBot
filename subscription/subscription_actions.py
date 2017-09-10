
# coding=utf-8

# Local modules
from common import debug, text_utils
from common.telegram import telegram_utils
from common.action import action_classes

from dailydevo import dailydevo_modules

PROMPT = "Choose any of these subscriptions to subscribe or unsubscribe!"
BADQUERY = "I don't have this subscription!"
CONFIRM_SUBSCRIBE = "I\'ve set up your subscription to {}!"
CONFIRM_UNSUBSCRIBE = "I\'ve unsubscribed you from {}!"


class DevoSubscriptionAction(action_classes.Action):
    def identifier(self):
        return "/dailydevosub"

    def name(self):
        return "Daily-Devo Subscriptions"

    def description(self):
        return "Subscribe to get reading material daily"

    def is_command(self):
        return True

    def resolve(self, userObj, msg):
        query = telegram_utils.strip_command(msg, self.identifier())
        subs = dailydevo_modules.get_hooks()

        if text_utils.is_valid(query):

            for sub in subs:

                if text_utils.fuzzy_compare(query, sub.name()):

                    if userObj.has_subscription(sub.identifier()):
                        userObj.remove_subscription(sub.identifier())

                        telegram_utils.send_close_keyboard(\
                        CONFIRM_UNSUBSCRIBE.format(sub.name()), userObj.get_uid())

                    else:
                        userObj.add_subscription(sub.identifier())

                        telegram_utils.send_close_keyboard(\
                        CONFIRM_SUBSCRIBE.format(sub.name()), userObj.get_uid())

                    userObj.set_state(None)
                    break
            else:
                telegram_utils.send_msg(BADQUERY, userObj.get_uid())

        else:
            subList = [sub.name() for sub in subs]

            for i in range(len(subList)):

                if userObj.has_subscription(subs[i].identifier()):
                    subList[i] = subList[i] + ' ' + telegram_utils.tick()

            telegram_utils.send_msg_keyboard(\
            PROMPT, userObj.get_uid(), subList, 1)

            userObj.set_state(self.identifier())

        return True

def get():
    return [
        DevoSubscriptionAction()
    ]