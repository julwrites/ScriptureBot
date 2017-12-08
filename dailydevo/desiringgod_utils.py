
# coding=utf-8

# Python modules
import urllib
from bs4 import BeautifulSoup

# Google App Engine Modules
from google.appengine.api import urlfetch, urlfetch_errors

# Local modules
from common import debug, html_utils, text_utils, constants
from common.telegram import telegram_utils

# Link to fetch html from
DG_URL = "https://www.desiringgod.org/articles"

# Coarse isolation of the html block we want
DG_START = "<main role='main'>"
DG_END = "</main>"


# Which class to isolate?
DG_SELECT = "share share--card js-share-values"

def fetch_desiringgod():
    formatUrl = DG_URL

    html = html_utils.fetch_html(formatUrl, DG_START, DG_END)
    if html is None:
        return None

    # debug.log("Html: " + html)

    soup = html_utils.html_to_soup(html)

    return soup 

def get_desiringgod_raw():
    soup = fetch_desiringgod()
    if soup is None:
        return None

    blocks = []
    for tag in soup(class_=DG_SELECT):
        blocks.append([tag["data-title"], tag["data-link"]])

    return blocks

def get_desiringgod():
    blocks = get_desiringgod_raw()

    if blocks is None:
        return None

    return blocks
