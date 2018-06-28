# coding=utf-8

# Local modules
from common import debug, text_utils


class Markup():
    def __init__(self):
        self.formatData = {}

    def set_field(self, key, value):
        self.formatData[key] = value

    def set_text(self, text):
        self.set_field("text", text_utils.to_utf8(text))

    def data(self):
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
                row.append(button.data())
                num -= 1

            keyboard.append(row)

        return keyboard


class ReplyKeyboard(Keyboard):
    def __init__(self):
        Keyboard.__init__(self)

    def data(self):
        self.set_field("keyboard", self.format())

        return self.formatData


class InlineKeyboard(Keyboard):
    def __init__(self):
        Keyboard.__init__(self)

    def data(self):
        self.set_field("inline_keyboard", self.format())

        return self.formatData


class CloseKeyboard(Markup):
    def __init__(self):
        Markup.__init__(self)

    def data(self):
        self.set_field("remove_keyboard", True)

        return self.formatData


class Post(Markup):
    def __init__(self):
        Markup.__init__(self)

        self.set_field("parse_mode", "Markdown")

    def set_user(self, user):
        self.set_field("chat_id", text_utils.to_utf8(user))

    def set_reply(self, reply):
        self.set_field("reply_markup", reply.data())
