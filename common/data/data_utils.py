
# Local Modules
from common.data.data_classes import Data


# Database util functions
def get_key(path, user_id):
    return db.Key.from_path(path, str(user_id))


# Functions for manipulation of data
def get_data(user_id):
    val = db.get(get_key('Data', user_id))
    return val

def has_data(user_id):
    try:
        val = db.get(get_key('Data', user_id))
    except db.KindError:
        val = None
    
    return val

def set_data(user_id, data):
    if has_data(user_id):
        val = get_data(user_id)
        val.data = str(data)
    else:
        val = Data(key_name=str(user_id), data=str(data))
        val.put()
    
    return val
