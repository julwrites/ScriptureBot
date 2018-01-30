# coding=utf-8

# Python std modules
import json

# Google App Engine API
from google.appengine.api import urlfetch

# Local modules
from common import debug, text_utils
from common.telegram import telegram_classes

from secret import BOT_ID

MAX_LENGTH = 4096

TELEGRAM_URL = "https://api.telegram.org/bot" + BOT_ID
TELEGRAM_URL_SEND = TELEGRAM_URL + "/sendMessage"

JSON_HEADER = {"Content-Type": "application/json;charset=utf-8"}


# Telegram message sending functionality
def send_post(post):
    data = json.dumps(post.data())
    debug.log("Performing send: " + data)

    try:
        urlfetch.fetch(
            url=TELEGRAM_URL_SEND,
            payload=data,
            method=urlfetch.POST,
            headers=JSON_HEADER)
    except Exception as e:
        debug.log("Send failed! " + TELEGRAM_URL_SEND + ", " + data)
        debug.err(e)


def find_md(text, symbols):
    md = []
    esc = []

    curr = 0
    for symbol in symbols:
        while True:
            start = text[curr:].find(symbol)
            if start == -1:
                break
            end = text[start:].find(symbol)
            if end == -1:
                esc.append(start)
                break
            md.append((start, end))
            curr = end + 1

    return md, esc


def format_msg(msg):
    debug.log("Formatting message")

    md, esc = find_md(msg, ["_", "*"])
    debug.log(md)
    debug.log(esc)

    for symbol in esc:
        msg = msg[:symbol] + "\\" + msg[symbol:]

    return msg


def split_msg(msg):
    debug.log("Splitting up message if necessary")

    chunks = []
    md, esc = find_md(msg, ["_", "*"])

    end = len(msg)
    max_pos = 0

    while len(msg[max_pos:]) > MAX_LENGTH:
        curr = max_pos
        max_pos += MAX_LENGTH

        for pair in md:
            if pair[0] < MAX_LENGTH and pair[1] > MAX_LENGTH:
                max_pos = pair[0]

        chunks.append(msg[curr:max_pos])
    chunks.append(msg[max_pos:])

    return chunks


def send_msg(user, text):
    debug.log("Preparing to send " + text_utils.stringify(user) + ": " + text)
    chunks = split_msg(format_msg(text))

    for chunk in chunks:
        post = telegram_classes.Post()
        post.set_user(user)
        post.set_text(chunk)
        send_post(post)


def send_reply(user, text, reply):
    post = telegram_classes.Post()
    post.set_user(user)
    post.set_text(format_msg(text))
    post.set_reply(reply)
    send_post(post)


def make_reply_button(text="", contact=False, location=False):
    button = telegram_classes.Markup()
    button.set_text(format_msg(text))
    button.set_field("request_contact", contact)
    button.set_field("request_location", location)
    return button


def make_reply_keyboard(buttons=[],
                        width=None,
                        resize=False,
                        one_time=False,
                        select=False):
    keyboard = telegram_classes.ReplyKeyboard()
    for button in buttons:
        keyboard.add_button(button)
    if width is not None:
        keyboard.set_width(width)
    keyboard.set_field("resize_keyboard", resize)
    keyboard.set_field("one_time_keyboard", one_time)
    keyboard.set_field("select", select)
    return keyboard


def make_inline_button(text="", url="", callback="", query=""):
    button = telegram_classes.Markup()
    button.set_text(format_msg(text))
    button.set_field("url", url)
    button.set_field("callback_data", callback)
    button.set_field("switch_inline_query", query)
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
    return text_utils.stringify(msg.get("text")).strip().replace(cmd, "")


# Telegram message prettifying
def surround(text, front, back=None):
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
    sups = {
        u"0": u"\u2070",
        u"1": u"\xb9",
        u"2": u"\xb2",
        u"3": u"\xb3",
        u"4": u"\u2074",
        u"5": u"\u2075",
        u"6": u"\u2076",
        u"7": u"\u2077",
        u"8": u"\u2078",
        u"9": u"\u2079",
        u"-": u"\u207b"
    }
    return "".join(sups.get(char, char) for char in text)
