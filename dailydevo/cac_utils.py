
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
CAC_DEVO_START = "<!--[Most recent post"
CAC_DEVO_END = '</div>'
CAC_DEVO_SELECT = 'bgw-devo-text'
CAC_DEVO_IGNORE = ''
CAC_DEVO_TITLE = 'h3'

REFERENCE = 'reference'
VERSION = 'version'
DEVO = 'devo'

def fetch_cac(version='NIV'):
    formatUrl = CAC_URL

    html = html_utils.fetch_html(formatUrl, CAC_DEVO_START, CAC_DEVO_END)
    if html is None:
        return None

    soup = html_utils.html_to_soup(html, CAC_DEVO_CLASS)

    return soup 

def get_cacdevo_raw(version='NIV'):
    soup = fetch_cac(version)
    if soup is None:
        return None

    # Remove the unnecessary tags
    for tag in soup.select(CAC_DEVO_IGNORE):
        tag.decompose()

    # Steps through all the html types and mark these
    soup = html_utils.stripmd_soup(soup)
    soup = html_utils.mark_soup(soup, 
    CAC_DEVO_SELECT,
    html_utils.HTML_HEADER_TAGS + html_utils.HTML_TEXT_TAGS)

    html_utils.foreach_header(soup, telegram_utils.bold)

    # Only at the last step do we do other destructive formatting
    soup = html_utils.strip_soup(soup=soup)

    blocks = []
    for tag in soup(class_=CAC_DEVO_SELECT):
        blocks.append(tag.text)

    passage = telegram_utils.join(blocks, '\n\n')

    debug.log("Finished parsing soup")

    return passage

def get_cacdevo(version='NIV'):
    devo = get_cacdevo_raw(version)

    if devo is None:
        return None

    return devo