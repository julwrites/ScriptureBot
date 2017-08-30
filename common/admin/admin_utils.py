
# Local modules
from common import debug

from secret import BOT_ADMIN

def access(user_id):
    debug.log('Admin Check for ' + str(user_id))

    if str(user_id) == str(BOT_ADMIN):
        return True
    return False

