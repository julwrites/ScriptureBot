
# Python modules
import urllib
from bs4 import BeautifulSoup

# Google App Engine Modules
from google.appengine.api import urlfetch, urlfetch_errors

# Local modules
from common import debug, html_utils, constants
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
    subscribeBlocks = []
    for tag in soup.findAll(MCHEYNE_SELECT):
        subscribeBlocks.append(tag.text)

    debug.log("Finished parsing soup" + '*'.join(subscribeBlocks))

    return subscribeBlocks

def get_mcheyne(version="NIV"):
    subscribeRefs = get_mcheyne_raw()
    subscribeBlocks = []

    if subscribeRefs is None:
        return None

    for ref in subscribeRefs:
        passage = bible_utils.get_passage(ref, version)

        if passage is not None:
            subscribeBlocks.append(passage)

    return subscribeBlocks
