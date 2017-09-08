
# -*- coding: utf-8 -*-

# Python modules
import urllib
from bs4 import BeautifulSoup

# Google App Engine Modules
from google.appengine.api import urlfetch, urlfetch_errors

# Local modules
from common import debug, html_utils, constants
from common.telegram import telegram_utils


CAC_URL = 'https://cac.org/category/daily-meditations/'

CAC_DEVO_START = '<hr>'
#'<p><!--{}--></p>'.format('[Most recent post will go here, with week title, day title, and date headings—body of post itself, no banner image or title field.]')
CAC_DEVO_END = '</div>'
CAC_DEVO_SELECT = 'cac-devo-text'
CAC_DEVO_IGNORE = 'h2'
CAC_DEVO_LINKS = 'href'
CAC_DEVO_TITLE = 'h3'

REFERENCE = 'reference'
VERSION = 'version'
DEVO = 'devo'

def fetch_cac(version='NIV'):
    formatUrl = CAC_URL

    html = html_utils.fetch_html(formatUrl, CAC_DEVO_START, CAC_DEVO_END)
    if html is None:
        return None

    soup = html_utils.html_to_soup(html)

    return soup 

def get_cacdevo_raw(version='NIV'):
    soup = fetch_cac(version)
    if soup is None:
        return None

    for tag in soup.select(CAC_DEVO_IGNORE):
        tag.decompose()

    debug.log('Decomposed soup')
    # Steps through all the html types and mark these
    soup = html_utils.stripmd_soup(soup)
    soup = html_utils.link_soup(soup, telegram_utils.link)
    soup = html_utils.mark_soup(soup, 
    CAC_DEVO_SELECT, html_utils.html_common_tags())

    html_utils.foreach_header(soup, telegram_utils.bold)

    # Only at the last step do we do other destructive formatting
    soup = html_utils.strip_soup(soup=soup)

    blocks = []
    for tag in soup(class_=CAC_DEVO_SELECT):
        blocks.append(tag.text)

    passage = telegram_utils.join(blocks, '\n\n')

    return passage

def get_cacdevo(version='NIV'):
    passage = get_cacdevo_raw(version)

    if passage is None:
        return None

    return passage