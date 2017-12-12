
# coding=utf-8

# Local modules
from common import debug, text_utils
from common.telegram import telegram_utils

from secret import BOT_ADMIN

def access(userId):
    debug.log("Admin Check for " + unicode(userId))

    if unicode(userId) == unicode(BOT_ADMIN):
        return True
    return False
