# coding=utf-8

# Python modules
import urllib
from bs4 import BeautifulSoup

# Google App Engine Modules
from google.appengine.api import urlfetch, urlfetch_errors

# Local modules
from common import debug, html_utils, text_utils, constants
from common.telegram import telegram_utils

from bible import bible_utils

# Link to fetch html from
MCHEYNE_URL = "http://www.edginet.org/mcheyne/rss_feed.php?type=rss_2.0&tz=8&cal=classic&bible=esv&conf=no"

# Which class to isolate?
MCHEYNE_SELECT = "title"


def fetch_mcheyne():
    formatUrl = MCHEYNE_URL

    rss = html_utils.fetch_rss(formatUrl)
    if rss is None:
        return None

    # debug.log("RSS: " + rss)

    soup = html_utils.rss_to_soup(rss)

    return soup


def get_mcheyne_raw():
    soup = fetch_mcheyne()
    if soup is None:
        return None

    # Steps through all the html types and mark these
    blocks = []
    for tag in soup.findAll(MCHEYNE_SELECT):
        ref = text_utils.strip_block(tag.text, "(", ")")
        if bible_utils.fetch_passage_html(ref) is not None:
            blocks.append(ref)

    return blocks


def get_mcheyne():
    blocks = get_mcheyne_raw()

    if blocks is None:
        return None

    return blocks
