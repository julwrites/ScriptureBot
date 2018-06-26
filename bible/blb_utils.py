# coding=utf-8

# Python modules
import re
from bs4 import BeautifulSoup

# Local modules
from common import debug, html_utils, constants, text_utils, web_utils
from common.telegram import telegram_utils

from bible import blb_classes

# This URL will return search results
BLB_SEARCH_URL = "https://www.blueletterbible.org/search/preSearch.cfm?Criteria={}&t={}&ss=1"

BLB_VERSIONS = ["NIV", "ESV", "KJV", "NASB", "RSV", "NKJV"]

BLB_VERSE_CLASS = "columns tablet-8 small-10 tablet-order-3 small-order-2"


def fetch_blb(query, version="NASB"):
    debug.log("Querying for {}", [query])

    query = query.lower().strip()

    if query is None:
        return None

    formatUrl = BLB_SEARCH_URL.format(query, version)

    url, html = html_utils.fetch_html(formatUrl)

    if html is None:
        return None

    soup = html_utils.html_to_soup(html, "nocrumbs")

    if soup is None:
        return None

    return url, html, soup


def get_search_raw(soup, version="NASB"):
    debug.log("Parsing search results")

    # We will need a BLBSearchResult for this
    return None


def get_search(query, version="NASB"):
    debug.log("Word search: {}", [query])

    url, html, soup = fetch_blb(query, version)

    if soup is None:
        return None

    header = "\n".join([tag.text for tag in soup.select("h1")])

    if header == "Search Results":
        results = get_search_raw(soup)

        if results is None:
            return None

        return telegram_utils.join(results.get_results(), "\n\n")

    return None


def get_strongs_link(soup):
    debug.log("Fetching Strongs: {}", [query])

    header = "\n".join([tag.text for tag in soup.select("h1")])

    return "{}", telegram_utils.link(header, formatUrl)


def get_passage_raw(html, soup, version="NASB"):
    debug.log("Parsing passage")

    # Prepare the title and header
    reference = soup.select_one("h1").text.strip()

    blocks = []
    lexicon = []

    # Parse raw html; bs4 ignores too much, we need the raw text
    debug.log("raw html: {}", [html])

    # Break up the html into verses first
    verse_pos = []
    cache_pos = []
    beg = 0
    while True:
        beg = beg + html[beg:].find(BLB_VERSE_CLASS)
        end = beg + html[beg + 1:].find(BLB_VERSE_CLASS) + 1
        if end == -1 or end >= len(html) or beg == end:
            break
        debug.log("beg: {}, end:{}", [beg, end])
        cache_pos.append({"begin": beg, "end": end})
        beg = end

    # Filter each position into more specific chunks containing only the verse data
    verse_pos = cache_pos
    cache_pos = []
    for pos in verse_pos:
        beg = html[pos["begin"]:pos["end"]].find('class="hide-for-tablet">')
        end = html[beg + 1:pos["end"]].find("</div></div>")
        if beg != -1 and end != -1:
            cache_pos.append({"begin": beg, "end": end})

    # Split up the verses into blocks of text and links
    verse_blocks = []
    for pos in verse_pos:
        blocks = html[pos["begin"]:pos["end"]].split("sup")
        debug.log("verse_block: {}", [blocks])
        verse_blocks.append(blocks)

    text = telegram_utils.join(verse_blocks, "\n\n")

    debug.log("Finished parsing soup")

    return blb_classes.BLBPassage(reference, version, text, lexicon)


def get_strongs(query, version="NASB"):
    debug.log("Fetching Strongs: {}", [query])

    url, html, soup = fetch_blb(query, version)

    if soup is None:
        return None

    header = "\n".join([tag.text for tag in soup.select("h1")])

    if header.find("Lexicon") != -1:
        return telegram_utils.link(header, url)
    else:
        passage = get_passage_raw(html, soup, version)

        if passage is None:
            return None

        text = telegram_utils.bold(passage.get_reference())
        text += " " + telegram_utils.bracket(passage.get_version())
        text += "\n\n" + passage.get_text()

        return text, passage.get_strongs()


def get_versions():
    return BLB_VERSIONS
