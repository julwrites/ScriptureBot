# Python std modules
import json
import string

# Google App Engine API
from google.appengine.api import urlfetch

# Local modules
from common import debug

from common.constants import *

TELEGRAM_URL = 'https://api.telegram.org/bot' + BOT_ID

TELEGRAM_URL_SEND = TELEGRAM_URL + '/sendMessage'
TELEGRAM_MAX_LENGTH = 4096

TELEGRAM_OPTION_REPLY_KEYBOARD = 'reply_markup'

class TelegramPost():
    def __init__(self, id):
        self.format_data = {
            'chat_id': str(id), 
            'parse_mode': 'Markdown'
        }

    def send(self):
        data = json.dumps(self.format_data)

        try:
            debug.log('Sending ' + chunk)

            result = urlfetch.fetch(
                url=TELEGRAM_URL_SEND, 
                payload=data,
                method=urlfetch.POST, 
                headers=JSON_HEADER
                )
        except:
            return
        else:
            debug.log(data)

    def add_keyboard(self, options=[]):
        debug.log('Adding keyboard for ' + str(id) + ': ' + str(options))

        keyboard_data = []
        for data in options:
            keyboard_data.append({'text': data})
        
        self.format_data['reply_markup'] = {
            'keyboard': keyboard_data
        }

    def add_text(self, msg):
        debug.log('Adding text for ' + str(id) + ': ' + str(msg))

        self.format_data['text'] = msg

def send_msg(msg, id):
    debug.log('Sending message to ' + str(id) + ': ' +  msg)

    last = None
    chunks = []
    while len(msg) > TELEGRAM_MAX_LENGTH:
        last = msg.rfind(' ', 0, TELEGRAM_MAX_LENGTH)
        if last == -1:
            last = TELEGRAM_MAX_LENGTH

        debug.log('Chunk: ' + msg[:last])
        chunks.append(msg[:last])
        msg = msg[last:]
        last = None

    chunks.append(msg[last:])

    for chunk in chunks:
        post = TelegramPost(id)
        post.add_text(chunk)
        post.send()

def send_msg_keyboard(msg, id, options=[]):
    post = TelegramPost(id)
    post.add_text(msg)
    post.add_keyboard(options)
    post.send()