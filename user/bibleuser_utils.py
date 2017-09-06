
# Google App Engine Modules
from google.appengine.ext import db

# Local Modules
from common import debug, text_utils
from user import bibleuser_classes


# Database util functions
def get_key(path, userId):
    return db.Key.from_path(path, str(userId))

# Functions for manipulation of user info
def get_user(userId):
    userObj = db.get(get_key('BibleUser', userId))
    return userObj

def get_uid(userIdObj):
    try:
        userId = userIdObj.get_uid()
    except AttributeError:
        userId = userIdObj

    return userId

def set_profile(userId, uname, fname, lname):
    existingUser = get_user(userId)

    uname = text_utils.stringify(uname)
    fname = text_utils.stringify(fname)
    lname = text_utils.stringify(fname)

    if existingUser:
        existingUser.username = uname
        existingUser.firstName = fname
        existingUser.lastName = lname
        existingUser.update_last_received()
        return existingUser
    else:
        userObj = bibleuser_classes.BibleUser(key_name=str(userId), username=uname, firstName=fname, lastName=lname)
        userObj.put()
        return userObj

def get_user_query():
    return bibleuser_classes.BibleUser.all()

def for_each_user(fn):
    debug.log('Running ' + str(fn) + ' for each user')
    
    # Read user database
    query = get_user_query()
    query.filter('active =', True)

    try:
        for dbUser in query.run(batch_size=500):
            fn(get_user(get_uid(dbUser)))
    except Exception as e:
        debug.log(str(e))

def migrate(userObj):
    newUserObj = bibleuser_classes.BibleUser(key_name=str(userObj.get_uid()))

    newUserObj.clone(userObj)

    userObj.delete()
    newUserObj.put()

    return newUserObj