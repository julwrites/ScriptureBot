# coding=utf-8

# Python modules
import re
from bs4 import BeautifulSoup

# Local modules
from common import debug, text_utils, web_utils, constants

HTML_HEADER_TAGS = ["h1", "h2", "h3", "h4", "h5", "h6"]
HTML_TEXT_TAGS = ["p"]

HTML_ITEM_TAG = "a"
HTML_LINK_TAG = "href"


# Tags
def html_common_tags():
    tags = []
    tags.extend(HTML_HEADER_TAGS)
    tags.extend(HTML_TEXT_TAGS)

    return tags


def soupify_tags(tags):
    return ",".join(tags)


def html_p_tag():
    return ",".join(HTML_TEXT_TAGS)


def get_url(html):
    return html.URL


def extract_html(html, top=None, bottom=None):
    if top is None or bottom is None:
        return html

    start = html.find(top)
    if start == -1:
        return html

    end = html.find(bottom, start)
    return html[start:end]


def fetch_html(url, start=None, end=None):
    debug.log("Fetching html from: {}", [url])

    url, html = web_utils.fetch_url(url)

    html = extract_html(html, start, end)

    return url, html


def replace_html(html, tag, rep):
    if text_utils.is_valid(html):
        html = html.replace(tag, rep)
    return html


def html_to_soup(html, select=None):
    debug.log("Parsing html to soup")

    soup = BeautifulSoup(html, "lxml")

    if text_utils.is_valid(select):
        soup = soup.select_one(".{}".format(select))

    debug.log("Soup has been made")

    return soup


def fetch_rss(url):
    debug.log("Fetching rss: {}", [url])

    url, html = web_utils.fetch_url(url)

    debug.log("rss: {}", [html])

    return url, html


def rss_to_soup(rss, select=None):
    soup = BeautifulSoup(rss, "xml")

    if text_utils.is_valid(select):
        soup = soup.select_one(".{}".format(select))

    debug.log("Soup has been made")

    return soup


# BeautifulSoup Functionalities
def strip_md(s):
    return s.replace("*", "\*").replace("_", "\_").replace("`", "\`").replace(
        "[", "\[")


def unstrip_md(s):
    return s.replace("\*", "*").replace("\_", "_").replace("\`", "`").replace(
        "\[", "[")


def foreach_tag(soup, tags, fn):
    for tag in soup.select(tags):
        tag.string = fn(tag.text)


def forall(soup, tag, fn):
    for tag in soup.find_all(tag):
        tag.string = fn(tag.text)


def foreach_header(soup, fn):
    foreach_tag(soup, soupify_tags(HTML_HEADER_TAGS), fn)


def foreach_text(soup, fn):
    foreach_tag(soup, soupify_tags(HTML_TEXT_TAGS), fn)


def foreach_br(soup, fn):
    foreach_tag(soup, "br", fn)


def foreach_all(soup, fn):
    foreach_tag(soup, soupify_tags(html_common_tags()), fn)


def soup_tags(soup):
    return "|".join([tag.name for tag in soup.find_all(True)])


def strip_soup(soup):
    debug.log("Stripping soup: ")

    foreach_all(soup, text_utils.strip_whitespace)

    return soup


def stripmd_soup(soup):
    debug.log("Stripping soup markdown: ")

    foreach_header(soup, strip_md)

    for tag in soup.select(soupify_tags(HTML_TEXT_TAGS)):
        badStrings = tag(text=re.compile("(\*|\_|\`|\[)"))
        for badString in badStrings:
            strippedText = strip_md(text_utils.stringify(badString))
            badString.replace_with(strippedText)

    return soup


def mark_soup(soup, mark, tags=[]):
    tags = soupify_tags(tags)
    debug.log("Marking tags: {}", [tags])

    for tag in soup.select(tags):
        # debug.log("Marking {}", [tag.text])
        tag["class"] = mark

    return soup


def link_soup(soup, fn):
    for tag in soup.find_all(HTML_ITEM_TAG, href=True):
        # debug.log("Converting link: {}", [tag.text])
        tag.string = fn(tag.text, tag["href"])

    return soup


def style_soup(soup, fn, find=True):
    for tag in soup.find_all(find, style=True):
        # debug.log("Styling tag: {}", [tag.text])
        tag.string = fn(tag.text)