
# coding=utf-8

# Google App Engine API
from google.appengine.api import urlfetch

# Local modules
from common import debug, text_utils

from secret import BOT_ID


class Markup():
    def __init__(self):
        self.formatData = {}

    def field(key, value):
        self.formatData[key] = value

    def jsonify(self):
        return self.formatData


class Button(Markup):
    def __init__(self):
        super(Button, self).__init__()

        self.text = ""
        self.fields = {}

    def jsonify(self):
        self.field("text", self.text)

        for key in self.fields.keys():
            self.field(key, self.fields[key])

        return self.formatData


KEYBOARD_WIDTH = 3
class Keyboard(Markup):
    def __init__(self):
        super(Keyboard, self).__init__()

        self.buttons = []
        self.fields = {}
        self.width = KEYBOARD_WIDTH

    def add_button(self, button):
        self.buttons.append(button)

    def set_width(self, width):
        self.width = width

    def format(self):
        num = len(self.buttons)
        modulus = 1 if num % self.width else 0
        rows = int(num / self.width) + modulus

        keyboard = []
        for i in range(0, rows):
            row = []

            for j in range(0, self.width):
                if num == 0:
                    break

                button = self.buttons[i * self.width + j]
                row.append(button.jsonify())
                num -= 1

            keyboard.append(row)

        return keyboard


class ReplyKeyboard(Keyboard):
    def __init__(self):
        super(ReplyKeyboard, self).__init__()

    def jsonify(self):
        self.field("keyboard", self.format())

        return self.formatData

class InlineKeyboard(Keyboard):
    def __init__(self):
        super(InlineKeyboard, self).__init__()

    def jsonify(self):
        self.field("inline_keyboard", self.format())

        return self.formatData

class CloseKeyboard(Markup):
    def __init__(self):
        super(CloseKeyboard, self).__init__()

    def jsonify(self):
        self.field("remove_keyboard", True)

        return self.formatData

class Post(Markup):
    def __init__(self):
        super(Post, self).__init__()

        self.formatData = {
            "parse_mode": "Markdown"
        }

    def set_id(self, id):
        return self.formatData["chat_id"] = text_utils.stringify(id)

    def add_text(self, msg):
        debug.log("Adding text for " + self.uid())
        self.field("text", msg)

    def add_keyboard(self, keyboard):
        debug.log("Adding keyboard for " + self.uid() + ": " + text_utils.stringify(keyboard))
        self.field("reply_markup", keyboard.jsonify())

    def jsonify(self):
        return self.formatData


