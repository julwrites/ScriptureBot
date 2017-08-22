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
TELEGRAM_KEYBOARD_GRID_SIZE = 3

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
            return
        else:
            debug.log(data)

    def add_text(self, msg):
        debug.log('Adding text for ' + str(id) + ': ' + msg)

        self.format_data['text'] = msg

    def add_keyboard(self, options=[]):
        debug.log('Adding keyboard for ' + str(id) + ': ' + str(options))

        size = len(options)
        keyboard_data = []
        for i in range(0, size, TELEGRAM_KEYBOARD_GRID_SIZE):
            keyboard_row = []

            for j in range(0, TELEGRAM_KEYBOARD_GRID_SIZE):
                data = options[i+j]
                keyboard_row.append({'text': data})
            
            keyboard_data.append(keyboard_row)
        
        self.format_data['reply_markup'] = [
            {
                'keyboard': keyboard_data
            }
        ]

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