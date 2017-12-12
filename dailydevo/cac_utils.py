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
CAC_URL = "https://cac.org/category/daily-meditations/"

# Coarse isolation of the html block we want
CAC_DEVO_START = "<hr>"
CAC_DEVO_END = "<p><strong>References:"

# Which class to isolate?
CAC_DEVO_SELECT = "cac-devo-text"
# Which tags to ignore?
CAC_DEVO_IGNORE = "h2"


def fetch_cac(version="NIV"):
    formatUrl = CAC_URL

    html = html_utils.fetch_html(formatUrl, CAC_DEVO_START, CAC_DEVO_END)
    if html is None:
        return None

    # debug.log("Html: " + html)

    soup = html_utils.html_to_soup(html)

    return soup


def get_cacdevo_raw(version="NIV"):
    soup = fetch_cac(version)
    if soup is None:
        return None

    for tag in soup.select(CAC_DEVO_IGNORE):
        tag.decompose()

    # Steps through all the html types and mark these
    soup = html_utils.stripmd_soup(soup)

    # Finds all links and converts to markdown
    soup = html_utils.link_soup(soup, telegram_utils.link)

    # Marks the parts of the soup that we want
    soup = html_utils.mark_soup(soup, CAC_DEVO_SELECT,
                                html_utils.html_common_tags())

    # Prettifying the stuffs
    html_utils.foreach_header(soup, telegram_utils.bold)
    html_utils.style_soup(soup, html_utils.unstrip_md, html_utils.html_p_tag())
    html_utils.style_soup(soup, telegram_utils.italics,
                          html_utils.html_p_tag())

    blocks = []
    for tag in soup(class_=CAC_DEVO_SELECT):
        blocks.append(tag.text)

    return blocks


def get_cacdevo(version="NIV"):
    blocks = get_cacdevo_raw(version)

    if blocks is None:
        return None

    passage = telegram_utils.join(blocks, "\n\n")

    return passage