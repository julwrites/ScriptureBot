# coding=utf-8

import firebase_admin
from firebase_admin import credentials
from firebase_admin import firestore

# Use the application default credentials
cred = credentials.ApplicationDefault()
firebase_admin.initialize_app(cred, {
    'projectId': project_id,
})

db = firestore.client()


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