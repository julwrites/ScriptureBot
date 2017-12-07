
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

# Coarse isolation of the html block we want
DG_DEVO_START = "<main role='main'>"
DG_DEVO_END = "</main>"

# Which class to isolate?
DG_DEVO_SELECT = "card__shadow"

def fetch_desiringgod(query=""):
    formatUrl = DG_URL + "/" + query

    html = html_utils.fetch_html(formatUrl, DG_DEVO_START, DG_DEVO_END)
    if html is None:
        return None

    debug.log("Html: " + html)

    soup = html_utils.html_to_soup(html)

    return soup 

def get_desiringgod_raw(query=""):
    soup = fetch_desiringgod(query)
    if soup is None:
        return None

    # Steps through all the html types and mark these
    soup = html_utils.stripmd_soup(soup)

    # Finds all links and converts to markdown
    soup = html_utils.link_soup(soup, telegram_utils.link)

    # Marks the parts of the soup that we want
    soup = html_utils.mark_soup(soup, DG_DEVO_SELECT, html_utils.html_common_tags())

    # Prettifying the stuffs
    html_utils.foreach_header(soup, telegram_utils.bold)
    html_utils.style_soup(soup, html_utils.unstrip_md, html_utils.html_p_tag())
    html_utils.style_soup(soup, telegram_utils.italics, html_utils.html_p_tag())

    blocks = []
    for tag in soup(class_=DG_DEVO_SELECT):
        blocks.append(tag.text)

    passage = telegram_utils.join(blocks, "\n\n")

    return passage

def get_desiringgod(query=""):
    passage = get_desiringgod_raw(query)

    if passage is None:
        return None

    return passage