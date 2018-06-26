# coding=utf-8

import random

# Local Modules
from common import chrono, text_utils, debug, database


class BibleUser(database.Item):
    username = database.StringProperty(indexed=True)
    firstName = database.StringProperty(multiline=True, indexed=True)
    lastName = database.StringProperty(multiline=True, indexed=True)
    created = database.DateTimeProperty(indexed=True)
    lastReceived = database.DateTimeProperty(indexed=True)
    lastSent = database.DateTimeProperty(indexed=True)
    lastAuto = database.DateTimeProperty()
    active = database.BooleanProperty(indexed=True, default=True)
    state = database.StringProperty(indexed=True)
    version = database.StringProperty(indexed=True, default="NIV")
    subscriptions = database.StringListProperty(indexed=True)
    subscriptionTime = database.DateTimeProperty(indexed=True)

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
        return self.name()

    def get_name_string(self, verbose=False):
        fname = text_utils.stringify(self.firstName)
        lname = text_utils.stringify(self.lastName)
        uname = text_utils.stringify(self.username)

        name = fname

        if text_utils.is_valid(name):
            name += text_utils.stringify(" ") + lname if verbose else ""
        else:
            name = lname

        if text_utils.is_valid(name):
            name += text_utils.stringify(" @") + uname if verbose else ""
        else:
            name = uname

        return name

    def get_reply_string(self, strings):
        choose = random.randint(0, len(strings) - 1)
        reply = text_utils.stringify(strings[choose]).format(
            self.get_name_string())

        return reply

    def get_description(self):
        userType = "Group " if self.is_group() else "User "
        return text_utils.stringify(userType) + self.get_name_string(
            verbose=True)

    def is_group(self):
        return int(self.get_uid()) < 0

    def is_active(self):
        return self.active

    def set_active(self, active):
        self.active = active
        self.update()

    def get_state(self):
        return self.state

    def set_state(self, state):
        self.state = state
        self.update()

    def get_version(self):
        return self.version

    def set_version(self, version):
        self.version = version
        self.update()

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
        self.update()

    def remove_subscription(self, subId):
        try:
            self.subscriptions.remove(subId)
            self.update()
        except:
            return

    def update_last_received(self):
        self.lastReceived = chrono.now()
        self.update()

    def update_last_sent(self):
        self.lastSent = chrono.now()
        self.update()

    def update_last_auto(self):
        self.lastAuto = chrono.now()
        self.update()

    def refresh(self):
        self.update()
