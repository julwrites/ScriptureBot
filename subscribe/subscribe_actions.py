# coding=utf-8

# Local modules
from common import debug, text_utils
from common.telegram import telegram_utils
from common.action import action_classes

from dailydevo import dailydevo_modules

PROMPT = "Choose any of these subscriptions to subscribe or unsubscribe!"
BADQUERY = "I don't have this subscription!"
CONFIRM_SUBSCRIBE = "I've set up your subscription to {}!"
CONFIRM_UNSUBSCRIBE = "I've unsubscribed you from {}!"


class SubscribeAction(action_classes.Action):
    def identifier(self):
        return "/subscribe"

    def name(self):
        return "Subscriptions"

    def description(self):
        return "Subscribe to / Unsubscribe from daily reading material"

    def is_command(self):
        return True

    def resolve(self, userObj, msg):
        query = telegram_utils.strip_command(msg, self.identifier())
        subs = dailydevo_modules.get_hooks()

        debug.log("Querying " + query)

        if text_utils.is_valid(query):

            for sub in subs:

                if text_utils.fuzzy_compare(query, sub.name()):

                    if userObj.has_subscription(sub.identifier()):
                        userObj.remove_subscription(sub.identifier())

                        telegram_utils.send_reply(
                            user=userObj.get_uid(),
                            text=CONFIRM_UNSUBSCRIBE.format(sub.name()),
                            reply=telegram_utils.make_close_keyboard())

                    else:
                        userObj.add_subscription(sub.identifier())

                        telegram_utils.send_reply(
                            user=userObj.get_uid(),
                            text=CONFIRM_SUBSCRIBE.format(sub.name()),
                            reply=telegram_utils.make_close_keyboard())

                    userObj.set_state(None)
                    break
            else:
                telegram_utils.send_msg(user=userObj.get_uid(), text=BADQUERY)

        else:
            subList = [sub.name() for sub in subs]

            for i in range(len(subList)):

                if userObj.has_subscription(subs[i].identifier()):
                    subList[i] = subList[i] + " " + telegram_utils.tick()

            options = [telegram_utils.make_reply_button(text=sub) for sub in subList]

            telegram_utils.send_reply(
                user=userObj.get_uid(),
                text=PROMPT,
                reply=telegram_utils.make_reply_keyboard(buttons=options, width=1))

            userObj.set_state(self.identifier())

        return True

class ScheduleAction(action_classes.Action):
    def identifier(self):
        return "/schedule"

    def name(self):
        return "Schedule"

    def description(self):
        return "Set subscription delivery time"

    def is_command(self):
        return False

    def resolve(self, userObj, msg):

        return True

def get():
    return [
        SubscribeAction(),
        ScheduleAction()
    ]