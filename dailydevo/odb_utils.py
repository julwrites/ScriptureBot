# coding=utf-8

# Python modules
import urllib
from bs4 import BeautifulSoup

# Google App Engine Modules
from google.appengine.api import urlfetch, urlfetch_errors

# Local Modules
from common import debug, html_utils, text_utils
from common.telegram import telegram_utils
from bible import bible_utils

# Link to fetch html from
ODB_URL = "https://odb.org"

# Coarse isolation of the html block we want
ODB_START = "<article"
ODB_END = "</article>"

# Which class to isolate?
ODB_VERSE = "verse-box"
ODB_SCRIPTURE_LINK = "scripture-link"
ODB_PASSAGE = "post-content"


def fetch_odb():
    formatUrl = ODB_URL

    html = html_utils.fetch_html(formatUrl, ODB_START, ODB_END)
    if html is None:
        return None

    # debug.log("Html: " + html)

    soup = html_utils.html_to_soup(html)

    return soup


def get_odb_raw():
    soup = fetch_odb()
    if soup is None:
        return None

    # Strips the markdown from the html
    soup = html_utils.stripmd_soup(soup)

    blocks = []
    for tag in soup(class_=ODB_VERSE):
        for link in tag(class_=ODB_SCRIPTURE_LINK):
            verse = bible_utils.get_passage(
                link.text, text_md=telegram_utils.italics)
            blocks.append(verse)

    for tag in soup(class_=ODB_PASSAGE):
        for p in tag.select(html_utils.html_p_tag()):
            blocks.append(text_utils.strip_whitespace(p.text))

    return blocks


def get_odb():
    blocks = get_odb_raw()

    if blocks is None:
        return None

    for block in blocks:
        debug.log("Block:" + block)

    passage = telegram_utils.join(blocks, "\n\n")

    return passage
