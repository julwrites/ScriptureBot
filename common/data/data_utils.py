
# Local Modules
from common.data.data_classes import Data


# Database util functions
def get_key(path, userId):
    return db.Key.from_path(path, str(userId))


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
        val.data = str(data)
    else:
        val = Data(key_name=str(userId), data=str(data))
        val.put()
    
    return val
