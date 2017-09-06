
# Python modules
import json

# Local modules
from common import debug

# Defines an interface for all functionality that can be executed by the bot
class Action():
    # To be inherited and overwritten with a check for whether this is waiting for a response
    def waiting(self, userObj):
        if userObj.get_state() == self.identifier():
            debug.log('Waiting for ' + self.identifier())
            return True
        return False

    # Do not overwrite if possible, this performs basic checks and resolves the action
    def execute(self, userObj, msg):
        try:
            if userObj is not None:
                if self.match(msg) or self.waiting(userObj):
                    debug.log('Action being executed: ' + self.identifier())
                    return self.resolve(userObj, msg)
        except:
            debug.log('Execute failed! ' + self.identifier())
            return False

    # Do not overwrite if possible, this checks the message text against the command name
    def match(self, msg):
        msgText = msg.get('text').strip() 
        if msgText.find(self.identifier()) != -1:
            debug.log('Matched with ' + self.identifier())
            return True
        return False



    # To be inherited and overwritten with the name of this action
    def identifier(self):
        return ''

    # To be inherited and overwritten with functionality
    def resolve(self, userObj, msg):
        return False
