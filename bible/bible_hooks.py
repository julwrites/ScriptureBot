
# Local modules
from common import debug
from common.telegram import telegram_utils
from common.action import hook_classes

from bible import bible_utils


class BibleActionHook(hook_classes.Hook):
    def identifier(self):
        return '/dailyverse'

    def resolve(self, userObj):
        if userObj is not None:
            debug.log("Sending verse: " + '')

            # telegram_utils.send_msg(verseMsg, userObj.get_uid())

HOOK = BibleActionHook()
def get():
    return HOOK