# Python std modules
import json

# Google App Engine API
from google.appengine.api import urlfetch

# Local modules
from common import debug

from secret import BOT_ID

TELEGRAM_URL = 'https://api.telegram_utils.org/bot' + BOT_ID
TELEGRAM_URL_SEND = TELEGRAM_URL + '/sendMessage'

JSON_HEADER = {'Content-Type': 'application/json;charset=utf-8'}

class TelegramPost():
    def __init__(self, id):
        self.format_data = {
            'chat_id': str(id), 
            'parse_mode': 'Markdown'
        }

    def send(self):
        data = json.dumps(self.format_data)
        debug.log('Performing send: ' + data)

        try:
            result = urlfetch.fetch(
                url=TELEGRAM_URL_SEND, 
                payload=data,
                method=urlfetch.POST, 
                headers=JSON_HEADER
                )
        except:
            debug.log('Send failed! ' + TELEGRAM_URL_SEND + ', ' + data)

    def add_text(self, msg):
        debug.log('Adding text for ' + str(id) + ': ' + msg)

        self.format_data['text'] = msg

    def add_keyboard(self, keyboard=[], one_time=False):
        debug.log('Adding keyboard for ' + str(id) + ': ' + str(keyboard))

        self.format_data['reply_markup'] = {
            'keyboard': keyboard,
            'one_time_keyboard': one_time
        }
    
    def close_keyboard(self):
        debug.log('Removing keyboard for ' + str(id))

        self.format_data['reply_markup'] = {
            'remove_keyboard': True
        }

    def add_inline_keyboard(self, keyboard=[]):
        debug.log('Adding inline keyboard for ' + str(id) + ': ' + str(keyboard))
       
        self.format_data['reply_markup'] = {
            'inline_keyboard': keyboard
        }

