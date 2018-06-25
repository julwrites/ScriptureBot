# coding=utf-8

# Google App Engine API
from google.appengine.api import urlfetch, urlfetch_errors

import urllib2

# Local modules
from common import debug, constants


def post_http(url, data, headers):
    urlfetch.fetch(
        url=url, payload=data, method=urlfetch.POST, headers=headers)


# HTML to BeautifulSoup
def fetch_url(url):
    try:
        result = urllib2.urlopen(url)
        result = result.read()
    except urllib2.URLError:
        debug.log("Error fetching: " + text_utils.stringify(e))
        debug.err(e)
        return None

    return result
