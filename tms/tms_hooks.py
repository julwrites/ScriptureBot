# coding=utf-8

# Local modules
from common.utils import debug
from common.telegram import telegram_utils
from common.action import hook_classes
from user import user_utils

from tms import tms_utils
from bible import bible_utils


class DailyTMSHook(hook_classes.Hook):
    def identifier(self):
        return "/tms"

    def name(self):
        return "Topical Memory System"

    def description(self):
        return "The Navigators' Topical Memory System"

    def resolve(self, userObj):
        if userObj is not None:
            verse = tms_utils.get_random_verse()
            passage = bible_utils.get_passage_raw(verse.reference,
                                                  userObj.get_version())
            verseMsg = tms_utils.format_verse(verse, passage)

            debug_utils.log("Sending verse: {}", [verseMsg])

            telegram_utils.send_msg(user=userObj.get_uid(), text=verseMsg)


def get():
    return [DailyTMSHook()]
