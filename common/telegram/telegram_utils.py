# coding=utf-8

# Python std modules
import json

# Local modules
from common.utils import debug_utils, text_utils, web_utils
from common.telegram import telegram_classes

from secret import BOT_ID

MAX_LENGTH = 4096

TELEGRAM_URL = "https://api.telegram.org/bot" + BOT_ID
TELEGRAM_URL_SEND = TELEGRAM_URL + "/sendMessage"

JSON_HEADER = {"Content-Type": "application/json;charset=utf-8"}


# Telegram message sending functionality
def send_post(post):
    data = json.dumps(post.data())

    debug_utils.log("Performing send: {}", [data])

    try:
        web_utils.post_http(TELEGRAM_URL_SEND, data, JSON_HEADER)
    except Exception as e:
        debug_utils.err(e)


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


def fix_sentences(text):
    # Special cases

    # Quotes should always be preceeded by a space when open, and when closed
    text = text.replace('"', ' " ')
    text = text.replace("-", " - ")

    # Ensuring spaces after
    text = text.replace(":", ": ")
    text = text.replace(";", "; ")
    text = text.replace(",", ", ")
    text = text.replace(".", ". ")
    text = text.replace("!", "! ")
    text = text.replace("?", "? ")

    # Ensuring no spaces before
    text = text.replace(" :", ":")
    text = text.replace(" ;", ";")
    text = text.replace(" ,", ",")
    text = text.replace(" .", ".")
    text = text.replace(" !", "!")
    text = text.replace(" ?", "?")

    # Removing extra spaces
    text = text.replace("    ", " ")
    text = text.replace("   ", " ")
    text = text.replace("  ", " ")
    return text


def format_msg(msg):
    debug_utils.log("Formatting message")

    msg = fix_sentences(msg)
    symbols = find_symbols(msg, ["_", "*"])
    pairs = find_symbol_pairs(msg, ["_", "*"])
    md, esc = set_intersect(symbols, pairs)

    for symbol in esc:
        msg = msg[:symbol] + "\\" + msg[symbol:]

    return msg


def split_msg(msg):
    debug_utils.log("Splitting up message if necessary")

    chunks = []

    symbols = find_symbols(msg, ["_", "*"])
    pairs = find_symbol_pairs(msg, ["_", "*"])
    md, esc = set_intersect(symbols, pairs)

    seps = find_symbols(msg, ["\n"])

    debug_utils.log("Markdown pairs: {}", [md])

    end_pos = 0

    while len(msg[end_pos:]) > MAX_LENGTH:
        curr = end_pos
        max_pos = end_pos + MAX_LENGTH
        end_pos = max_pos

        for i in range(0, len(md), 2):
            if md[i] < max_pos and md[i + 1] >= max_pos:
                end_pos = md[i]

        for sep in reversed(seps):
            if end_pos > sep:
                end_pos = sep
                break

        chunk = msg[curr:end_pos]
        debug_utils.log("Chunk: {}", [chunk])
        chunks.append(chunk)

    chunks.append(msg[end_pos:])

    return chunks


def send_msg(user, text, args=[]):
    debug_utils.log("Preparing to send {}: {}", [user, text])

    fmt_msg = format_msg(text)

    if len(args) > 0:
        debug_utils.log("Detected arguments: {}", [args])
        fmt_msg = fmt_msg.format(*[text_utils.to_utf8(arg) for arg in args])

    chunks = split_msg(fmt_msg)

    for chunk in chunks:
        post = telegram_classes.Post()
        post.set_user(user)
        post.set_text(chunk)
        send_post(post)


def send_reply(user, text, reply):
    fmt_msg = format_msg(text)

    post = telegram_classes.Post()
    post.set_user(user)
    post.set_text(fmt_msg)
    post.set_reply(reply)
    send_post(post)


def make_reply_button(text="", contact=False, location=False):
    fmt_msg = format_msg(text)

    button = telegram_classes.Markup()
    button.set_text(fmt_msg)
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
    fmt_msg = format_msg(text)

    button = telegram_classes.Markup()
    button.set_text(fmt_msg)
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
    return msg.get("text").replace(cmd, "").strip()


def strip_symbol(msg, sym):
    return msg.get("text").replace(sym, "").strip()


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
    return "\\u2714"


def to_sup(text):
    sups = {
        "0": "\\u2070",
        "1": "\xb9",
        "2": "\xb2",
        "3": "\xb3",
        "4": "\\u2074",
        "5": "\\u2075",
        "6": "\\u2076",
        "7": "\\u2077",
        "8": "\\u2078",
        "9": "\\u2079",
        "-": "\\u207b"
    }
    return "".join(sups.get(char, char) for char in text)
