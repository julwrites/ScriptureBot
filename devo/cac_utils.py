
# Python modules
import urllib
from bs4 import BeautifulSoup

# Google App Engine Modules
from google.appengine.api import urlfetch, urlfetch_errors

# Local modules
from common import debug, html_utils, constants
from common.telegram import telegram_utils


CAC_URL = 'https://cac.org/category/daily-meditations/'

CAC_subscribe_CLASS = 'subscribe-text'
CAC_subscribe_START = '<div class="{}">'.format(CAC_subscribe_CLASS)
CAC_subscribe_END = '<!--END .{}-->'.format(CAC_subscribe_CLASS)
CAC_subscribe_SELECT = 'bgw-subscribe-text'
CAC_subscribe_IGNORE = '.subscribe-display, .footnote, .footnotes, .crossrefs, .publisher-info-bottom'
CAC_subscribe_TITLE = '.subscribe-display-bcv'

REFERENCE = 'reference'
VERSION = 'version'
subscribe = 'subscribe'

def fetch_cac(version='NIV'):
    formatUrl = CAC_URL

    html = html_utils.fetch_html(formatUrl, CAC_subscribe_START, CAC_subscribe_END)
    soup = html_utils.html_to_soup(html, CAC_subscribe_CLASS)

    return soup 

def get_subscribe_raw(version='NIV'):
    subscribeSoup = fetch_cac(version)
    if subscribeSoup is None:
        return None

    # Remove the unnecessary tags
    for tag in subscribeSoup.select(CAC_subscribe_IGNORE):
        tag.decompose()

    # Steps through all the html types and mark these
    soup = html_utils.stripmd_soup(subscribeSoup)
    soup = html_utils.mark_soup(subscribeSoup, 
    CAC_subscribe_SELECT,
    html_utils.HTML_HEADER_TAGS + html_utils.HTML_TEXT_TAGS)

    html_utils.foreach_header(subscribeSoup, telegram_utils.bold)

    # Special formatting for chapter and verse
    for tag in soup.select('.chapternum'):
        tag.string = telegram_utils.bold(tag.text)
    for tag in soup.select('.versenum'):
        tag.string = telegram_utils.italics(telegram_utils.to_sup(tag.text))

    # Only at the last step do we do other destructive formatting
    soup = html_utils.strip_soup(soup=subscribeSoup)

    subscribeBlocks = []
    for tag in soup(class_=CAC_subscribe_SELECT):
        subscribeBlocks.append(tag.text)

    subscribeText = telegram_utils.join(subscribeBlocks, '\n\n')

    debug.log("Finished parsing soup")

    return subscribeText

def get_subscribe(version='NIV'):
    subscribe = get_subscribe_raw(version)

    if subscribe is None:
        return None

    return subscribe

