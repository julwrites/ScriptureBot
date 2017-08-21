# Python std modules
import json
import string

# Google App Engine API
from google.appengine.api import urlfetch

# Local modules
from common import debug
from common.constants import *

def send_msg(msg, id):
    debug.log(str(id) + ': ' +  msg)

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
        data_prep = {
            'chat_id': str(id), 
            'text': chunk,
            'parse_mode': 'Markdown'
        }
        data = json.dumps(data_prep)

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
