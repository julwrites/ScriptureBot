
# Local modules
from common import debug, text_utils
from common.action import hook_classes
from common.telegram import telegram_utils

from subscribe import cac_utils
from bible import bible_utils

class CACsubscribeHook(hook_classes.Hook):
    def identifier(self):
        return '/cacsubscribe'

    def name(self):
        return 'Center for Action and Contemplation subscribetional'

    def resolve(self, userObj):
        subscribe = cac_utils.get_subscribe(userObj.get_version())
        debug.log('subscribe fetched: ' + subscribe)

        return True

def get():
    return [
        CACsubscribeHook()
    ]