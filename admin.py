
# Local modules
import debug

from secret import BOT_ADMIN

def access(uid):
    debug.log('Admin Check for ' + str(uid))

    if str(uid) == str(BOT_ADMIN):
        return True
    return False

