
# Python modules
import urllib
from bs4 import BeautifulSoup

# Google App Engine Modules
from google.appengine.api import urlfetch, urlfetch_errors
from google.appengine.ext import db

# Local modules
from common import debug, html_utils, constants
from common.telegram import telegram_utils


BGW_URL = 'http://www.biblegateway.com/passage/?search={}&version={}&interface=print'

BGW_PASSAGE_CLASS = 'passage-text'
BGW_PASSAGE_START = '<div class="{}">'.format(BGW_PASSAGE_CLASS)
BGW_PASSAGE_END = '<!--END .{}-->'.format(BGW_PASSAGE_CLASS)
BGW_PASSAGE_SELECT = 'bgw-passage-text'
BGW_PASSAGE_IGNORE = '.passage-display, .footnote, .footnotes, .crossrefs, .publisher-info-bottom'
BGW_PASSAGE_TITLE = '.passage-display-bcv'

REFERENCE = 'reference'
VERSION = 'version'
PASSAGE = 'passage'

class BGWPassage():
    def __init__(self, ref, ver, txt):
        self.reference = ref
        self.version = ver
        self.text = txt

    def get_reference(self):
        return self.reference

    def get_version(self):
        return self.version

    def get_text(self):
        return self.text

def extract_passage(html):
    return html_utils.sub_html(html, BGW_PASSAGE_START, BGW_PASSAGE_END)

def fetch_bgw(query, version='NIV'):
    format_ref = urllib.quote(query.lower().strip())
    format_url = BGW_URL.format(format_ref, version)

    try:
        debug.log('Attempting to fetch: ' + query + ' ' + version + ' from ' + format_url)
        result = urlfetch.fetch(format_url, deadline=constants.URL_TIMEOUT)
    except urlfetch_errors.Error as e:
        debug.log('Error fetching: ' + str(e))
        return None

    # Format using BS4 into a form we can use for extraction
    passage_html = extract_passage(result.content)
    if passage_html is None:
        return None

    soup = BeautifulSoup(passage_html, 'lxml').select_one('.{}'.format(BGW_PASSAGE_CLASS))
    
    debug.log("Soup has been made")

    return soup 

def get_passage_raw(ref, version='NIV'):
    debug.log('Querying for passage ' + ref)

    passage_soup = fetch_bgw(ref, version)
    if passage_soup is None:
        return None

    # Prepare the title and header
    passage_reference = passage_soup.select_one(BGW_PASSAGE_TITLE).text.strip()
    passage_version = version

    # Remove the unnecessary tags
    for tag in passage_soup.select(BGW_PASSAGE_IGNORE):
        tag.decompose()

    # Steps through all the html types and mark these
    soup = html_utils.stripmd_soup(soup=passage_soup)
    soup = html_utils.mark_soup(soup=passage_soup, 
    html_mark=BGW_PASSAGE_SELECT,
    tags=html_utils.HTML_HEADER_TAGS + html_utils.HTML_TEXT_TAGS)

    html_utils.foreach_header(passage_soup, telegram_utils.bold)

    # Special formatting for chapter and verse
    for tag in soup.select('.chapternum'):
        tag.string = telegram_utils.bold(tag.text)
    for tag in soup.select('.versenum'):
        tag.string = telegram_utils.italics(html_utils.to_sup(tag.text))

    # Only at the last step do we do other destructive formatting
    soup = html_utils.strip_soup(soup=passage_soup)

    passage_blocks = []
    for tag in soup(class_=BGW_PASSAGE_SELECT):
        passage_blocks.append(tag.text)

    passage_text = telegram_utils.join(passage_blocks, '\n\n')

    debug.log("Finished parsing soup")

    return BGWPassage(passage_reference, passage_version, passage_text)

def get_passage(ref, version='NIV'):
    passage = get_passage_raw(ref, version)

    if passage is None:
        return None

    passage_format = telegram_utils.bold(passage.get_reference())
    passage_format += ' ' + telegram_utils.bracket(passage.get_version())
    passage_format += '\n\n' + passage.get_text()

    return passage_format

def get_reference(query):
    debug.log('Querying for reference ' + query)

    passage_soup = fetch_bgw(query)
    if passage_soup is None:
        return None

    reference = passage_soup.select_one(BGW_PASSAGE_TITLE).text
    reference = reference.strip()

    return reference
