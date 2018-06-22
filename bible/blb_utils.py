# coding=utf-8

# Python modules
import urllib
from bs4 import BeautifulSoup

# Local modules
from common import debug, html_utils, constants, text_utils, web_utils
from common.telegram import telegram_utils

from bible import blb_classes

# This URL will return search results
BLB_SEARCH_URL = "https://www.blueletterbible.org/search/preSearch.cfm?Criteria={}&t={}&ss=1"
BLB_STRONGS_URL = "https://www.blueletterbible.org/lang/Lexicon/Lexicon.cfm?strongs={}"

BLB_VERSIONS = ["NIV", "ESV", "KJV", "NASB", "RSV", "NKJV"]

# We want to grab all <div> with this id, and pick only the last bit.
BLB_VERSE_ID = "verse_{}{}"  # verse_<chapter num><3 digit verse num>
BLB_VERSE_START = "</span>"
BLB_VERS_END = " </div></div>"

REFERENCE = "reference"
VERSION = "version"
PASSAGE = "passage"


def fetch_blb(query, version="NASB"):
    debug.log("Querying for " + query)

    query = query.lower().strip()

    if query is None:
        return None

    formatRef = text_utils.strip_whitespace(query)
    formatRef = formatRef.replace(" ", "+")

    formatUrl = BLB_SEARCH_URL.format(formatRef, version)

    html = html_utils.fetch_html(formatUrl)

    if html is None:
        return None

    soup = html_utils.html_to_soup(html)

    return soup


def get_passage_raw(ref, version="NASB"):
    debug.log("Querying for passage " + ref)

    soup = fetch_blb(ref, version)
    if soup is None:
        return None

    blocks = []
    for tag in soup(class_="tools row align-middle"):
        blocks.append(tag.text)

    # # Prepare the title and header
    # reference = soup.select_one(BLB_PASSAGE_TITLE).text.strip()
    # version = version

    # # Remove the unnecessary tags
    # for tag in soup.select(BLB_PASSAGE_IGNORE):
    #     tag.decompose()

    # # Steps through all the html types and mark these
    # soup = html_utils.stripmd_soup(soup)

    # # Special formatting for chapter and verse
    # html_utils.foreach_tag(soup, ".chapternum", telegram_utils.bold)
    # html_utils.foreach_tag(soup, ".versenum", telegram_utils.to_sup)
    # html_utils.foreach_tag(soup, ".versenum", telegram_utils.italics)
    # html_utils.foreach_header(soup, telegram_utils.bold)

    # # Marking the parts of the soup we want to print
    # soup = html_utils.mark_soup(soup, BLB_PASSAGE_SELECT,
    #                             html_utils.html_common_tags())

    # # Only at the last step do we do other destructive formatting
    # soup = html_utils.strip_soup(soup)

    # blocks = []
    # for tag in soup(class_=BLB_PASSAGE_SELECT):
    #     debug.log("Joining " + tag.text)
    #     blocks.append(tag.text)

    # text = telegram_utils.join(blocks, "\n\n")

    debug.log("Finished parsing soup")

    return blb_classes.BLBPassage(reference, version, text)


def get_passage(ref, version="NASB"):
    passage = get_passage_raw(ref, version)

    if passage is None:
        return None

    text = telegram_utils.bold(passage.get_reference())
    text += " " + telegram_utils.bracket(passage.get_version())
    text += "\n\n" + passage.get_text()

    return text


def get_strongs_link(query):
    debug.log("Fetching Strongs: " + query)

    query = query.upper().strip()

    if query is None:
        return None

    formatRef = text_utils.strip_whitespace(query)

    formatUrl = BLB_STRONGS_URL.format(formatRef)

    if html_utils.fetch_html(formatUrl) is None:
        return None

    return telegram_utils.link(formatRef, formatUrl)


def get_search_raw(query):
    debug.log("Word search: " + query)

    passage = fetch_blb(ref, version)

    if passage is None:
        return None

    text = telegram_utils.bold(passage.get_reference())
    text += " " + telegram_utils.bracket(passage.get_version())
    text += "\n\n" + passage.get_text()

    return text


def get_search(query):
    soup = fetch_blb(query, version)

    if soup is None:
        return None


def get_versions():
    return BLB_VERSIONS
