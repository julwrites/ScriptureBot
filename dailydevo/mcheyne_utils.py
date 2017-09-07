
# Python modules
import urllib
from bs4 import BeautifulSoup

# Google App Engine Modules
from google.appengine.api import urlfetch, urlfetch_errors

# Local modules
from common import debug, html_utils, text_utils, constants
from common.telegram import telegram_utils

from bible import bible_utils


MCHEYNE_URL = "http://www.edginet.org/mcheyne/rss_feed.php?type=rss_2.0&tz=8&cal=classic&bible=esv&conf=no"

MCHEYNE_IGNORE = ""
MCHEYNE_SELECT = "title"

def get_mcheyne_raw():
    formatUrl = MCHEYNE_URL

    rss = html_utils.fetch_rss(formatUrl)
    soup = html_utils.rss_to_soup(rss)

    if soup is None:
        return None

    # Steps through all the html types and mark these
    blocks = []
    for tag in soup.findAll(MCHEYNE_SELECT):
        ref = text_utils.strip_block(tag.text, '(', ')')
        blocks.append(ref)

    return blocks 

def get_mcheyne(version="NIV"):
    refs = get_mcheyne_raw()
    blocks = []

    if refs is None:
        return None

    for ref in refs:
        passage = bible_utils.get_passage(ref, version)

        if passage is not None:
            blocks.append(passage)

    return blocks
