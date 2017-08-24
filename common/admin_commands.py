
# Local modules
from common import debug
from common import admin
from common import database
from common import telegram

from user.bibleuser_utils import *

ADMIN_DUMP = '/dump'
ADMIN_DEBUG = '/doggle'
ADMIN_CLEAN = '/clean'
ADMIN_RAGNAROK = '/ragnarok'

# List of commands to run through
def cmds(uid, cmd, msg):
    debug.log('Running admin commands')

    try:
        result = ( \
        cmd_dump(uid, cmd, msg)       \
        or cmd_doggle(uid, cmd, msg)   \
        or cmd_clean(uid, cmd, msg)   \
        or cmd_ragnarok(uid, cmd, msg)   \
        )
    except:
        debug.log('Exception in Admin Commands')
        return False
    return result

# Debug Commands
def cmd_dump(uid, cmd, msg):
    if admin.access(uid) and cmd == ADMIN_DUMP:
        debug.log_cmd(cmd)

        # Read user database
        query = get_user_query()
        query.filter('active =', True)

        try:
            user_list = []
            for user in query.run(batch_size=10):
                dbUser = get_user(get_uid(user))
                user_list.append(dbUser.get_description())
            user_list_msg = '\n'.join(user_list)
            telegram.send_msg(user_list_msg, uid)
        except Exception as e:
            debug.log(str(e))

        return True 

    return False

def cmd_doggle(uid, cmd, msg):
    if admin.access(uid) and cmd == ADMIN_DEBUG:
        debug.log_cmd(cmd)

        debug.toggle()

        return True

    return False

def cmd_clean(uid, cmd, msg):
    if admin.access(uid) and cmd == ADMIN_CLEAN:
        debug.log_cmd(cmd)

        # Read user database
        query = get_user_query

        try:
            for user in query.run():
                dbUser = get_user(get_uid(user))
                if dbUser.get_name_string() == '-':
                    debug.log('Deleting: ' + dbUser.get_uid())
                    dbUser.delete()

            for user in query.run():
                dbUser = get_user(get_uid(user))
                count = 0

                for dup in query.run():
                    dbDup = get_user(get_uid(dup))
                    if dbDup.get_uid() == user.get_uid():
                        count += 1
                        if count > 1:
                            dbDup.delete()

        except Exception as e:
            debug.log(str(e))
        
        return True

    return False

def cmd_ragnarok(uid, cmd, msg):
    if admin.access(uid) and cmd == ADMIN_RAGNAROK:
        debug.log_cmd(cmd)

        # Read user database
        query = get_user_query

        try:
            for user in query.run(batch_size=500):
                dbUser = get_user(get_uid(user))
                debug.log('Deleting: ' + dbUser.get_uid())
                dbUser.delete()
        except Exception as e:
            debug.log(str(e))

        telegram.send_msg("Baboomz~", uid)
        
        return True
    
    return False
 