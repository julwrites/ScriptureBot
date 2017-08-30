
# Local Modules
from common.user import bibleuser_utils


# Functions for manipulation of user info
def get_user(user_id):
    return bibleuser_utils.get_user(user_id)

def get_uid(user_id_obj):
    return bibleuser_utils.get_uid(user_id_obj)

def set_profile(user_id, uname, fname, lname):
    return bibleuser_utils.set_profile(user_id, uname, fname, lname)

def get_user_query():
    return bibleuser_utils.get_user_query()

def for_each_user(fn):
    return bibleuser_utils.for_each_user(fn)