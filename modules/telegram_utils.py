
# Local modules
import debug
import admin
import chrono
from modules import telegram

from bible_user import *

from constants import *

def foreach_user(fn):
    debug.log('Running ' + str(fn) + ' for each user')
    
    # Read user database
    query = get_user_query()
    query.filter('active =', True)

    try:
        for user in query.run(batch_size=500):
            fn(database.get_uid(user))
    except Exception as e:
        debug.log(str(e))
 
def blast_msg(msg):
    debug.log('Blasting message: ' + msg)

    def send(uid):
        if not debug.debug() or admin.access(uid):
            telegram.send_msg(msg, uid)
    
    foreach_user(send)


# Telegram message parsing
def parse_payload(msg):
    if msg is None:
        return None

    text = msg.get('text')
    if text is not None:
        return text

    audio = msg.get('audio')
    if audio is not None:
        return audio

    document = msg.get('document')
    if document is not None:
        return document

    photo = msg.get('photo')
    if photo is not None:
        return photo

    sticker = msg.get('sticker')
    if sticker is not None:
        return sticker

    video = msg.get('video')
    if video is not None:
        return video

    voice = msg.get('voice')
    if voice is not None:
        return voice

    return None


# Telegram message prettifying
def surround(text, front, back = None):
    if back is None:
        back = front

    return front + text + back

def bold(text):
    return surround(text, '*')

def italics(text):
    return surround(text, '_')

def bracket(text):
    return surround(text, '(', ')')

def bracket_square(text):
    return surround(text, '[', ']')

def link(text, link):
    return bracket_square(text) + bracket(link)

def join(blocks, separator):
    return separator.join(blocks)