# coding=utf-8

# Python std modules
import json

# Local modules
from common import debug, text_utils
from common.telegram import telegram_classes

MAX_LENGTH = 4096

TELEGRAM_URL = "https://api.telegram.org/bot" + BOT_ID
TELEGRAM_URL_SEND = TELEGRAM_URL + "/sendMessage"

JSON_HEADER = {"Content-Type": "application/json;charset=utf-8"}

# Telegram message sending functionality
def send_post(userId, post):
    data = json.dumps(post.jsonify())
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

def last_md(chunk):
    start = chunk.rfind("\n")
    end = chunk.rfind("\n")
    if end < start:
        return start

    start = chunk.rfind("_ ")
    end = chunk.rfind(" _")
    if end < start:
        return start

    start = chunk.rfind("* ")
    end = chunk.rfind(" *")
    if end < start:
        return start

    return -1

def format_msg(msg):
    debug.log("Splitting up message if necessary: " +  msg)

    last = None
    chunks = []
    while len(msg) > MAX_LENGTH:
        last = last_md(msg[:MAX_LENGTH])

        if last <= 0:
            last = min(last, msg.rfind(" ", 0, MAX_LENGTH))
        if last <= 0:
            last = MAX_LENGTH

        debug.log("Chunk: " + msg[:last])
        chunks.append(msg[:last])
        msg = msg[last:]
        last = None

    chunks.append(msg[last:])

    return chunks

def send_msg(user, text):
    chunks = format_msg(text)

    for chunk in chunks:
        post = telegram_classes.Post()
        post.add_text(chunk)
        send_post(user, post)

def send_keyboard(user, text, keyboard):
    post = telegram_classes.Post()
    post.add_text(text)
    post.add_keyboard(keyboard)
    send_post(user, post)

def make_reply_button(text="", contact=False, location=False):
    button = telegram_classes.Button()
    button.add_text(text)
    button.field("request_contact", contact)
    button.field("request_location", location)
    return button

def make_reply_keyboard(buttons=[], width=None, resize=False, one_time=False, select=False):
    keyboard = telegram_classes.ReplyKeyboard()
    for button in buttons:
        keyboard.add_button(button)
    if width is not None:
        keyboard.set_width(width)
    keyboard.field("resize_keyboard", resize)
    keyboard.field("one_time_keyboard", one_time)
    keyboard.field("select", select)
    return keyboard

def make_inline_button(text="", url="", callback="", query=""):
    button = telegram_classes.Button()
    button.add_text(text)
    button.field("url", url)
    button.field("callback_data", callback)
    button.field("switch_inline_query", query)
    return button

def make_inline_keyboard(buttons=[], width=None):
    keyboard = telegram_classes.InlineKeyboard()
    for button in buttons:
        keyboard.add_button(button)
    if width is not None:
        keyboard.set_width(width)
    return keyboard

def make_close_keyboard():
    keyboard = telegram_classes.CloseKeyboard()
    return keyboard


# Telegram message parsing
def parse_payload(msg):
    if msg is None:
        return None

    text = msg.get("text")
    if text is not None:
        return text

    audio = msg.get("audio")
    if audio is not None:
        return audio

    document = msg.get("document")
    if document is not None:
        return document

    photo = msg.get("photo")
    if photo is not None:
        return photo

    sticker = msg.get("sticker")
    if sticker is not None:
        return sticker

    video = msg.get("video")
    if video is not None:
        return video

    voice = msg.get("voice")
    if voice is not None:
        return voice

    return None

def strip_command(msg, cmd):
    return msg.get("text").strip().replace(cmd, "")


# Telegram message prettifying
def surround(text, front, back = None):
    if back is None:
        back = front

    return front + text + back

def bold(text):
    return surround(text, "* ", " *")

def italics(text):
    return surround(text, "_ ", " _")

def bracket(text):
    return surround(text, "(", ")")

def bracket_square(text):
    return surround(text, "[", "]")

def link(text, hyperlink):
    return bracket_square(text) + bracket(hyperlink)

def join(blocks, separator):
    return separator.join(blocks)


# Telegram special symbols
def tick():
    return u"\u2714"

def to_sup(text):
    sups = {u"0": u"\u2070",
            u"1": u"\xb9",
            u"2": u"\xb2",
            u"3": u"\xb3",
            u"4": u"\u2074",
            u"5": u"\u2075",
            u"6": u"\u2076",
            u"7": u"\u2077",
            u"8": u"\u2078",
            u"9": u"\u2079",
            u"-": u"\u207b"}
    return "".join(sups.get(char, char) for char in text)
