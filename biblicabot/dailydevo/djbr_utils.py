# coding=utf-8

import djbr_data
import datetime

# Local modules
from common import constants
from common.utils import debug_utils, html_utils, text_utils
from common.telegram import telegram_utils

from bible import bible_utils


def fetch_djbr():
    data = djbr_data.get()

    month_length = len(data) / 12

    # We will read the entry using the date, format: Year, Month, Day
    date = text_utils.to_string(datetime.date.today()).split("-")
    year = int(date[0])
    month = int(date[1])
    day = int(date[2])

    if day < month_length:
        passages = data[(month - 1) * month_length + (day - 1)]

        return passages

    return None


def get_djbr_raw():
    debug_utils.log("Getting DJBR")

    passages = fetch_djbr()

    if passages is None:
        return None

    # Steps through all the html types and mark these
    blocks = []
    for ref in passages:
        blocks.append(ref)

    return blocks


def get_djbr():
    blocks = get_djbr_raw()

    if blocks is None:
        return None

    return blocks
