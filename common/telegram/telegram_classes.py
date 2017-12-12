
# coding=utf-8

# Google App Engine API
from google.appengine.api import urlfetch

# Local modules
from common import debug, text_utils


class Markup():
    def __init__(self):
        self.formatData = {}

    def field(self, key, value):
        debug.log("Adding field " + str(key) + ": " + str(value))
        self.formatData[key] = value

    def jsonify(self):
        return self.formatData


KEYBOARD_WIDTH = 3
class Keyboard(Markup):
    def __init__(self):
        Markup.__init__(self)

        self.buttons = []
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
        Keyboard.__init__(self)

    def jsonify(self):
        self.field("keyboard", self.format())

        return self.formatData

class InlineKeyboard(Keyboard):
    def __init__(self):
        Keyboard.__init__(self)

    def jsonify(self):
        self.field("inline_keyboard", self.format())

        return self.formatData

class CloseKeyboard(Markup):
    def __init__(self):
        Keyboard.__init__(self)

    def jsonify(self):
        self.field("remove_keyboard", True)

        return self.formatData

class Post(Markup):
    def __init__(self):
        Markup.__init__(self)

        self.formatData = {
            "parse_mode": "Markdown"
        }

    def set_user(self, user):
        self.field("chat_id", text_utils.stringify(user))

    def set_reply(self, reply):
        self.field("reply_markup", reply.jsonify())

    def set_text(self, text):
        self.field("text", text)