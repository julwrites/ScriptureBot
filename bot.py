
# coding=utf-8

# Python std modules
import webapp2
import json

# Google App Engine modules
from google.appengine.api import urlfetch

# Local modules
from common import debug
from common.telegram import telegram_utils
from common.action import action_utils
from user import user_utils

import actions

from secret import BOT_ID
APP_BOT_URL = "/" + BOT_ID

class BotHandler(webapp2.RequestHandler):
    def get(self):
        self.post()

    def post(self):
        data = json.loads(self.request.body)
        debug.log(data)

        if data.get("message"):
            msg = data.get("message")

            # Read the user to echo back
            userId = user_utils.get_uid(msg.get("from").get("id"))
            userObj = user_utils.get_user(userId)

            if action_utils.execute(actions.get(), userObj, msg):
                return

            telegram_utils.send_msg("Hello I am bot", msg.get("from").get("id"))


app = webapp2.WSGIApplication([
    (APP_BOT_URL, BotHandler),
], debug=True)
