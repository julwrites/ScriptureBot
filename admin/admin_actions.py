# coding=utf-8

# Local modules
from common import debug, text_utils
from common.telegram import telegram_utils
from common.action import action_classes

from user import user_utils
from admin import admin_utils


class AdminNotifyAction(action_classes.Action):
    def identifier(self):
        return "/notify"

    def name(self):
        return "Notifies users about Bot Update"

    def resolve(self, userObj, msg):
        if admin_utils.access(userObj.get_uid()):
            # Read user database
            query = user_utils.get_user_query()
            query.filter("active =", True)

            msg = telegram_utils.strip_command(msg, self.identifier())

            userList = []
            for dbUser in query.run(batch_size=10):
                dbUserObj = user_utils.get_user(user_utils.get_uid(dbUser))
                telegram_utils.send_msg(
                    user=dbUserObj.get_uid(),
                    text=msg.format(dbUserObj.get_name_string(verbose=True)))

            return True
        return False


class AdminDumpAction(action_classes.Action):
    def identifier(self):
        return "/dump"

    def name(self):
        return "Dump User Database"

    def resolve(self, userObj, msg):
        if admin_utils.access(userObj.get_uid()):
            # Read user database
            query = user_utils.get_user_query()
            query.filter("active =", True)

            userList = []
            for dbUser in query.run(batch_size=10):
                dbUserObj = user_utils.get_user(user_utils.get_uid(dbUser))
                userList.append(dbUserObj.get_description())
            userListMsg = "\n".join(userList[start:end])
            telegram_utils.send_msg(user=userObj.get_uid(), text=userListMsg)

            return True
        return False


class AdminCleanAction(action_classes.Action):
    def identifier(self):
        return "/clean"

    def name(self):
        return "Clean User Database"

    def resolve(self, userObj, msg):
        if admin_utils.access(userObj.get_uid()):
            # Read user database
            query = user_utils.get_user_query()

            for dbUser in query.run():
                dbUserObj = user_utils.get_user(user_utils.get_uid(dbUser))
                if dbUserObj.get_name_string() == "-":
                    debug.log("Deleting: " + dbUserObj.get_uid())
                    dbUserObj.delete()

            for dbUser in query.run():
                dbUserObj = user_utils.get_user(user_utils.get_uid(dbUser))
                count = 0

                for dbUserDup in query.run():
                    dbUserObjDup = user_utils.get_user(
                        user_utils.get_uid(dbUserDup))
                    if dbUserObj.get_uid() == dbUserObjDup.get_uid():
                        count += 1
                        if count > 1:
                            dbUserObjDup.delete()

            return True
        return False


class AdminMigrateAction(action_classes.Action):
    def identifier(self):
        return "/refresh"

    def name(self):
        return "Migrate User Database"

    def resolve(self, userObj, msg):
        if admin_utils.access(userObj.get_uid()):
            # Read user database
            query = user_utils.get_user_query()

            for dbUser in query.run():
                dbUserObj = user_utils.get_user(user_utils.get_uid(dbUser))
                user_utils.migrate(dbUserObj)

            return True
        return False


class AdminRagnarokAction(action_classes.Action):
    def identifier(self):
        return "/ragnarok"

    def name(self):
        return "Kill User Database"

    def resolve(self, userObj, msg):
        if admin_utils.access(userObj.get_uid()):
            # Read user database
            query = user_utils.get_user_query()

            for dbUser in query.run(batch_size=500):
                dbUserObj = user_utils.get_user(user_utils.get_uid(dbUser))
                debug.log("Deleting: " + dbUserObj.get_uid())
                dbUserObj.delete()

            telegram_utils.send_msg(user=userObj.get_uid(), text="Baboomz~")

            return True

        return False


class AdminFeedbackAction(action_classes.Action):
    def identifier(self):
        return "/feedback"

    def name(self):
        return "Feedback"

    def description(self):
        return "Send Feedback to Administrator"

    def resolve(self, userObj, msg):
        query = telegram_utils.strip_command(msg, self.identifier())

        if text_utils.is_valid(query):
            telegram_utils.send_msg(user=admin_utils.BOT_ADMIN, text=query)
        else:
            userObj.set_state(self.identifier())
        return True


def get():
    return [
        AdminNotifyAction(),
        AdminDumpAction(),
        AdminCleanAction(),
        AdminMigrateAction(),
        AdminRagnarokAction(),
        AdminFeedbackAction(),
    ]