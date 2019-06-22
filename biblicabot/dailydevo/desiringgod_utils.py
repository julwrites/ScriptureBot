# coding=utf-8

# Python modules
import urllib
from bs4 import BeautifulSoup
# Local modules
from common.utils import debug_utils, html_utils, text_utils
from common.telegram import telegram_utils

# Link to fetch html from
DG_URL = "https://www.desiringgod.org/articles"

# Coarse isolation of the html block we want
DG_START = "<main role='main'>"
DG_END = "</main>"

# Which class to isolate?
DG_TODAY = "section-1"
DG_SELECT = "share share--card js-share-values"


def fetch_desiringgod():
    formatUrl = DG_URL

    url, html = html_utils.fetch_html(formatUrl, DG_START, DG_END)
    if html is None:
        return None

    # debug_utils.log("Html: {}", [html])

    soup = html_utils.html_to_soup(html)

    return soup


def get_desiringgod_raw():
    soup = fetch_desiringgod()
    if soup is None:
        return None

    soup(id=DG_TODAY)

    blocks = []
    for tag in soup(class_=DG_SELECT):
        blocks.append({"title": tag["data-title"], "link": tag["data-link"]})

    return blocks[:3]


def get_desiringgod():
    blocks = get_desiringgod_raw()

    if blocks is None:
        return None

    return blocks
