
# Local modules
from common import debug
from common import database
from common import telegram
from common import telegram_utils
from bgw import bgw

CMD_PASSAGE = '/passage'

def cmds(user, cmd, msg):
    debug.log('Running BGW commands')

    return (    \
    cmd_passage(user, cmd, msg) \
    )

def cmd_passage(user, cmd, msg):
    if user is not None:
        if cmd == CMD_PASSAGE:
            debug.log('Command: ' + cmd)

            query = msg.get('text')
            query = query.replace(cmd, '')

            text = bgw.get_passage(query)
            if text is not None:
                telegram.send_msg(text, user.get_uid())

                return True

    return False
