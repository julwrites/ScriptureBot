# coding=utf-8

# Python std modules
import webapp2
import json

# Local modules
from .common.utils import debug_utils, text_utils
from .common.telegram import telegram_utils
from .common.action import action_utils
from .user import user_utils

from . import actions

from secret import BOT_ID
APP_BOT_URL = "/" + BOT_ID


class BotHandler(webapp2.RequestHandler):
    def get(self):
        self.post()

    def post(self):
        data = json.loads(self.request.body)
        debug_utils.log(data)

        if data.get("message"):
            msg = data.get("message")

            # Read the user to echo back
            userId = user_utils.get_uid(msg.get("from").get("id"))
            userObj = user_utils.get_user(userId)

            if action_utils.execute(actions.get(), userObj, msg):
                return

            telegram_utils.send_msg(
                user=msg.get("from").get("id"), text="Hello, I am bot")


app = webapp2.WSGIApplication([
    (APP_BOT_URL, BotHandler),
], debug=True)
