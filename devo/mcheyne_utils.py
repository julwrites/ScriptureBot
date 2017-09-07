
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

MCHEYNE_DEVO_START = '<rss version="2.0">'
MCHEYNE_DEVO_END = '</rss>'
MCHEYNE_IGNORE = ""
MCHEYNE_SELECT = "title"

def get_mcheyne_raw():
    formatUrl = MCHEYNE_URL

    devoSoup = html_utils.fetch_html(formatUrl, MCHEYNE_DEVO_START, MCHEYNE_DEVO_END)
    if devoSoup is None:
        return None

    # Steps through all the html types and mark these
    devoBlocks = []
    for tag in devoSoup(class_=MCHEYNE_SELECT):
        devoBlocks.append(tag.text)

    debug.log("Finished parsing soup")

    return devoBlocks

def get_mcheyne(version="NIV"):
    devoRefs = get_mcheyne_raw()
    devoBlocks = []

    if devoRefs is not None:
        for ref in devoRefs:
            passage = bible_utils.get_passage(ref, version)

            if passage is not None:
                devoBlocks.append(passage)

        return devoBlocks
    return None
