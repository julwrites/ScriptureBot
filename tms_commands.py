
# Local modules
import debug
import database
import telegram
import telegram_utils
import biblegateway
import tms

CMD_TMS = '/tms'

def cmds(user, cmd, msg):
    debug.log('Running TMS commands')

    return (    \
    cmd_tms(user, cmd, msg) \
    )

def cmd_tms(user, cmd, msg):
    if user is not None:
        if cmd == CMD_TMS:
            debug.log('Command: ' + cmd)

            query = msg.get('text')
            query = query.replace(cmd, '').strip()
            query = query.split(' ')

            debug.log('Attempting to get ' + '|'.join(query))

            verse = None
            if len(query) == 2:
                verse = tms.get_verse_by_pack(query[0], int(query[1]))

            if verse is not None:
                verse_text = biblegateway.get_passage(verse.reference, user.version)
                verse_msg = tms.format_verse(verse, verse_text)
                
                debug.log("Sending TMS verse: " + verse_msg)

                telegram.send_msg(verse_msg, user.get_uid())

                return True

    return False

