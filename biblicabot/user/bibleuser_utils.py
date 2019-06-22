# coding=utf-8

# Local Modules
from common import database
from common.utils import debug_utils, text_utils
from user import bibleuser_classes


# Functions for manipulation of user info
def get_user(userId):
    return database.retrieve("BibleUser", text_utils.to_string(userId))


def get_uid(userIdObj):
    try:
        userId = userIdObj.get_uid()
    except AttributeError:
        userId = userIdObj

    return userId


def set_profile(userId, uname, fname, lname):
    debug_utils.log("Setting profile of {}", [userId])

    existingUser = get_user(userId)

    uname = text_utils.to_string(uname)
    fname = text_utils.to_string(fname)
    lname = text_utils.to_string(lname)

    if existingUser:
        debug_utils.log("Updating names... {} {} {}", [uname, fname, lname])
        existingUser.username = uname
        existingUser.firstName = fname
        existingUser.lastName = lname

        debug_utils.log("Updating the user... {} {} {}", [
            existingUser.username, existingUser.firstName,
            existingUser.lastName
        ])
        existingUser.update_last_received()

        return existingUser
    else:
        debug_utils.log("New user: {} {} {}", [uname, fname, lname])
        userObj = bibleuser_classes.BibleUser(
            key_name=text_utils.to_string(userId),
            username=uname,
            firstName=fname,
            lastName=lname)
        userObj.update()
        return userObj


def get_user_query():
    return bibleuser_classes.BibleUser.all()


def for_each_user(fn):
    debug_utils.log("Running {} for each user", [fn])

    # Read user database
    query = get_user_query()
    query.filter("active =", True)

    for dbUser in query.run(batch_size=500):
        fn(get_user(get_uid(dbUser)))


def migrate(userObj):
    newUserObj = bibleuser_classes.BibleUser(
        key_name=text_utils.to_string(userObj.get_uid()))

    newUserObj.clone(userObj)

    userObj.delete()
    newUserObj.update()

    return newUserObj