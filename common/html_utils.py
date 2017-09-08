
# Python modules
import re
from bs4 import BeautifulSoup

# Google App Engine Modules
from google.appengine.api import urlfetch, urlfetch_errors

# Local modules
from common import debug, text_utils, constants


HTML_HEADER_TAGS = ['h1', 'h2', 'h3', 'h4', 'h5', 'h6']
HTML_TEXT_TAGS = ['p']

HTML_ITEM_TAG = 'a'
HTML_BREAK_TAG = 'br'


# Tags
def html_all_tags():
    tags = []
    tags.extend(HTML_HEADER_TAGS)
    tags.extend(HTML_TEXT_TAGS)

    return tags

def soupify_tags(tags):
    return ','.join(tags)


# HTML to BeautifulSoup
def fetch_url(url):
    try:
        debug.log('Attempting to fetch: ' + url)
        result = urlfetch.fetch(url, deadline=constants.URL_TIMEOUT)
    except urlfetch_errors.Error as e:
        debug.log('Error fetching: ' + str(e))
        return None

    return result

def extract_html(html, top, bottom):
    start = html.find(top)
    if start == -1:
        return None
    end = html.find(bottom, start)
    return html[start:end]

def fetch_html(url, start, end):
    result = fetch_url(url)

    # Format using BS4 into a form we can use for extraction
    return extract_html(result.content, start, end)

def html_to_soup(html, select=None):
    soup = BeautifulSoup(html, 'lxml')

    if text_utils.is_valid(select):
        soup = soup.select_one('.{}'.format(select))

    debug.log("Soup has been made")

    return soup 

def fetch_rss(url):
    result = fetch_url(url)

    return result.content

def rss_to_soup(rss, select=None):
    soup = BeautifulSoup(rss, 'xml')

    if text_utils.is_valid(select):
        soup = soup.select_one('.{}'.format(select))

    debug.log("Soup has been made")

    return soup


# BeautifulSoup Functionalities
def strip_md(string):
    return string.replace('*', '\*').replace('_', '\_').replace('`', '\`').replace('[', '\[')

def foreach_tag(soup, tags, fn):
    for tag in soup.select(tags):
        tag.string = fn(tag.text)

def foreach_header(soup, fn):
    foreach_tag(soup, soupify_tags(HTML_HEADER_TAGS), fn)

def foreach_text(soup, fn):
    foreach_tag(soup, soupify_tags(HTML_TEXT_TAGS), fn)

def foreach_break(soup, fn):
    for tag in soup.find_all(HTML_BREAK_TAG):
        tag.string = fn(tag.text)

def foreach_all(soup, fn):
    foreach_tag(soup, soupify_tags(html_all_tags()), fn)

def strip_soup(soup):
    debug.log('Stripping soup')

    foreach_all(soup, text_utils.strip_whitespace)

    return soup

def stripmd_soup(soup):
    debug.log('Stripping soup markdown')

    foreach_header(soup, strip_md)

    for tag in soup.select(soupify_tags(HTML_TEXT_TAGS)):
        badStrings = tag(text=re.compile('(\*|\_|\`|\[)'))
        for badString in badStrings:
            strippedText = strip_md(unicode(badString))
            badString.replace_with(strippedText)

    return soup

def mark_soup(soup, htmlMark, tags=[], markbr_=True):
    tags = soupify_tags(tags)
    debug.log('Marking tags: ' + tags)

    for tag in soup.select(tags):
        tag['class'] = htmlMark
    
    if markbr_:
        for tag in soup.find_all(HTML_BREAK_TAG):
            debug.log('Marking br')
            tag['class'] = htmlMark

    return soup

def link_soup(soup, fn):
    for tag in soup.find_all(HTML_ITEM_TAG, href=True):
        debug.log('Converting link: ' + tag.text)
        tag.string = fn(tag.text, tag['href'])
    
    return soup
