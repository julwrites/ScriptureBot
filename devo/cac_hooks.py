
# Local modules
from common import debug, text_utils
from common.action import hook_classes
from common.telegram import telegram_utils

from devo import cac_utils
from bible import bible_utils

CMD_PASSAGE_PROMPT = "Give me a Bible reference"
CMD_PASSAGE_BADQUERY = "Sorry, I can't find this reference"


class CACDevoHook(hook_classes.Hook):
    def identifier(self):
        return '/cacdevo'

    def name(self):
        return 'Center for Action and Contemplation Devotional'

    def resolve(self, userObj):
        devo = cac_utils.get_devo(userObj.get_version())
        debug.log('Devo fetched: ' + devo)

        return True

def get():
    return [
        CACDevoHook()
    ]