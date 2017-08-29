
# Python modules
import json

# Defines an interface for all functionality that can be executed by the bot
class Action():
    # Do not overwrite if possible, this checks the message text against the command name
    def match(self, msg_text):
        return msg_text.find(self.identifier()) != -1

    # To be inherited and overwritten with a check for whether this is waiting for a response
    def waiting(self, user):
        return user.get_state() == self.identifier()

    # Do not overwrite if possible, this performs basic checks and resolves the action
    def execute(self, user, msg):
        try:
        if user is not None:
            if self.match(msg.get('text').strip()) or self.waiting(user):
                return self.resolve(user, msg)
        return False


    # To be inherited and overwritten with the name of this action
    def identifier(self):
        return ''

    # To be inherited and overwritten with functionality
    def resolve(self, user, msg):
        return False

 