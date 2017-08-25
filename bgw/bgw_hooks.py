
# Local modules
from tms import tms_utils
from common import debug
from common import telegram
from common import telegram_utils
from bgw import bgw_utils

from user.bibleuser_utils import *

HOOK_DAILYBGW = "/dailybgw"
SUBSCRIPTION_DAILYBGW = "/*dailybgw*/"

def hooks(data):
    debug.log('Running BGW hooks')

    return (    \
    hook_dailybgw()   \
    )

def resolve_dailybgw(user):
    if user is not None:
        
        if user.has_subscription(SUBSCRIPTION_DAILYBGW):
            verse = tms_utils.get_random_verse()
            passage = bgw_utils.get_passage_raw(verse.reference, user.get_version())
            verse_msg = tms_utils.format_verse(verse, passage)

            debug.log("Sending verse: " + verse_msg)
            
            telegram.send_msg(verse_msg, user.get_uid())


def hook_dailybgw():
    debug.log_hook(HOOK_DAILYBGW)

    telegram_utils.foreach_user(resolve_dailybgw)
 