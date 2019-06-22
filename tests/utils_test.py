# coding=utf-8

import json
import unittest

from common.utils import debug_utils, web_utils, text_utils, html_utils


class test_text_methods(unittest.TestCase):
    def test_to_utf8(self):
        debug_utils.log("===\tTesting text_utils.to_utf8\t===")
        test_string = "utf-8: 👍🏻\nObject:{}\n:: Pass"
        norm_string = test_string.format(["ListObject"])

        utf8_string = norm_string.encode("utf-8")

        try:
            debug_utils.log(text_utils.to_utf8(norm_string))
            debug_utils.log(text_utils.to_utf8(utf8_string))
        except Exception as e:
            self.assertTrue(False, e)


class test_html_methods(unittest.TestCase):
    def test_get_url(self):
        debug_utils.log("===\tTesting html_utils.get_url\t===")
        try:
            url, html = html_utils.fetch_html("https://julwrites.github.io")
        except Exception as e:
            self.assertTrue(False, e)


class test_web_methods(unittest.TestCase):
    def test_fetch_url(self):
        debug_utils.log("===\tTesting web_utils.fetch_url\t===")

        url = "http://www.tehj.org"

        try:
            url, html = web_utils.fetch_url(url)

            debug_utils.log("Url: {}\nHtml: {}", [url, html])
        except Exception as e:
            self.assertTrue(False, e)


unittest.main()