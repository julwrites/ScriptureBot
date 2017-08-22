
# Local modules
from common import debug
from common import database
from common import telegram
from common import telegram_utils

from bgw import bgw_utils

CMD_PASSAGE = '/passage'
CMD_PASSAGE_PROMPT = 'Please enter /passage followed by the Bible passage you desire'

def cmds(user, cmd, msg):
    if user is None:
        return False

    debug.log('Running BGW commands')

    return (    \
    cmd_passage(user, cmd, msg) \
    )

def cmd_passage(user, cmd, msg):
    if cmd == CMD_PASSAGE:
        debug.log('Command: ' + cmd)

        query = msg.get('text')
        query = query.replace(cmd, '')

        text = bgw_utils.get_passage(query, user.get_version())
        if text is not None:
            telegram.send_msg(text, user.get_uid())
        else:
            telegram.send_msg(CMD_PASSAGE_PROMPT, user.get_uid())

        return True

    return False
