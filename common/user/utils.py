
# Local Modules
from common.user import bibleuser_utils


# Functions for manipulation of user info
def get_user(uid):
    return bibleuser_utils.get_user(uid)

def get_uid(user):
    return bibleuser_utils.get_uid(user)

def set_profile(uid, uname, fname, lname):
    return bibleuser_utils.set_profile(uid, uname, fname, lname)

def get_user_query():
    return bibleuser_utils.get_user_query()

def for_each_user(fn):
    return bibleuser_utils.for_each_user(fn)