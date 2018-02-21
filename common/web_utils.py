# coding=utf-8

# Google App Engine API
from google.appengine.api import urlfetch

# Local modules
from common import debug


def post_http(url, data, headers):
    urlfetch.fetch(
        url=url, payload=data, method=urlfetch.POST, headers=headers)


# HTML to BeautifulSoup
def fetch_url(url):
    try:
        debug.log("Attempting to fetch: " + url)
        result = urlfetch.fetch(url, deadline=constants.URL_TIMEOUT)
    except urlfetch_errors.Error as e:
        debug.log("Error fetching: " + text_utils.stringify(e))
        debug.err(e)
        return None

    return result
