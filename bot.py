# Python std modules
import webapp2
import json

# Google App Engine modules
from google.appengine.api import urlfetch

# Local modules
from common import debug
from common.telegram import telegram_utils
from user import user_utils

import components

from secret import BOT_ID
APP_BOT_URL = "/" + BOT_ID

CMD_START = '/start'
CMD_START_PROMPT = 'Hello {}, I\'m Biblica! I hope I will be helpful as a tool for you to handle the Bible!'

# This is a special command, specialized to this bot
def start(msg):
    # Register User
    userJson = msg.get('from')
    userId = user_utils.get_uid(userJson.get('id'))
    userObj = user_utils.get_user(userId)

    # This runs to update the user's info, or register
    if userJson is not None:
        debug.log('Updating user info')
        user_utils.set_profile(
            userJson.get('id'), 
            userJson.get('username'), 
            userJson.get('first_name'), 
            userJson.get('last_name'))

    # If this is the user's first time registering
    if userObj is None:
        debug.log_cmd('New user!')
        userObj = user_utils.get_user(userId)

        telegram_utils.send_msg(CMD_START_PROMPT.format(userObj.get_name_string()), userObj.get_uid())
        debug.log('Registering ' + userObj.get_name_string())

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

            # TODO: Replace?
            # Runs to register new users
            start(msg)

            # Read the user to echo back
            userId = user_utils.get_uid(msg.get('from').get('id'))
            userObj = user_utils.get_user(userId)

            actions = components.actions()

            for action in actions:
                if action.execute(userObj, msg):
                    return

            telegram_utils.send_msg('Hello I am bot', msg.get('from').get('id'))


app = webapp2.WSGIApplication([
    (APP_BOT_URL, BotHandler),
], debug=True)
