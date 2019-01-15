# coding=utf-8

# Local modules
from common import debug, text_utils
from common.telegram import telegram_utils
from common.action import action_classes

from user import user_actions

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

        done = user_actions.UserDoneAction()
        if done.try_execute(userObj, msg):
            return True

        if text_utils.is_valid(query):
            fuzz = [
                i for i in range(len(subs))
                if text_utils.fuzzy_compare(query, subs[i].name())
            ]

            if len(fuzz) > 0:
                max_fuzz = fuzz[0]
                for i in fuzz[1:]:
                    if text_utils.overlap_compare(
                            query, subs[i]) > text_utils.overlap_comare(
                                query, subs[max_fuzz]):
                        max_fuzz = i
                sub = subs[max_fuzz]

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
            else:
                telegram_utils.send_msg(user=userObj.get_uid(), text=BADQUERY)

        else:
            subList = [sub.name() for sub in subs]

            for i in range(len(subList)):
                if userObj.has_subscription(subs[i].identifier()):
                    subList[i] = subList[i] + " " + telegram_utils.tick()

            options = [
                telegram_utils.make_reply_button(text=sub) for sub in subList
            ]

            options.append(telegram_utils.make_reply_button(text=done.name()))

            telegram_utils.send_reply(
                user=userObj.get_uid(),
                text=PROMPT,
                reply=telegram_utils.make_reply_keyboard(
                    buttons=options, width=1))

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
    return [SubscribeAction(), ScheduleAction()]
