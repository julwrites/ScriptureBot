# coding=utf-8

from common import debug, web_utils
from common.telegram import telegram_utils


def test_post():
    try:
        telegram_utils.send_msg("Testing")
    except Exception as e:
        debug.log(e)


tests = [test_post]

for test in tests:
    test()