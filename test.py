# coding=utf-8

import json
import logging
import unittest

from common import debug, web_utils, text_utils, html_utils
from common.telegram import telegram_utils

import secret

logging.basicConfig(level=logging.DEBUG)


class test_web_methods(unittest.TestCase):
    def test_post(self):
        debug.log("Testing web_utils.http_post")

        TELEGRAM_URL = "https://api.telegram.org/bot" + secret.BOT_ID
        TELEGRAM_URL_SEND = TELEGRAM_URL + "/sendMessage"
        JSON_HEADER = {"Content-Type": "application/json;charset=utf-8"}

        data = {
            "text": "Testing",
            "chat_id": text_utils.stringify(secret.BOT_ADMIN),
            "parse_mode": "Markdown"
        }
        data = json.dumps(data)
        debug.log("Performing send: " + text_utils.stringify(str(data)))

        try:
            web_utils.post_http(TELEGRAM_URL_SEND, data, JSON_HEADER)
        except Exception as e:
            self.assertTrue(False, str(e))


class test_telegram_methods(unittest.TestCase):
    def test_send_msg(self):
        debug.log("Testing telegram_utils.send_msg")
        try:
            telegram_utils.send_msg(
                text="Testing", user=text_utils.stringify(secret.BOT_ADMIN))
        except Exception as e:
            self.assertTrue(False, str(e))


class test_html_methods(unittest.TestCase):
    def test_get_url(self):
        debug.log("Testing html_utils.get_url")
        try:
            url, html = html_utils.fetch_html("https://julwrites.github.io")
        except Exception as e:
            self.assertTrue(False, str(e))


unittest.main()