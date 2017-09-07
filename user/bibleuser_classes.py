
# Google App Engine Modules
from google.appengine.ext import db

# Local Modules
from common import chrono, text_utils, debug

class BibleUser(db.Model):
    username = db.StringProperty(indexed=True)
    firstName = db.StringProperty(multiline=True, indexed=True)
    lastName = db.StringProperty(multiline=True, indexed=True)
    created = db.DateTimeProperty(auto_now_add=True, indexed=True)
    lastReceived = db.DateTimeProperty(auto_now_add=True, indexed=True)
    lastSent = db.DateTimeProperty(indexed=True)
    lastAuto = db.DateTimeProperty(auto_now_add=True)
    active = db.BooleanProperty(default=True)
    state = db.StringProperty(indexed=True)
    version = db.StringProperty(indexed=True, default='NIV')
    subscription = db.StringProperty(indexed=True)
    subscriptions = db.StringListProperty(indexed=True)
    subscriptionTime = db.IntegerProperty(indexed=True)

    def clone(self, obj):
        self.username = obj.username
        self.firstName = obj.firstName
        self.lastName = obj.lastName
        self.created = obj.created
        self.lastReceived = obj.lastReceived
        self.lastSent = obj.lastSent
        self.lastAuto = obj.lastAuto
        self.active = obj.active
        self.state = obj.state
        self.version = obj.version
        self.subscriptions = obj.subscriptions
        self.subscriptionTime = obj.subscriptionTime
        return self

    def get_uid(self):
        return self.key().name()

    def get_name_string(self):
        def prep(string):
            return string.encode('utf-8', 'ignore').strip()

        name = prep(self.firstName)
        if self.lastName:
            name += ' ' + prep(self.lastName)
        if self.username:
            name += ' @' + prep(self.username)

        return name

    def get_description(self):
        userType = 'Group' if self.is_group() else 'User'
        return userType + ' ' + self.get_name_string()

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
        return ','.join(self.subscription)

    def add_subscription(self, subId):
        if self.has_subscription(subId):
            return

        self.subscription.append(subId)
        self.put()

    def remove_subscription(self, subId):
        try:
            self.subscription.remove(subId)
            self.put()
        except:
            return

    def has_subscription(self, subId):
        try:
            self.subscription.index(subId)
        except:
            return False
        return True

    def get_subscription_time(self):
        return self.subscriptionTime

    def set_subscription_time(self, time):
        self.subscriptionTime = time
        self.put()

    def update_last_received(self):
        self.lastReceived = chrono.now()
        self.put()

    def update_last_sent(self):
        self.lastSent = chrono.now()
        self.put()

    def update_last_auto(self):
        self.lastAuto = chrono.now()
        self.put()

    def refresh(self):
        self.put()
