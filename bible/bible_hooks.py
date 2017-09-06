
# Local modules
from common import debug
from common.user import user_utils
from common.telegram import telegram_utils

from tms import tms_utils
from bible import bible_utils


HOOK_DAILYVERSE = "/dailyverse"
SUBSCRIPTION_DAILYVERSE = "/*dailyverse*/"

def hooks(data):
    debug.log('Running Bible Query hooks')

    return (    \
    hook_dailyverse()   \
    )

def resolve_dailyverse(userObj):
    if userObj is not None:
        
        if userObj.has_subscription(SUBSCRIPTION_DAILYVERSE):
            verse = tms_utils.get_random_verse()
            passage = bible_utils.get_passage_raw(verse.reference, userObj.get_version())
            verseMsg = tms_utils.format_verse(verse, passage)

            debug.log("Sending verse: " + verseMsg)
            
            telegram_utils.send_msg(verseMsg, userObj.get_uid())


def hook_dailyverse():
    debug.log_hook(HOOK_DAILYVERSE)

    user_utils.for_each_user(resolve_dailyverse)
 