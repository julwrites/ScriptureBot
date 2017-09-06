
# Local modules
from common import debug

from secret import BOT_ADMIN

def access(userId):
    debug.log('Admin Check for ' + str(userId))

    if str(userId) == str(BOT_ADMIN):
        return True
    return False

