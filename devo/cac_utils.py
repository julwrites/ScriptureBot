
# Python modules
import urllib
from bs4 import BeautifulSoup

# Google App Engine Modules
from google.appengine.api import urlfetch, urlfetch_errors

# Local modules
from common import debug, html_utils, constants
from common.telegram import telegram_utils


CAC_URL = 'https://cac.org/category/daily-meditations/'

CAC_DEVO_CLASS = 'devo-text'
CAC_DEVO_START = '<div class="{}">'.format(CAC_DEVO_CLASS)
CAC_DEVO_END = '<!--END .{}-->'.format(CAC_DEVO_CLASS)
CAC_DEVO_SELECT = 'bgw-devo-text'
CAC_DEVO_IGNORE = '.devo-display, .footnote, .footnotes, .crossrefs, .publisher-info-bottom'
CAC_DEVO_TITLE = '.devo-display-bcv'

REFERENCE = 'reference'
VERSION = 'version'
DEVO = 'devo'

def extract_devo(html):
    return html_utils.sub_html(html, CAC_DEVO_START, CAC_DEVO_END)

def fetch_cac(version='NIV'):
    formatUrl = CAC_URL

    soup = html_utils.fetch_html(formatUrl, CAC_DEVO_START, CAC_DEVO_END, CAC_DEVO_CLASS)

    return soup 

def get_devo_raw(version='NIV'):
    devoSoup = fetch_cac(version)
    if devoSoup is None:
        return None

    # Remove the unnecessary tags
    for tag in devoSoup.select(CAC_DEVO_IGNORE):
        tag.decompose()

    # Steps through all the html types and mark these
    soup = html_utils.stripmd_soup(devoSoup)
    soup = html_utils.mark_soup(devoSoup, 
    CAC_DEVO_SELECT,
    html_utils.HTML_HEADER_TAGS + html_utils.HTML_TEXT_TAGS)

    html_utils.foreach_header(devoSoup, telegram_utils.bold)

    # Special formatting for chapter and verse
    for tag in soup.select('.chapternum'):
        tag.string = telegram_utils.bold(tag.text)
    for tag in soup.select('.versenum'):
        tag.string = telegram_utils.italics(telegram_utils.to_sup(tag.text))

    # Only at the last step do we do other destructive formatting
    soup = html_utils.strip_soup(soup=devoSoup)

    devoBlocks = []
    for tag in soup(class_=CAC_DEVO_SELECT):
        devoBlocks.append(tag.text)

    devoText = telegram_utils.join(devoBlocks, '\n\n')

    debug.log("Finished parsing soup")

    return devoText

def get_devo(version='NIV'):
    devo = get_devo_raw(version)

    if devo is None:
        return None

    return devo

