
# Google App Engine Modules
from google.appengine.ext import db

# Local Modules
from common import chrono, text_utils

class BibleUser(db.Model):
    username = db.StringProperty(indexed=False)
    first_name = db.StringProperty(multiline=True, indexed=False)
    last_name = db.StringProperty(multiline=True, indexed=False)
    created = db.DateTimeProperty(auto_now_add=True)
    last_received = db.DateTimeProperty(auto_now_add=True, indexed=False)
    last_sent = db.DateTimeProperty(indexed=False)
    last_auto = db.DateTimeProperty(auto_now_add=True)
    active = db.BooleanProperty(default=True)
    state = db.StringProperty(indexed=False)
    version = db.StringProperty(indexed=False, default='NIV')
    subscription = db.StringProperty(indexed=False)

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
        return user_type + ' ' + self.get_name_string()

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

    def get_subscription(self):
        return self.subscription

    def add_subscription(self, subscription):
        if self.has_subscription(subscription):
            return

        self.subscription += subscription
        self.put()

    def has_subscription(self, subscription):
        if text_utils.is_valid(self.subscription):
            return self.subscription.find(subscription) != -1
        return False

    def update_last_received(self):
        self.last_received = chrono.now()
        self.put()

    def update_last_sent(self):
        self.last_sent = chrono.now()
        self.put()

    def update_last_auto(self):
        self.last_auto = chrono.now()
        self.put()

    def migrate_to(self, user_id):
        props = dict((prop, getattr(self, prop)) for prop in self.properties().keys())
        props.update(key_name=str(user_id))
        new_user = BibleUser(**props)
        new_user.put()
        self.delete()
        return new_user
