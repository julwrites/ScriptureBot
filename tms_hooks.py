

# Local modules
import tms
import debug
from modules import telegramtelegram
from modules import telegramtelegram_utils
import biblegateway

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
        verse_text = biblegateway.get_passage(verse.reference, user.version)
        verse_msg = tms.format_verse(verse, verse_text)

        debug.log("Sending verse: " + verse_msg)
        
        telegram.send_msg(verse_msg, uid)
   
    telegram_utils.foreach_user(send_verse)
 