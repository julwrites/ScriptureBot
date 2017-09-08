
# coding=utf-8

# Local Modules
from common import text_utils
from common.data.data_classes import Data


# Database util functions
def get_key(path, userId):
    return db.Key.from_path(path, text_utils.stringify(userId))


# Functions for manipulation of data
def get_data(userId):
    val = db.get(get_key('Data', userId))
    return val

def has_data(userId):
    try:
        val = db.get(get_key('Data', userId))
    except db.KindError:
        val = None
    
    return val

def set_data(userId, data):
    if has_data(userId):
        val = get_data(userId)
        val.data = text_utils.stringify(data)
    else:
        val = Data(key_name=text_utils.stringify(userId), data=text_utils.stringify(data))
        val.put()
    
    return val
