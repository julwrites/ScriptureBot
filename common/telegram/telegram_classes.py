
# coding=utf-8

# Python std modules
import json

# Google App Engine API
from google.appengine.api import urlfetch

# Local modules
from common import debug, text_utils

from secret import BOT_ID

TELEGRAM_URL = "https://api.telegram.org/bot" + BOT_ID
TELEGRAM_URL_SEND = TELEGRAM_URL + "/sendMessage"

JSON_HEADER = {"Content-Type": "application/json;charset=utf-8"}

class TelegramPost():
    def __init__(self, userId):
        self.formatData = {
            "chat_id": text_utils.stringify(userId), 
            "parse_mode": "Markdown"
        }

    def uid(self):
        return self.formatData.get("chat_id")

    def send(self):
        data = json.dumps(self.formatData)
        debug.log("Performing send: " + data)

        try:
            urlfetch.fetch(
                url=TELEGRAM_URL_SEND, 
                payload=data,
                method=urlfetch.POST, 
                headers=JSON_HEADER
                )
        except Exception as e:
            debug.log("Send failed! " + TELEGRAM_URL_SEND + ", " + data)
            debug.err(e)

    def add_text(self, msg):
        debug.log("Adding text for " + self.uid())

        self.formatData["text"] = msg

    def add_keyboard(self, keyboard=[], oneTime=False):
        debug.log("Adding keyboard for " + self.uid() + ": " + text_utils.stringify(keyboard))

        self.formatData["reply_markup"] = {
            "keyboard": keyboard,
            "one_time_keyboard": oneTime
        }
    
    def close_keyboard(self):
        debug.log("Removing keyboard for " + self.uid())

        self.formatData["reply_markup"] = {
            "remove_keyboard": True
        }

    def add_inline_keyboard(self, keyboard=[]):
        debug.log("Adding inline keyboard for " + self.uid() + ": " + text_utils.stringify(keyboard))
       
        self.formatData["reply_markup"] = {
            "inline_keyboard": keyboard
        }

