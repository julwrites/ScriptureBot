# coding=utf-8

# Google App Engine API
from google.appengine.api import urlfetch

# Local modules
from common import debug


def post(url, data, headers):
    urlfetch.fetch(
        url=url, payload=data, method=urlfetch.POST, headers=headers)
