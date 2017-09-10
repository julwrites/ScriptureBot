
# coding=utf-8

# Local modules
from common import debug, text_utils

from secret import BOT_ADMIN

def access(userId):
    debug.log("Admin Check for " + text_utils.stringify(userId))

    if text_utils.stringify(userId) == text_utils.stringify(BOT_ADMIN):
        return True
    return False

