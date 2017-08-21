

# Local modules
from tms import tms
from common import debug
from common import telegram
from common import telegram_utils
from bgw import bgw

from bible_user import *

HOOK_DAILYTMS = '/dailytms'

def hooks(data):
    debug.log('Running TMS hooks')

    return (    \
    hook_dailytms()   \
    )

def hook_dailytms():
    debug.log('Hook: ' + HOOK_DAILYTMS)

    def send_verse(uid):
        user = get_user(uid)

        verse = tms.get_verse_by_pack(user.get_current_pack(), user.get_current_verse())
        verse_text = bgw.get_passage(verse.reference, user.version)
        verse_msg = tms.format_verse(verse, verse_text)

        debug.log("Sending verse: " + verse_msg)
        
        telegram.send_msg(verse_msg, uid)
   
    telegram_utils.foreach_user(send_verse)
 