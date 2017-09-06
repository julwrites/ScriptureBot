
# Local Modules
from user import bibleuser_utils


# Functions for manipulation of user info
def get_user(userId):
    return bibleuser_utils.get_user(userId)

def get_uid(userIdObj):
    return bibleuser_utils.get_uid(userIdObj)

def set_profile(userId, uname, fname, lname):
    return bibleuser_utils.set_profile(userId, uname, fname, lname)

def get_user_query():
    return bibleuser_utils.get_user_query()

def for_each_user(fn):
    return bibleuser_utils.for_each_user(fn)