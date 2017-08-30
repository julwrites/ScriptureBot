
# Local modules
from common import debug
from common import user
from common import telegram
from common import admin


ADMIN_DUMP = '/dump'
ADMIN_DEBUG = '/doggle'
ADMIN_CLEAN = '/clean'
ADMIN_RAGNAROK = '/ragnarok'

# List of commands to run through
def cmds(user_id, cmd, msg):
    debug.log('Running admin commands')

    if admin.utils.access(user_id):
        debug.log('Welcome, Master')
        try:
            return ( \
            cmd_dump(user_id, cmd, msg)       \
            or cmd_doggle(user_id, cmd, msg)   \
            or cmd_clean(user_id, cmd, msg)   \
            or cmd_ragnarok(user_id, cmd, msg)   \
            )
        except:
            debug.log('Exception in Admin Commands')
            return False

# Debug Commands
def cmd_dump(user_id, cmd, msg):
    if admin.utils.access(user_id) and cmd == ADMIN_DUMP:
        debug.log_cmd(cmd)

        # Read user database
        query = user.utils.get_user_query()
        query.filter('active =', True)

        try:
            user_list = []
            for db_user in query.run(batch_size=10):
                user_obj = user.utils.get_user(user.utils.get_uid(db_user))
                user_list.append(user_obj.get_description())
            user_list_msg = '\n'.join(user_list)
            telegram.utils.send_msg(user_list_msg, user_id)
        except Exception as e:
            debug.log(str(e))

        return True 

    return False

def cmd_doggle(user_id, cmd, msg):
    if admin.utils.access(user_id) and cmd == ADMIN_DEBUG:
        debug.log_cmd(cmd)

        debug.toggle()

        return True

    return False

def cmd_clean(user_id, cmd, msg):
    if admin.utils.access(user_id) and cmd == ADMIN_CLEAN:
        debug.log_cmd(cmd)

        # Read user database
        query = user.utils.get_user_query()

        try:
            for db_user in query.run():
                user_obj = user.utils.get_user(user.utils.get_uid(db_user))
                if user_obj.get_name_string() == '-':
                    debug.log('Deleting: ' + user_obj.get_uid())
                    user_obj.delete()

            for db_user in query.run():
                user_obj = user.utils.get_user(user.utils.get_uid(db_user))
                count = 0

                for db_user_dup in query.run():
                    user_obj_dup = user.utils.get_user(user.utils.get_uid(db_user_dup))
                    if user_obj.get_uid() == user_obj_dup.get_uid():
                        count += 1
                        if count > 1:
                            user_obj_dup.delete()

        except Exception as e:
            debug.log(str(e))
        
        return True

    return False

def cmd_ragnarok(user_id, cmd, msg):
    if admin.utils.access(user_id) and cmd == ADMIN_RAGNAROK:
        debug.log_cmd(cmd)

        # Read user database
        query = user.utils.get_user_query()

        try:
            for db_user in query.run(batch_size=500):
                user_obj = user.utils.get_user(user.utils.get_uid(db_user))
                debug.log('Deleting: ' + user_obj.get_uid())
                user_obj.delete()
        except Exception as e:
            debug.log(str(e))

        telegram.utils.send_msg("Baboomz~", user_id)
        
        return True
    
    return False
 