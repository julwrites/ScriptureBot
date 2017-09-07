
# Local modules
from common import debug, text_utils
from common.action import action_classes
from common.telegram import telegram_utils

from devo import mcheyne_utils

class McheyneDevoAction(action_classes.Action):
    def identifier(self):
        return '/mcheynedevo'
    
    def name(self):
        return 'Mcheyne Bible Reading Plan'

    def resolve(self, userObj, msg):
        passages = mcheyne_utils.get_mcheyne(userObj.get_version())

        for passage in passages:
            telegram_utils.send_msg(passage, userObj.get_uid())

        return True

def get():
    return [
        McheyneDevoAction()
    ]

