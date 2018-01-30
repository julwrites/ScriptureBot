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


def set_intersect(lhs, rhs):
    rhs = set(rhs)
    both = [item for item in lhs if item in rhs]
    non = [item for item in lhs if not item in rhs]
    return both, non


def find_symbols(text, symbols):
    sym = []

    for symbol in symbols:
        curr = 0
        slength = len(symbol)
        while True:
            pos = text[curr:].find(symbol)
            if pos == -1:
                break
            pos += curr
            curr = pos + slength

            sym.append(pos)

    return sym


def find_symbol_pairs(text, symbols, exclusive=True):
    sym = []

    for symbol in symbols:
        curr = 0
        slength = len(symbol)
        while True:
            first = text[curr:].find(symbol)
            if first == -1:
                break
            first += curr
            curr = first + slength

            last = text[curr:].find(symbol)
            if last == -1:
                break
            last += curr
            curr = last + (slength if exclusive else 0)

            sym.append(first)
            sym.append(last)

    return sym


def format_msg(msg):
    debug.log("Formatting message")

    symbols = find_symbols(msg, ["_", "*"])
    pairs = find_symbol_pairs(msg, ["_", "*"])
    md, esc = set_intersect(symbols, pairs)

    for symbol in esc:
        msg = msg[:symbol] + "\\" + msg[symbol:]

    return msg


def split_msg(msg):
    debug.log("Splitting up message if necessary")

    chunks = []

    symbols = find_symbols(msg, ["_", "*"])
    pairs = find_symbol_pairs(msg, ["_", "*"])
    md, esc = set_intersect(symbols, pairs)
    pairs = find_symbol_pairs(msg, ["\n"], False)
    md.append(pairs)

    max_pos = 0

    while len(msg[max_pos:]) > MAX_LENGTH:
        curr = max_pos
        max_pos += MAX_LENGTH

        for i in range(0, len(md), 2):
            if md[i] < max_pos and md[i + 1] >= max_pos:
                max_pos = md[i]

        debug.log("Chunk: " + msg[curr:max_pos])

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
