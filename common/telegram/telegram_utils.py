
# Local modules
from common import debug
from common.telegram.telegram_classes import TelegramPost

MAX_LENGTH = 4096

OPTION_REPLY_KEYBOARD = 'reply_markup'
KEYBOARD_WIDTH = 3

# Telegram message sending functionality
def format_keyboard(options=[], width=KEYBOARD_WIDTH):
    numButtons = len(options)
    modulus = 1 if numButtons % width else 0
    numRows = int(numButtons / width) + modulus

    keyboardData = []
    for i in range(0, numRows):
        keyboardRow = []

        for j in range(0, width):
            if numButtons == 0:
                break

            data = options[i * width + j]
            keyboardRow.append({'text': data})
            numButtons -= 1
        
        keyboardData.append(keyboardRow)

    return keyboardData

def send_msg(msg, userId):
    debug.log('Sending message to ' + str(userId) + ': ' +  msg)

    last = None
    chunks = []
    while len(msg) > MAX_LENGTH:
        last = msg.rfind(' ', 0, MAX_LENGTH)
        if last == -1:
            last = MAX_LENGTH

        debug.log('Chunk: ' + msg[:last])
        chunks.append(msg[:last])
        msg = msg[last:]
        last = None

    chunks.append(msg[last:])

    for chunk in chunks:
        post = TelegramPost(userId)
        post.add_text(chunk)
        post.send()

def send_msg_keyboard(msg, userId, options=[], width=KEYBOARD_WIDTH, inline=False, oneTime=False):
    post = TelegramPost(userId)
    post.add_text(msg)
    if inline:
        post.add_inline_keyboard(format_keyboard(options, width))
    else:
        post.add_keyboard(format_keyboard(options, width), oneTime)
    post.send()

def send_close_keyboard(msg, userId):
    post = TelegramPost(userId)
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

def strip_command(msg, cmd):
    return msg.get('text').strip().replace(cmd, '')


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