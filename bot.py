# Python std modules
import webapp2
import json

# Google App Engine modules
from google.appengine.api import urlfetch

# Local modules
from common import debug
from common import telegram
from common import user
from common import admin

import bible
import tms


from secret import BOT_ID
APP_BOT_URL = "/" + BOT_ID

CMD_START = '/start'
CMD_START_PROMPT = 'Hello {}, I\'m Biblica! I hope I will be helpful as a tool for you to handle the Bible!'

# This is a special command, specialized to this bot
def cmd_start(cmd, msg):
    # Register User
    user_json = msg.get('from')
    uid = user.utils.get_uid(user_json.get('id'))
    user = user.utils.get_user(uid)

    # This runs to update the user's info, or register
    if user_json is not None:
        debug.log('Updating user info')
        user.utils.set_profile(
            user_json.get('id'), 
            user_json.get('username'), 
            user_json.get('first_name'), 
            user_json.get('last_name'))

    # If this is the user's first time registering
    if user is None:
        debug.log_cmd(cmd)
        user = user.utils.get_user(uid)

        telegram.utils.send_msg(CMD_START_PROMPT.format(user.get_name_string()), user.get_uid())
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

            # Read the user to echo back
            uid = user.utils.get_uid(msg.get('from').get('id'))
            user = user.utils.get_user(uid)

            actions = tms.actions.get() + bible.actions.get() + user.actions.get()

            for action in actions:
                if action.execute(user, msg):
                    return

            telegram.utils.send_msg('Hello I am bot', msg.get('from').get('id'))


app = webapp2.WSGIApplication([
    (APP_BOT_URL, BotHandler),
], debug=True)
