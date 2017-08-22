
# Python modules
import re
from bs4 import BeautifulSoup

# Google App Engine Modules
from google.appengine.api import urlfetch, urlfetch_errors

# Local modules
from common import debug
from common import text_utils

HTML_HEADER_TAGS = ['h1', 'h2', 'h3', 'h4', 'h5', 'h6']
HTML_TEXT_TAGS = ['p']
HTML_DIV_TAGS = ['br']

# Tags
def html_all_tags():
    tags = []
    tags.extend(HTML_HEADER_TAGS)
    tags.extend(HTML_TEXT_TAGS)
    tags.extend(HTML_DIV_TAGS)

    return tags

def soupify_tags(tags):
    return ','.join(tags)

# HTML Parsing
def sub_html(html, top_tag, bottom_tag):
    start = html.find(top_tag)
    if start == -1:
        return None
    end = html.find(bottom_tag, start)
    return html[start:end]
   
def strip_md(string):
    return string.replace('*', '\*').replace('_', '\_').replace('`', '\`').replace('[', '\[')

def to_sup(text):
        sups = {u'0': u'\u2070',
                u'1': u'\xb9',
                u'2': u'\xb2',
                u'3': u'\xb3',
                u'4': u'\u2074',
                u'5': u'\u2075',
                u'6': u'\u2076',
                u'7': u'\u2077',
                u'8': u'\u2078',
                u'9': u'\u2079',
                u'-': u'\u207b'}
        return ''.join(sups.get(char, char) for char in text)

def foreach_tag(soup, tags, fn):
    for tag in soup.select(tags):
        tag.string = fn(tag.text)

def foreach_header(soup, fn):
    foreach_tag(soup, soupify_tags(HTML_HEADER_TAGS), fn)

def foreach_text(soup, fn):
    foreach_tag(soup, soupify_tags(HTML_TEXT_TAGS), fn)

def foreach_all(soup, fn):
    foreach_tag(soup, soupify_tags(html_all_tags()), fn)

def strip_soup(soup):
    debug.log('Stripping soup')

    foreach_header(soup, strip_md)

    for tag in soup.select(soupify_tags(HTML_TEXT_TAGS)):
        bad_strings = tag(text=re.compile('(\*|\_|\`|\[)'))
        for bad_string in bad_strings:
            stripped_text = strip_md(unicode(bad_string))
            bad_string.replace_with(stripped_text)

        tag.replace_with(text_utils.strip_whitespace(tag.text))

    return soup

def mark_soup(soup, html_mark, tags=[]):
    tags = soupify_tags(tags)
    debug.log('Marking tags: ' + tags)

    for tag in soup.select(tags):
        tag['class'] = html_mark

    return soup
 
    