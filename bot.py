# Python std modules
import webapp2
import json

# Google App Engine modules
from google.appengine.api import urlfetch

# Local modules
from common import debug
from common import telegram
from common import user

import bible
import tms


from secret import BOT_ID
APP_BOT_URL = "/" + BOT_ID

CMD_START = '/start'
CMD_START_PROMPT = 'Hello {}, I\'m Biblica! I hope I will be helpful as a tool for you to handle the Bible!'

# This is a special command, specialized to this bot
def start(msg):
    # Register User
    user_json = msg.get('from')
    user_id = user.utils.get_uid(user_json.get('id'))
    user_obj = user.utils.get_user(user_id)

    # This runs to update the user's info, or register
    if user_json is not None:
        debug.log('Updating user info')
        user.utils.set_profile(
            user_json.get('id'), 
            user_json.get('username'), 
            user_json.get('first_name'), 
            user_json.get('last_name'))

    # If this is the user's first time registering
    if user_obj is None:
        debug.log_cmd('New user!')
        user_obj = user.utils.get_user(user_id)

        telegram.utils.send_msg(CMD_START_PROMPT.format(user_obj.get_name_string()), user_obj.get_uid())
        debug.log('Registering ' + user_obj.get_name_string())

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
            user_id = user.utils.get_uid(msg.get('from').get('id'))
            user_obj = user.utils.get_user(user_id)

            actions = tms.actions.get() + bible.actions.get() + user.actions.get()

            for action in actions:
                if action.execute(user_obj, msg):
                    return

            telegram.utils.send_msg('Hello I am bot', msg.get('from').get('id'))


app = webapp2.WSGIApplication([
    (APP_BOT_URL, BotHandler),
], debug=True)
