# coding=utf-8

import urllib
import urllib2

# Local modules
from common import debug, constants


def post_http(url, data, headers):
    debug.log("Post request to " + url)

    debug.log("headers : " + str(headers))

    # data = urllib.urlencode(data)
    debug.log("data : " + str(data))

    request = urllib2.Request(url, data, headers)

    response = urllib2.urlopen(request)
    debug.log("response : " + str(response))


# HTML to BeautifulSoup
def fetch_url(url):
    debug.log("Fetching url: " + url)

    try:
        response = urllib2.urlopen(url)
        html = response.read()
        url = response.geturl()
    except urllib2.URLError:
        debug.err(e)
        return None

    return url, html
