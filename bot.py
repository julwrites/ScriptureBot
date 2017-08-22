# Python std modules
import webapp2
import json

# Google App Engine modules
from google.appengine.api import urlfetch

# Local modules
from common import database
from common import debug
from common import telegram 
from common import admin
from tms import tms_utils

from user.bibleuser_utils import *

from common import bot_commands
from common import admin_commands
from bgw import bgw_commands
from tms import tms_commands

from common.constants import APP_BOT_URL

CMD_START = '/start'
CMD_START_PROMPT = 'Hello {}, I\'m Biblica! I hope I will be helpful as a tool for you to handle the Bible!'

# This is a special command, specialized to this bot
def cmd_start(cmd, msg):
    if cmd == CMD_START:
        debug.log('Command: ' + cmd)

        # Register User
        user_json = msg.get('from')

        # Read the user to echo back
        uid = get_uid(user_json.get('id'))
        user = get_user(uid)

        # This runs to update the user's info, or register
        if user_json is not None:
            debug.log(str(user_json.get('first_name'))
            + ' ' + str(user_json.get('last_name'))
            + ': ' + str(user_json.get('username'))
            + ' - ' + str(user_json.get('id')))

            set_profile(
                user_json.get('id'), 
                user_json.get('username'), 
                user_json.get('first_name'), 
                user_json.get('last_name'))

        # If this is not the user's first time registering
        if user is None:
            user = get_user(uid)

            # Initializes this user's data
            verse = tms_utils.get_start_verse()
            user.set_current_pack(verse.get_pack())
            user.set_current_verse(verse.get_position())

            telegram.send_msg(CMD_START_PROMPT.format(user.get_name_string()), user.get_uid())
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

            telegram.send_msg('Hello I am bot', msg.get('from').get('id'))
            
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

        if admin.access(uid):
            debug.log('Welcome, Master')
            if admin_commands.cmds(uid, cmd, msg):
                return True

        if user is None:
            debug.log('This user does not exist')

        if cmd_start(cmd, msg):
            return True
        else:
            debug.log('Running all commands')

            if( \
            bot_commands.cmds(user, cmd, msg)    \
            or bgw_commands.cmds(user, cmd, msg)    \
            or tms_commands.cmds(user, cmd, msg)    \
            ):
                return True

        return False

    def handle_state(self, msg):
        # States
        if bot_commands.states(msg):
            return True

        return False

app = webapp2.WSGIApplication([
    (APP_BOT_URL, BotHandler),
], debug=True)
