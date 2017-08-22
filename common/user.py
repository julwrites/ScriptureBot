
# Google App Engine Modules
from google.appengine.ext import db

# Local Modules
from common import chrono
from common import database
from common import text_utils

class User(db.Model):
    username = db.StringProperty(indexed=False)
    first_name = db.StringProperty(multiline=True, indexed=False)
    last_name = db.StringProperty(multiline=True, indexed=False)
    created = db.DateTimeProperty(auto_now_add=True)
    last_received = db.DateTimeProperty(auto_now_add=True, indexed=False)
    last_sent = db.DateTimeProperty(indexed=False)
    last_auto = db.DateTimeProperty(auto_now_add=True)
    active = db.BooleanProperty(default=True)
    state = db.StringProperty(indexed=False)

    def get_uid(self):
        return self.key().name()

    def get_name_string(self):
        def prep(string):
            return string.encode('utf-8', 'ignore').strip()

        name = prep(self.first_name)
        if self.last_name:
            name += ' ' + prep(self.last_name)
        if self.username:
            name += ' @' + prep(self.username)

        return name

    def get_description(self):
        user_type = 'Group' if self.is_group() else 'User'
        return user_type + ' ' + self.get_name_string() + ' ' + self.last_sent

    def is_group(self):
        return int(self.get_uid()) < 0

    def is_active(self):
        return self.active

    def set_active(self, active):
        self.active = active
        self.put()

    def get_state(self):
        return self.state

    def set_state(self, state):
        self.state = state
        self.put()

    def get_version(self):
        return self.version

    def set_version(self, version):
        self.version = version
        self.put()

    def update_last_received(self):
        self.last_received = chrono.now()
        self.put()

    def update_last_sent(self):
        self.last_sent = chrono.now()
        self.put()

    def update_last_auto(self):
        self.last_auto = chrono.now()
        self.put()

    def migrate_to(self, uid):
        props = dict((prop, getattr(self, prop)) for prop in self.properties().keys())
        props.update(key_name=str(uid))
        new_user = User(**props)
        new_user.put()
        self.delete()
        return new_user

# Functions for manipulation of user info
def get_user(uid):
    user = db.get(database.get_key('User', uid))
    return user

def get_uid(user):
    try:
        uid = user.get_uid()
    except AttributeError:
        uid = user

    return uid

def set_profile(uid, uname, fname, lname):
    existing_user = get_user(uid)

    uname = text_utils.stringify(uname)
    fname = text_utils.stringify(fname)
    lname = text_utils.stringify(fname)

    if existing_user:
        existing_user.username = uname
        existing_user.first_name = fname
        existing_user.last_name = lname
        existing_user.update_last_received()
        return existing_user
    else:
        user = User(key_name=str(uid), username=uname, first_name=fname, last_name=lname)
        user.put()
        return user

def get_user_query():
    return User.all()

