

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

def resolve_dailytms(user):
    if user is not None:
        if user.has_subscription(SUBSCRIPTION_DAILYTMS):
            verse = tms_utils.get_random_verse()
            passage = bgw_utils.get_passage_raw(verse.reference, user.get_version())
            verse_msg = tms_utils.format_verse(verse, passage)

            debug.log("Sending verse: " + verse_msg)
            
            telegram.send_msg(verse_msg, user.get_uid())

def hook_dailytms():
    debug.log_hook(HOOK_DAILYTMS)

    telegram_utils.foreach_user(resolve_dailytms)
 