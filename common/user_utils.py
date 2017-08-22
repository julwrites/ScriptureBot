
# Google App Engine Modules
from google.appengine.ext import db

# Local Modules
from common import chrono
from common import database
from common import text_utils
from common import bible_user


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

    uname = text_utils.stringify(uname)
    fname = text_utils.stringify(fname)
    lname = text_utils.stringify(fname)

    if existing_user:
        existing_user.username = uname
        existing_user.first_name = fname
        existing_user.last_name = lname
        existing_user.update_last_received()
        return existing_user
    else:
        user = bible_user.BibleUser(key_name=str(uid), username=uname, first_name=fname, last_name=lname)
        user.put()
        return user

def get_user_query():
    return bible_user.BibleUser.all()

