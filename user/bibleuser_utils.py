# coding=utf-8

# Local Modules
from common import debug, text_utils, database
from user import bibleuser_classes


# Functions for manipulation of user info
def get_user(userId):
    return database.retrieve("BibleUser", text_utils.stringify(userId))


def get_uid(userIdObj):
    try:
        userId = userIdObj.get_uid()
    except AttributeError:
        userId = userIdObj

    return userId


def set_profile(userId, uname, fname, lname):
    debug.log("Setting profile of {}", [userId])

    existingUser = get_user(userId)

    uname = text_utils.stringify(uname)
    fname = text_utils.stringify(fname)
    lname = text_utils.stringify(lname)

    if existingUser:
        debug.log("User exists, updating")

        existingUser.username = uname
        existingUser.firstName = fname
        existingUser.lastName = lname
        existingUser.update_last_received()
        return existingUser
    else:
        debug.log("New user!")
        userObj = bibleuser_classes.BibleUser(
            key_name=text_utils.stringify(userId),
            username=uname,
            firstName=fname,
            lastName=lname)
        userObj.put()
        return userObj


def get_user_query():
    return bibleuser_classes.BibleUser.all()


def for_each_user(fn):
    debug.log("Running {} for each user", [fn])

    # Read user database
    query = get_user_query()
    query.filter("active =", True)

    for dbUser in query.run(batch_size=500):
        fn(get_user(get_uid(dbUser)))


def migrate(userObj):
    newUserObj = bibleuser_classes.BibleUser(
        key_name=text_utils.stringify(userObj.get_uid()))

    newUserObj.clone(userObj)

    userObj.delete()
    newUserObj.put()

    return newUserObj