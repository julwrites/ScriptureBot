
# Google App Engine Modules
from google.appengine.ext import db

# Local Modules
from common import debug
from common import text_utils
from common.user import bibleuser_class


# Database util functions
def get_key(path, uid):
    return db.Key.from_path(path, str(uid))

# Functions for manipulation of user info
def get_user(uid):
    user = db.get(get_key('BibleUser', uid))
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
        user = bibleuser_class.BibleUser(key_name=str(uid), username=uname, first_name=fname, last_name=lname)
        user.put()
        return user

def get_user_query():
    return bibleuser_class.BibleUser.all()

def for_each_user(fn):
    debug.log('Running ' + str(fn) + ' for each user')
    
    # Read user database
    query = get_user_query()
    query.filter('active =', True)

    try:
        for user in query.run(batch_size=500):
            fn(get_user(get_uid(user)))
    except Exception as e:
        debug.log(str(e))

