# coding=utf-8

import urllib
import urllib2

# Local modules
from common.utils import debug_utils
from common import constants


def post_http(url, data, headers):
    debug_utils.log("Post request to {}", [url])

    try:
        request = urllib2.Request(url, data, headers)

        response = urllib2.urlopen(request)
    except Exception as e:
        debug_utils.err(e)
        raise


def fetch_url(url):
    debug_utils.log("Fetching url: {}", [url])

    try:
        response = urllib2.urlopen(url)
        html = response.read()
        url = response.geturl()
    except urllib2.URLError as e:
        debug_utils.err(e)
        raise

    return url, html
