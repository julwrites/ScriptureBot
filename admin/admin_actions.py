
# Local modules
from common import debug
from common.user import user_utils
from common.telegram import telegram_utils
from common.action import action_classes

from admin import admin_utils


ADMIN_DUMP = '/dump'
ADMIN_DEBUG = '/doggle'
ADMIN_CLEAN = '/clean'
ADMIN_RAGNAROK = '/ragnarok'

class AdminDumpAction(action_classes.Action):
    def identifier(self):
        return '/dump'

    def resolve(self, userObj, msg):
        if admin_utils.access(userObj.get_uid()):
            # Read user database
            query = user_utils.get_user_query()
            query.filter('active =', True)

            try:
                userList = []
                for dbUser in query.run(batch_size=10):
                    userObj = user_utils.get_user(user_utils.get_uid(dbUser))
                    userList.append(userObj.get_description())
                userListMsg = '\n'.join(userList)
                telegram_utils.send_msg(userListMsg, userObj.get_uid())

            except Exception as e:
                debug.log(str(e))

            return True 
        return False

class AdminCleanAction(action_classes.Action):
    def identifier(self):
        return '/clean'

    def resolve(self, userObj, msg):
        if admin_utils.access(userObj.get_uid()):
            # Read user database
            query = user_utils.get_user_query()

            try:
                for dbUser in query.run():
                    userObj = user_utils.get_user(user_utils.get_uid(dbUser))
                    if userObj.get_name_string() == '-':
                        debug.log('Deleting: ' + userObj.get_uid())
                        userObj.delete()

                for dbUser in query.run():
                    userObj = user_utils.get_user(user_utils.get_uid(dbUser))
                    count = 0

                    for dbUserDup in query.run():
                        userObjDup = user_utils.get_user(user_utils.get_uid(dbUserDup))
                        if userObj.get_uid() == userObjDup.get_uid():
                            count += 1
                            if count > 1:
                                userObjDup.delete()

            except Exception as e:
                debug.log(str(e))

            return True
        return False

class AdminMigrateAction(action_classes.Action):
    def identifier(self):
        return '/refresh'

    def resolve(self, userObj, msg):
        if admin_utils.access(userObj.get_uid()):
            # Read user database
            query = user_utils.get_user_query()

            try:
                for dbUser in query.run():
                    userObj = user_utils.get_user(user_utils.get_uid(dbUser))
                    user_utils.migrate(userObj)
            except Exception as e:
                debug.log(str(e))

            return True
        return False

class AdminRagnarokAction(action_classes.Action):
    def identifier(self):
        return '/ragnarok'

    def resolve(self, userObj, msg):
        if admin_utils.access(userObj.get_uid()):
            # Read user database
            query = user_utils.get_user_query()

            try:
                for dbUser in query.run(batch_size=500):
                    userObj = user_utils.get_user(user_utils.get_uid(dbUser))
                    debug.log('Deleting: ' + userObj.get_uid())
                    userObj.delete()
            except Exception as e:
                debug.log(str(e))

            telegram_utils.send_msg("Baboomz~", userObj.get_uid())
            
            return True
        
        return False


def get():
    return [
        AdminDumpAction(),
        AdminCleanAction(),
        AdminMigrateAction(),
        AdminRagnarokAction(),
    ]