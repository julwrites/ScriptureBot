# coding=utf-8

# Local modules
from common.utils import debug_utils
from common.action import action_classes

from user import user_utils


# A hook is an action which is dispatched on all users.
# It checks for a subscription, and if the subscription matches itself, it executes the action.
class Hook():
    # Do not overwrite if possible, this performs basic checks and resolves the action
    def try_execute(self, userObj):
        try:
            if userObj is not None:
                if self.match(userObj):
                    debug_utils.log_hook(self.identifier())
                    self.resolve(userObj)
        except Exception as e:
            debug_utils.log("Hook failed! {}", [self.identifier()])
            debug_utils.err(e)

    def dispatch(self):
        user_utils.for_each_user(self.try_execute)

    # Do not overwrite if possible, this checks the message text against the command name
    def match(self, userObj):
        subs = userObj.get_subscription()
        if subs.find(self.identifier()) != -1:
            debug_utils.log("Matched with {}", [self.identifier()])
            return True
        return False

    # To be inherited and overwritten with the command name of this action
    def identifier(self):
        return ""

    # To be inherited and overwritten with the display name of this action
    def name(self):
        return ""

    # To be inherited and overwritten with functionality
    def resolve(self, userObj):
        return False