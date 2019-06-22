# coding=utf-8

# Python modules
import json

# Local modules
from common.utils import debug_utils, text_utils


# Defines an interface for all functionality that can be executed by the bot
class Action():
    # To be inherited and overwritten with a check for whether this is waiting for a response
    def waiting(self, userObj):
        if userObj is not None:
            if userObj.get_state() == self.identifier():
                debug_utils.log("Waiting for {}", [self.identifier()])
                return True
        return False

    # Do not overwrite if possible, this checks the message text against the command name
    def match(self, msg):
        return self.match_command(msg) or self.match_name(msg)

    def match_command(self, msg):
        if msg is not None:
            msgText = msg.get("text").strip()
            if text_utils.overlap_compare(msgText, self.identifier()):
                debug_utils.log("Matched with {}", [self.identifier()])
                return True
        return False

    def match_name(self, msg):
        if msg is not None:
            msgText = msg.get("text").strip()
            if text_utils.text_compare(msgText, self.name()):
                debug_utils.log("Matched with {}", [self.name()])
                return True
        return False

    def try_execute(self, userObj, msg):
        try:
            if self.waiting(userObj) or self.match(msg):
                return self.resolve(userObj, msg)
        except Exception as e:
            debug_utils.log("Tried, but failed to execute {}", [self.identifier()])
            debug_utils.err(e)
        return False

    # To be inherited if this action is to be exposed as a command
    def is_command(self):
        return False

    # To be inherited and overwritten with the command name of this action
    def identifier(self):
        return ""

    # To be inherited and overwritten with the display name of this action
    def name(self):
        return ""

    # To be inherited and overwritten with the description of this action
    def description(self):
        return ""

    # To be inherited and overwritten with functionality
    def resolve(self, userObj, msg):
        return False
