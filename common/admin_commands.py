
# Local modules
from common import debug
from common import admin
from common import database
from common import telegram

from bible_user import *

ADMIN_DUMP = '/dump'
ADMIN_DEBUG = '/doggle'
ADMIN_CLEAN = '/clean'
ADMIN_RAGNAROK = '/ragnarok'

# List of commands to run through
def cmds(user, cmd, msg):
    debug.log('Running admin commands')

    return ( \
    cmd_dump(user, cmd, msg)       \
    or cmd_doggle(user, cmd, msg)   \
    or cmd_clean(user, cmd, msg)   \
    or cmd_ragnarok(user, cmd, msg)   \
    )

# Debug Commands
def cmd_dump(user, cmd, msg):
    adminID = user.get_uid()
    if admin.access(adminID) and cmd == ADMIN_DUMP:
        debug.log('Command: ' + cmd)

        # Read user database
        query = get_user_query()
        query.filter('active =', True)

        try:
            for user in query.run(batch_size=500):
                dbUser = get_user(get_uid(user))
                telegram.send_msg('User: ' 
                + dbUser.get_description()
                , adminID)
        except Exception as e:
            debug.log(str(e))

        # Log user database
        debug.log('User list: ' + str(query))

        return True 

    return False

def cmd_doggle(user, cmd, msg):
    adminID = user.get_uid()
    if admin.access(adminID) and cmd == ADMIN_DEBUG:
        debug.log('Command: ' + cmd)

        debug.toggle()

        return True

    return False

def cmd_clean(user, cmd, msg):
    adminID = user.get_uid()
    if admin.access(adminID) and cmd == ADMIN_CLEAN:
        debug.log('Command: ' + cmd)

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

def cmd_ragnarok(user, cmd, msg):
    adminID = user.get_uid()
    if admin.access(adminID) and cmd == ADMIN_RAGNAROK:
        debug.log('Command: ' + cmd)

        # Read user database
        query = get_user_query

        try:
            for user in query.run(batch_size=500):
                dbUser = get_user(get_uid(user))
                debug.log('Deleting: ' + dbUser.get_uid())
                dbUser.delete()
        except Exception as e:
            debug.log(str(e))

        telegram.send_msg("Baboomz~", adminID)
        
        return True
    
    return False
 