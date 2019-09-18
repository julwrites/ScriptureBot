# coding=utf-8

import urllib
import urllib2

# Local modules
from common import debug, constants


def post_http(url, data, headers):
    debug.log("Post request to {}", [url])

    try:
        request = urllib2.Request(url, data, headers)

        response = urllib2.urlopen(request)
    except Exception as e:
        debug.err(e)
        raise


def fetch_url(url):
    hdr = {
        'User-Agent':
        'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11',
        'Accept':
        'text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8',
        'Accept-Charset': 'ISO-8859-1,utf-8;q=0.7,*;q=0.3',
        'Accept-Encoding': 'none',
        'Accept-Language': 'en-US,en;q=0.8',
        'Connection': 'keep-alive'
    }

    debug.log("Fetching url: {}", [url])

    try:
        req = urllib2.Request(url, headers=hdr)
        res = urllib2.urlopen(req)
        html = res.read()
        url = res.geturl()
    except urllib2.URLError as e:
        debug.err(e)
        raise

    return url, html

