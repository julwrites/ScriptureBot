
# Local modules
from common.telegram import telegram_utils
from common.action import hook_classes

from devo import mcheyne_utils

class McheyneDevoHook(hook_classes.Hook):
    def identifier(self):
        return '/mcheyne'

    def name(self):
        return 'Mcheyne Bible Reading Plan'

    def resolve(self, userObj):
        passages = mcheyne_utils.get_mcheyne(userObj.get_version())

        if passages is not None:

            for passage in passages:
                telegram_utils.send_msg(passage, userObj.get_uid())
        return True

def get():
    return [
        McheyneDevoHook()
    ]