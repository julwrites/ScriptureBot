
# Google App Engine Modules
from google.appengine.ext import db

# Local Modules
from common import debug, text_utils
from common.user import bibleuser_class


# Database util functions
def get_key(path, user_id):
    return db.Key.from_path(path, str(user_id))

# Functions for manipulation of user info
def get_user(user_id):
    user_obj = db.get(get_key('BibleUser', user_id))
    return user_obj

def get_uid(user_id_obj):
    try:
        user_id = user_id_obj.get_uid()
    except AttributeError:
        user_id = user_id_obj

    return user_id

def set_profile(user_id, uname, fname, lname):
    existing_user = get_user(user_id)

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
        user_obj = bibleuser_class.BibleUser(key_name=str(user_id), username=uname, first_name=fname, last_name=lname)
        user_obj.put()
        return user_obj

def get_user_query():
    return bibleuser_class.BibleUser.all()

def for_each_user(fn):
    debug.log('Running ' + str(fn) + ' for each user')
    
    # Read user database
    query = get_user_query()
    query.filter('active =', True)

    try:
        for db_user in query.run(batch_size=500):
            fn(get_user(get_uid(db_user)))
    except Exception as e:
        debug.log(str(e))

