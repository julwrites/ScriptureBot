# coding=utf-8

# Local modules
from common import debug
from common.telegram import telegram_utils
from common.action import hook_classes

from bible import bible_utils


class DailyVerseHook(hook_classes.Hook):
    def identifier(self):
        return "/dailyverse"

    def name(self):
        return "Daily Verse"

    def resolve(self, userObj):
        if userObj is not None:
            debug.log("Sending verse: " + "")


def get():
    return [DailyVerseHook()]
