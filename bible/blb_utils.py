# coding=utf-8

# Python modules
import urllib
from bs4 import BeautifulSoup

# Local modules
from common import debug, html_utils, constants, text_utils, web_utils
from common.telegram import telegram_utils

from bible import blb_classes

BLB_URL = "http://www.biblegateway.com/passage/?search={}&version={}&interface=print"

BLB_PASSAGE_CLASS = "passage-text"
BLB_PASSAGE_START = '<div class="{}">'.format(BLB_PASSAGE_CLASS)
BLB_PASSAGE_END = "<!--END .{}-->".format(BLB_PASSAGE_CLASS)
BLB_PASSAGE_SELECT = "blb-passage-text"
BLB_PASSAGE_IGNORE = ".passage-display, .footnote, .footnotes, .crossrefs, .publisher-info-bottom"
BLB_PASSAGE_TITLE = ".passage-display-bcv"

BLB_VERSIONS = ["NIV", "ESV", "KJV", "NASB", "NLT", "AMP"]

REFERENCE = "reference"
VERSION = "version"
PASSAGE = "passage"


def fetch_blb(query, version="NIV"):
    debug.log("Querying for " + query)

    query = query.lower().strip()

    if query is None:
        return None

    formatRef = urllib.quote(query)
    formatUrl = BLB_URL.format(formatRef, version)

    html = html_utils.fetch_html(formatUrl, BLB_PASSAGE_START, BLB_PASSAGE_END)

    if html is None:
        return None

    soup = html_utils.html_to_soup(html, BLB_PASSAGE_CLASS)

    return soup


def find_reference(ref):
    debug.log("Parsing reference " + ref)

    parts = ref.split(" ")
    book = text_utils.find_alpha(parts)

    # Just return it, let the caller handle it
    if book == -1:
        return ref

    parts = [
        parts[i] for i in range(len(parts))
        if (i == book) or not (parts[i].isalpha())
    ]

    debug.log("Reference parts: " + text_utils.stringify(str(parts)))

    return "".join(parts)


def get_passage_raw(ref, version="NIV"):
    debug.log("Querying for passage " + ref)

    ref = find_reference(ref)

    soup = fetch_blb(ref, version)
    if soup is None:
        return None

    # Prepare the title and header
    reference = soup.select_one(BLB_PASSAGE_TITLE).text.strip()
    version = version

    # Remove the unnecessary tags
    for tag in soup.select(BLB_PASSAGE_IGNORE):
        tag.decompose()

    # Steps through all the html types and mark these
    soup = html_utils.stripmd_soup(soup)

    # Special formatting for chapter and verse
    html_utils.foreach_tag(soup, ".chapternum", telegram_utils.bold)
    html_utils.foreach_tag(soup, ".versenum", telegram_utils.to_sup)
    html_utils.foreach_tag(soup, ".versenum", telegram_utils.italics)
    html_utils.foreach_header(soup, telegram_utils.bold)

    # Marking the parts of the soup we want to print
    soup = html_utils.mark_soup(soup, BLB_PASSAGE_SELECT,
                                html_utils.html_common_tags())

    # Only at the last step do we do other destructive formatting
    soup = html_utils.strip_soup(soup)

    blocks = []
    for tag in soup(class_=BLB_PASSAGE_SELECT):
        debug.log("Joining " + tag.text)
        blocks.append(tag.text)

    text = telegram_utils.join(blocks, "\n\n")

    debug.log("Finished parsing soup")

    return blb_classes.BLBPassage(reference, version, text)


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

    soup = fetch_blb(query)
    if soup is None:
        return None

    reference = soup.select_one(BLB_PASSAGE_TITLE).text
    reference = reference.strip()

    return reference


def get_link(query, version="NIV"):
    debug.log("Querying for link " + query)

    url = BLB_URL.format(query, version)

    html = web_utils.fetch_url(url)
    if html is None:
        return None

    return url


def get_versions():
    return BLB_VERSIONS
