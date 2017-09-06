
# Local modules
from common import debug
from common.user import user_utils
from common.telegram import telegram_utils
from common.admin import admin_utils


ADMIN_DUMP = '/dump'
ADMIN_DEBUG = '/doggle'
ADMIN_CLEAN = '/clean'
ADMIN_RAGNAROK = '/ragnarok'

# List of commands to run through
def cmds(userId, cmd, msg):
    debug.log('Running admin commands')

    if admin_utils.access(userId):
        debug.log('Welcome, Master')
        try:
            return ( \
            cmd_dump(userId, cmd, msg)       \
            or cmd_doggle(userId, cmd, msg)   \
            or cmd_clean(userId, cmd, msg)   \
            or cmd_ragnarok(userId, cmd, msg)   \
            )
        except:
            debug.log('Exception in Admin Commands')
            return False

# Debug Commands
def cmd_dump(userId, cmd, msg):
    if admin_utils.access(userId) and cmd == ADMIN_DUMP:
        debug.log_cmd(cmd)

        # Read user database
        query = user_utils.get_user_query()
        query.filter('active =', True)

        try:
            userList = []
            for dbUser in query.run(batch_size=10):
                userObj = user_utils.get_user(user_utils.get_uid(dbUser))
                userList.append(userObj.get_description())
            userListMsg = '\n'.join(userList)
            telegram_utils.send_msg(userListMsg, userId)
        except Exception as e:
            debug.log(str(e))

        return True 

    return False

def cmd_doggle(userId, cmd, msg):
    if admin_utils.access(userId) and cmd == ADMIN_DEBUG:
        debug.log_cmd(cmd)

        debug.toggle()

        return True

    return False

def cmd_clean(userId, cmd, msg):
    if admin_utils.access(userId) and cmd == ADMIN_CLEAN:
        debug.log_cmd(cmd)

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

def cmd_ragnarok(userId, cmd, msg):
    if admin_utils.access(userId) and cmd == ADMIN_RAGNAROK:
        debug.log_cmd(cmd)

        # Read user database
        query = user_utils.get_user_query()

        try:
            for dbUser in query.run(batch_size=500):
                userObj = user_utils.get_user(user_utils.get_uid(dbUser))
                debug.log('Deleting: ' + userObj.get_uid())
                userObj.delete()
        except Exception as e:
            debug.log(str(e))

        telegram_utils.send_msg("Baboomz~", userId)
        
        return True
    
    return False
 