# coding=utf-8

# Local modules
from common import debug, text_utils
from common.telegram import telegram_utils
from common.action import action_classes

from bible import bible_utils

PROMPT = "Please select a version of your choosing\n\
(if unsure, always go with the one you are comfortable with!)"

BADQUERY = "I don't have this version!"

STATE_VERSION_PROMPT = "I\'ve changed your version to {}!"


class BibleUserVersionAction(action_classes.Action):
    def identifier(self):
        return "/version"

    def name(self):
        return "Bible Version"

    def description(self):
        return "Choose your preferred Bible version"

    def is_command(self):
        return True

    def resolve(self, userObj, msg):
        query = telegram_utils.strip_command(msg, self.identifier())

        if text_utils.is_valid(query):

            for ver in bible_utils.get_versions():

                if text_utils.text_compare(query, ver):
                    userObj.set_version(ver)
                    userObj.set_state(None)

                    telegram_utils.send_reply(
                        user=userObj.get_uid(),
                        text=STATE_VERSION_PROMPT.format(ver),
                        reply=telegram_utils.make_close_keyboard())
                    break
            else:
                telegram_utils.send_msg(user=userObj.get_uid(), text=BADQUERY)

        else:
            options = [
                telegram_utils.make_reply_button(text=version)
                for version in bible_utils.get_versions()
            ]

            telegram_utils.send_reply(
                user=userObj.get_uid(),
                text=PROMPT,
                reply=telegram_utils.make_reply_keyboard(buttons=options))

            userObj.set_state(self.identifier())

        return True


def get():
    return [BibleUserVersionAction()]
