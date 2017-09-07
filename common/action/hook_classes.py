
# Local modules
from common import debug
from common.action import action_classes

from user import user_utils

# A hook is an action which is dispatched on all users.
# It checks for a subscription, and if the subscription matches itself, it executes the action.
class Hook():
    # Do not overwrite if possible, this performs basic checks and resolves the action
    def execute(self, userObj):
        try:
            if userObj is not None:
                if self.match(userObj):
                    debug.log_hook(self.identifier())
                    self.resolve(userObj)
        except:
            debug.log('Hook failed! ' + self.identifier())

    def dispatch(self):
        user_utils.for_each_user(self.execute)

    # Do not overwrite if possible, this checks the message text against the command name
    def match(self, userObj):
        subs = userObj.get_subscription()
        if subs.find(self.identifier()) != -1:
            debug.log('Matched with ' + self.identifier())
            return True
        return False



    # To be inherited and overwritten with the command name of this action
    def identifier(self):
        return ''

    # To be inherited and overwritten with the display name of this action
    def name(self):
        return ''

    # To be inherited and overwritten with functionality
    def resolve(self, userObj):
        return False