
# Local modules
from common import debug
from common import chrono
from common.telegram_utils.telegram_class import TelegramPost

from common.user.bibleuser_utils import *

from common.constants import *

TELEGRAM_MAX_LENGTH = 4096

TELEGRAM_OPTION_REPLY_KEYBOARD = 'reply_markup'
TELEGRAM_KEYBOARD_GRID_SIZE = 3

# Telegram message sending functionality
def format_keyboard(options=[]):
    num_buttons = len(options)
    modulus = 1 if num_buttons % TELEGRAM_KEYBOARD_GRID_SIZE else 0
    num_rows = int(num_buttons / TELEGRAM_KEYBOARD_GRID_SIZE) + modulus

    keyboard_data = []
    for i in range(0, num_rows):
        keyboard_row = []

        for j in range(0, TELEGRAM_KEYBOARD_GRID_SIZE):
            if num_buttons == 0:
                break

            data = options[i * TELEGRAM_KEYBOARD_GRID_SIZE + j]
            keyboard_row.append({'text': data})
            num_buttons -= 1
        
        keyboard_data.append(keyboard_row)

    return keyboard_data

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

def send_msg_keyboard(msg, id, options=[], inline=False, one_time=False):
    post = TelegramPost(id)
    post.add_text(msg)
    if inline:
        post.add_inline_keyboard(options)
    else:
        post.add_keyboard(options, one_time)
    post.send()

def send_close_keyboard(msg, id):
    post = TelegramPost(id)
    post.add_text(msg)
    post.close_keyboard()
    post.send()


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