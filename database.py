# Google App Engine Modules
from google.appengine.ext import db

# Local Modules
import chrono

from data import Data

# Database util functions
def get_key(path, uid):
    return db.Key.from_path(path, str(uid))


# Functions for manipulation of data
def get_data(uid):
    val = db.get(get_key('Data', uid))
    return val

def has_data(uid):
    try:
        val = db.get(get_key('Data', uid))
    except db.KindError:
        val = None
    
    return val

def set_data(uid, data):
    if has_data(uid):
        val = get_data(uid)
        val.data = str(data)
    else:
        val = Data(key_name=str(uid), data=str(data))
        val.put()
    
    return val
