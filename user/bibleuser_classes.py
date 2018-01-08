# coding=utf-8

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
    version = db.StringProperty(indexed=True, default="NIV")
    subscriptions = db.StringListProperty(indexed=True)
    subscriptionTime = db.DateTimeProperty(indexed=True)

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

    def get_name_string(self, username=False):
        name = text_utils.stringify(self.firstName)
        if text_utils.is_valid(
                self.lastName) and self.firstName != self.lastName:
            name += text_utils.stringify(" ") + text_utils.stringify(
                self.lastName)
        if username and text_utils.is_valid(self.username):
            name += text_utils.stringify(" @") + text_utils.stringify(
                self.username)

        return name

    def get_description(self):
        userType = "Group " if self.is_group() else "User "
        return text_utils.stringify(userType) + self.get_name_string(
            username=True)

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
        return ",".join(self.subscriptions)

    def has_subscription(self, subId):
        try:
            self.subscriptions.index(subId)
        except:
            return False
        return True

    def add_subscription(self, subId):
        if self.has_subscription(subId):
            return

        self.subscriptions.append(subId)
        self.put()

    def remove_subscription(self, subId):
        try:
            self.subscriptions.remove(subId)
            self.put()
        except:
            return

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
