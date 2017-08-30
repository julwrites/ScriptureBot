
# Local modules
from common import debug
from common.telegram import telegram_utils
from common.user import user_utils

from tms import tms_utils

import bible


HOOK_DAILYVERSE = "/dailyverse"
SUBSCRIPTION_DAILYVERSE = "/*dailyverse*/"

def hooks(data):
    debug.log('Running Bible Query hooks')

    return (    \
    hook_dailyverse()   \
    )

def resolve_dailyverse(user):
    if user is not None:
        
        if user.has_subscription(SUBSCRIPTION_DAILYVERSE):
            verse = tms_utils.get_random_verse()
            passage = bible.utils.get_passage_raw(verse.reference, user.get_version())
            verse_msg = tms_utils.format_verse(verse, passage)

            debug.log("Sending verse: " + verse_msg)
            
            telegram_utils.send_msg(verse_msg, user.get_uid())


def hook_dailyverse():
    debug.log_hook(HOOK_DAILYVERSE)

    user_utils.for_each_user(resolve_dailyverse)
 