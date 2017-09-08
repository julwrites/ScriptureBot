
# coding=utf-8

# Google App Engine Modules
from google.appengine.ext import db

class Data(db.Model):
    data = db.StringProperty(multiline=True, indexed=False)
   
    def get_uid(self):
        return self.key().name()

    def get_value(self):
        return self.data