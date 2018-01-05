# coding=utf-8

# Python modules
import urllib
from bs4 import BeautifulSoup

# Google App Engine Modules
from google.appengine.api import urlfetch, urlfetch_errors
from google.appengine.ext import db

# Local modules
from common import debug, html_utils, constants
from common.telegram import telegram_utils

from bible import bgw_classes

BGW_URL = "http://www.biblegateway.com/passage/?search={}&version={}&interface=print"

BGW_PASSAGE_CLASS = "passage-text"
BGW_PASSAGE_START = '<div class="{}">'.format(BGW_PASSAGE_CLASS)
BGW_PASSAGE_END = "<!--END .{}-->".format(BGW_PASSAGE_CLASS)
BGW_PASSAGE_SELECT = "bgw-passage-text"
BGW_PASSAGE_IGNORE = ".passage-display, .footnote, .footnotes, .crossrefs, .publisher-info-bottom"
BGW_PASSAGE_TITLE = ".passage-display-bcv"

BGW_VERSIONS = ["NIV", "ESV", "KJV", "NASB", "NLT", "AMP"]

REFERENCE = "reference"
VERSION = "version"
PASSAGE = "passage"


def fetch_bgw(query, version="NIV"):
    formatQuery = query.lower().strip()
    debug.log("Formatted query " + formatQuery)
    formatRef = urllib.quote(formatQuery)
    debug.log("Formatted ref " + formatRef)
    formatUrl = BGW_URL.format(formatRef, version)

    html = html_utils.fetch_html(formatUrl, BGW_PASSAGE_START, BGW_PASSAGE_END)

    if html is None:
        return None

    soup = html_utils.html_to_soup(html, BGW_PASSAGE_CLASS)

    return soup


def get_passage_raw(ref, version="NIV"):
    debug.log("Querying for passage " + ref)

    soup = fetch_bgw(ref, version)
    if soup is None:
        return None

    # Prepare the title and header
    reference = soup.select_one(BGW_PASSAGE_TITLE).text.strip()
    version = version

    # Remove the unnecessary tags
    for tag in soup.select(BGW_PASSAGE_IGNORE):
        tag.decompose()

    # Steps through all the html types and mark these
    soup = html_utils.stripmd_soup(soup)

    # Special formatting for chapter and verse
    html_utils.foreach_tag(soup, ".chapternum", telegram_utils.bold)
    html_utils.foreach_tag(soup, ".versenum", telegram_utils.to_sup)
    html_utils.foreach_tag(soup, ".versenum", telegram_utils.italics)
    html_utils.foreach_header(soup, telegram_utils.bold)

    # Marking the parts of the soup we want to print
    soup = html_utils.mark_soup(soup, BGW_PASSAGE_SELECT,
                                html_utils.html_common_tags())

    # Only at the last step do we do other destructive formatting
    soup = html_utils.strip_soup(soup)

    blocks = []
    for tag in soup(class_=BGW_PASSAGE_SELECT):
        debug.log("Joining " + tag.text)
        blocks.append(tag.text)

    text = telegram_utils.join(blocks, "\n\n")

    debug.log("Finished parsing soup")

    return bgw_classes.BGWPassage(reference, version, text)


def get_passage(ref, version="NIV"):
    passage = get_passage_raw(ref, version)

    if passage is None:
        return None

    text = telegram_utils.bold(passage.get_reference())
    text += " " + telegram_utils.bracket(passage.get_version())
    text += "\n\n" + passage.get_text()

    return text


def get_reference(query):
    debug.log("Querying for reference " + query)

    soup = fetch_bgw(query)
    if soup is None:
        return None

    reference = soup.select_one(BGW_PASSAGE_TITLE).text
    reference = reference.strip()

    return reference


def get_link(query, version="NIV"):
    debug.log("Querying for link " + query)

    url = BGW_URL.format(query, version)

    html = html_utils.fetch_url(url)
    if html is None:
        return None

    return url


def get_versions():
    return BGW_VERSIONS
