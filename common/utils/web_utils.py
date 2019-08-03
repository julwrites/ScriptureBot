# coding=utf-8

import urllib.request, urllib.parse, urllib.error
import urllib.request, urllib.error, urllib.parse

# Local modules
from common import constants
from common.utils import debug_utils


def post_http(url, data, headers):
    debug_utils.log("Post request to {}", [url])

    try:
        request = urllib.request.Request(url, data, headers)

        response = urllib.request.urlopen(request)
    except Exception as e:
        debug_utils.err(e)
        raise


def fetch_url(url):
    debug_utils.log("Fetching url: {}", [url])

    try:
        headers = {'User-Agent': user_agent}
        req = urllib.request.Request(url, None, headers)
        response = urllib.request.urlopen(req)
        html = response.read()
        url = response.geturl()
    except urllib.error.URLError as e:
        debug_utils.err(e)
        raise

    return url, html
