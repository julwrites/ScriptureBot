# coding=utf-8

from google.cloud import datastore

from secret import PROJECT_ID

db = datastore.Client()


def retrieve(kind, name):
    key = db.key(kind, str(name))
    return db.get(key)


# class BooleanProperty():
#     def __init__(self, default=None, indexed=False):
#         super(BooleanProperty, self).__init__(default=default, indexed=indexed)

# class StringProperty():
#     def __init__(self, default=None, indexed=False, multiline=False):
#         super(StringProperty, self).__init__(
#             default=default, indexed=indexed, multiline=multiline)

# class StringListProperty():
#     def __init__(self, default=None, indexed=False):
#         super(StringListProperty, self).__init__(
#             default=default, indexed=indexed)

# class DateTimeProperty():
#     def __init__(self, default=None, indexed=False):
#         super(DateTimeProperty, self).__init__(
#             default=default, indexed=indexed)


class Item(datastore.Entity):
    def update(self, data):
        self.put(data)

    def id(self):
        return self.key()

    def name(self):
        return self.key().name()