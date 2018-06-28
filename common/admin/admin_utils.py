# coding=utf-8

# Local modules
from common import debug, text_utils
from common.telegram import telegram_utils

from secret import BOT_ADMIN


def access(userId):
    debug.log("Admin Check for {}", [userId])

    if text_utils.to_utf8(userId) == text_utils.to_utf8(BOT_ADMIN):
        return True
    return False
