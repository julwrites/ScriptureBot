
# Google App Engine Modules
from google.appengine.ext import db

# Local modules
import database

from user import User

class BibleUser(User):
    version = db.StringProperty(indexed=False, default='NIV')
    current_verse = db.IntegerProperty(indexed=False, default=0)
    current_pack = db.StringProperty(indexed=False, default='')

    def get_version(self):
        return self.version

    def set_version(self, version):
        self.version = version
        self.put()

    def get_current_verse(self):
        return self.current_verse

    def set_current_verse(self, verse):
        self.current_verse = verse
        self.put()

    def get_current_pack(self):
        return self.current_pack

    def set_current_pack(self, pack):
        self.current_pack = pack
        self.put()


# Functions for manipulation of user info
def get_user(uid):
    user = db.get(database.get_key('BibleUser', uid))
    return user

def get_uid(user):
    try:
        uid = user.get_uid()
    except AttributeError:
        uid = user

    return uid

def set_profile(uid, uname, fname, lname):
    existing_user = get_user(uid)
    if existing_user:
        existing_user.username = uname
        existing_user.first_name = fname
        existing_user.last_name = lname
        existing_user.update_last_received()
        return existing_user
    else:
        user = BibleUser(key_name=str(uid), username=uname, first_name=fname, last_name=lname)
        user.put()
        return user

def get_user_query():
    return BibleUser.all()


        