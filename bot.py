# Python std modules
import webapp2
import json

# Google App Engine modules
from google.appengine.api import urlfetch

# Local modules
from common import database
from common import debug
from common.telegram import telegram_utils
from common import admin_utils
from tms import tms_utils

from common.user.bibleuser_utils import *

from common import bot_commands
from common import admin_commands
from bible import bible_commands
from tms import tms_commands
from common.user import bibleuser_commands

from secret import BOT_ID
APP_BOT_URL = "/" + BOT_ID

CMD_START = '/start'
CMD_START_PROMPT = 'Hello {}, I\'m Biblica! I hope I will be helpful as a tool for you to handle the Bible!'

# This is a special command, specialized to this bot
def cmd_start(cmd, msg):
    # Register User
    user_json = msg.get('from')
    uid = get_uid(user_json.get('id'))
    user = get_user(uid)

    # This runs to update the user's info, or register
    if user_json is not None:
        debug.log('Updating user info')
        set_profile(
            user_json.get('id'), 
            user_json.get('username'), 
            user_json.get('first_name'), 
            user_json.get('last_name'))

    # If this is the user's first time registering
    if user is None:
        debug.log_cmd(cmd)
        user = get_user(uid)

        telegram_utils.send_msg(CMD_START_PROMPT.format(user.get_name_string()), user.get_uid())
        debug.log('Registering ' + user.get_name_string())

        return True
    return False

class BotHandler(webapp2.RequestHandler):
    def get(self):
        self.post()

    def post(self):
        data = json.loads(self.request.body)
        debug.log(data)

        if data.get('message'):
            msg = data.get('message')

            cmd = self.read_cmd(msg.get('text'))
            if cmd is not None:
                if self.handle_command(cmd, msg):
                    return

            if self.handle_state(msg):
                return

            telegram_utils.send_msg('Hello I am bot', msg.get('from').get('id'))
            
    def read_cmd(self, text):
        debug.log('Message Text: ' + text)

        if text.startswith("/"):
            cmd_end = text.find(' ')
            if cmd_end == -1:
                return text

            return text[:cmd_end]

        return None

    def handle_command(self, cmd, msg):
        debug.log('Possible command detected: ' + cmd)

        # Read the user to echo back
        uid = get_uid(msg.get('from').get('id'))
        user = get_user(uid)

        if admin_commands.cmds(uid, cmd, msg):
            return True

        if user is None:
            debug.log('This user does not exist')

        if cmd_start(cmd, msg):
            return True
        else:
            debug.log('Running all commands')

            if( \
            bibleuser_commands.cmds(user, cmd, msg) \
            or bot_commands.cmds(user, cmd, msg)    \
            or bible_commands.cmds(user, cmd, msg)    \
            or tms_commands.cmds(user, cmd, msg)    \
            ):
                return True

        return False

    def handle_state(self, msg):
        debug.log('Handling state reaction')

        # Read the user to echo back
        uid = get_uid(msg.get('from').get('id'))
        user = get_user(uid)

        if admin.access(uid):
            debug.log('Welcome, Master')

        if user is None:
            debug.log('This user does not exist')
        else:
            debug.log('Running all states')
            
            # States
            if (    \
            bibleuser_commands.states(user, msg)    \
            or bot_commands.states(user, msg)       \
            or bible_commands.states(user, msg)       \
            or tms_commands.states(user, msg)       \
            ):
                return True

        return False

app = webapp2.WSGIApplication([
    (APP_BOT_URL, BotHandler),
], debug=True)
