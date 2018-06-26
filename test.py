# coding=utf-8

import json
import logging
import unittest

from common import debug, web_utils, text_utils, html_utils
from common.telegram import telegram_utils

import secret

logging.basicConfig(level=logging.DEBUG)


class test_text_methods(unittest.TestCase):
    def test_stringify(self):
        debug.log("===\tTesting text_utils.stringify\t===")
        test_string = u"utf-8: 👍🏻\nObject:{}\n:: Pass"
        norm_string = test_string.format(["ListObject"])

        utf8_string = norm_string.encode("utf-8")

        try:
            debug.log(text_utils.stringify(norm_string))
            debug.log(text_utils.stringify(utf8_string))
        except Exception as e:
            self.assertTrue(False, e)


class test_html_methods(unittest.TestCase):
    def test_get_url(self):
        debug.log("===\tTesting html_utils.get_url\t===")
        try:
            url, html = html_utils.fetch_html("https://julwrites.github.io")
        except Exception as e:
            self.assertTrue(False, e)


class test_web_methods(unittest.TestCase):
    def test_http_post(self):
        debug.log("===\tTesting web_utils.http_post\t===")

        TELEGRAM_URL = "https://api.telegram.org/bot" + secret.BOT_ID
        TELEGRAM_URL_SEND = TELEGRAM_URL + "/sendMessage"
        JSON_HEADER = {"Content-Type": "application/json;charset=utf-8"}

        data = {
            "text": "Test <http post>\n:: Pass",
            "chat_id": text_utils.stringify(secret.BOT_ADMIN),
            "parse_mode": "Markdown"
        }
        data = json.dumps(data)
        debug.log("Performing send: {}", [data])

        try:
            web_utils.post_http(TELEGRAM_URL_SEND, data, JSON_HEADER)
        except Exception as e:
            self.assertTrue(False, e)

    def test_fetch_url(self):
        debug.log("===\tTesting web_utils.fetch_url\t===")

        url = "http://www.tehj.org"

        try:
            url, html = web_utils.fetch_url(url)

            debug.log("Url: {}\nHtml: {}", [url, html])
        except Exception as e:
            self.assertTrue(False, e)


class test_telegram_methods(unittest.TestCase):
    def test_send_msg(self):
        debug.log("===\tTesting telegram_utils.send_msg\t===")

        msg = u"Test <send msg> with:\nutf-8: 👍🏻\nMarkdown: *bold*,_italic_,```monospace```\nArguments:{}\n:: Pass"
        args = [u"[tehj](https://tehj.org)"]

        debug.log("Validating Message: {}", [msg.format(args[0])])

        try:
            telegram_utils.send_msg(user=secret.BOT_ADMIN, text=msg, args=args)
        except Exception as e:
            self.assertTrue(False, e)


unittest.main()