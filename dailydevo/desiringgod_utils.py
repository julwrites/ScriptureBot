
# coding=utf-8

# Python modules
import urllib
from bs4 import BeautifulSoup

# Google App Engine Modules
from google.appengine.api import urlfetch, urlfetch_errors

# Local modules
from common import debug, html_utils, constants
from common.telegram import telegram_utils

# Link to fetch html from
DG_URL = "https://www.desiringgod.org/articles"

# Which class to isolate?
DG_SELECT = "share share--card js-share-values"

def fetch_desiringgod(query=""):
    formatUrl = DG_URL + "/" + query

    html = html_utils.fetch_html(formatUrl)
    if html is None:
        return None

    debug.log("Html: " + html)

    soup = html_utils.html_to_soup(html.main)

    return soup 

def get_desiringgod_raw(query=""):
    soup = fetch_desiringgod(query)
    if soup is None:
        return None

    blocks = []
    for tag in soup.select(DG_SELECT):
        blocks.append({"text":tag["data-title"], "url":tag["data-link"]})

    return blocks

def get_desiringgod(query=""):
    passage = get_desiringgod_raw(query)

    if passage is None:
        return None

    return passage