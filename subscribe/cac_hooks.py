
# Local modules
from common import debug, text_utils
from common.action import hook_classes
from common.telegram import telegram_utils

from devo import cac_utils
from bible import bible_utils

class CACDevoHook(hook_classes.Hook):
    def identifier(self):
        return '/cacdevo'

    def name(self):
        return 'Center for Action and Contemplation Devotional'

    def resolve(self, userObj):
        devo = cac_utils.get_devo(userObj.get_version())
        debug.log('devo fetched: ' + devo)

        return True

def get():
    return [
        CACDevoHook()
    ]