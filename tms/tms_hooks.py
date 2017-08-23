

# Local modules
from tms import tms_utils
from common import debug
from common import telegram
from common import telegram_utils
from bgw import bgw_utils

from user.bibleuser_utils import *

HOOK_DAILYTMS = '/dailytms'
SUBSCRIPTION_DAILYTMS = '/*dailytms*/'

def hooks(data):
    debug.log('Running TMS hooks')

    return (    \
    hook_dailytms()   \
    )

def hook_dailytms():
    debug.log_hook(HOOK_DAILYTMS)

    def send_verse(uid):
        user = get_user(uid)
        
        if user.get_daily_subscription().find(SUBSCRIPTION_DAILYTMS) is not -1:
            verse = tms_utils.get_random_verse()
            passage = bgw_utils.get_passage_raw(verse.reference, user.get_version())
            verse_msg = tms_utils.format_verse(verse, passage)

            debug.log("Sending verse: " + verse_msg)
            
            telegram.send_msg(verse_msg, uid)

    telegram_utils.foreach_user(send_verse)
 