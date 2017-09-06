
# Local modules
from common import debug
from common.actions import action_classes

from common.user import user_utils

class Hook(action_classes.Action):
    def execute(self, data):
        try:
            user_utils.for_each_user(self.resolve)
        except:
            debug.log('Hook failed! ' + self.identifier())
            return False



    # To be inherited and overwritten with functionality
    def resolve(self, userObj):
        return False