# coding=utf-8

# Google App Engine Modules
from google.appengine.ext import db


def retrieve(region, name):
    key = db.Key.from_path(region, str(name))
    return db.get(key)


class BooleanProperty(db.BooleanProperty):
    def __init__(self, default=None, indexed=False):
        super(BooleanProperty, self).__init__(default=default, indexed=indexed)


class StringProperty(db.StringProperty):
    def __init__(self, default=None, indexed=False, multiline=False):
        super(StringProperty, self).__init__(
            default=default, indexed=indexed, multiline=multiline)


class StringListProperty(db.StringListProperty):
    def __init__(self, default=None, indexed=False):
        super(StringListProperty, self).__init__(
            default=default, indexed=indexed)


class DateTimeProperty(db.DateTimeProperty):
    def __init__(self, default=None, indexed=False):
        super(DateTimeProperty, self).__init__(
            default=default, indexed=indexed)


class Item(db.Model):
    def update(self):
        self.put()

    def id(self):
        return self.key()

    def name(self):
        return self.key().name()