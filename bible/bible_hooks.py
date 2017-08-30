
# Local modules
from common import debug, telegram
from common.user import user_utils

from tms import tms_utils
from bible import bible_utils


HOOK_DAILYVERSE = "/dailyverse"
SUBSCRIPTION_DAILYVERSE = "/*dailyverse*/"

def hooks(data):
    debug.log('Running Bible Query hooks')

    return (    \
    hook_dailyverse()   \
    )

def resolve_dailyverse(user_obj):
    if user_obj is not None:
        
        if user_obj.has_subscription(SUBSCRIPTION_DAILYVERSE):
            verse = tms_utils.get_random_verse()
            passage = bible_utils.get_passage_raw(verse.reference, user_obj.get_version())
            verse_msg = tms_utils.format_verse(verse, passage)

            debug.log("Sending verse: " + verse_msg)
            
            telegram_utils.send_msg(verse_msg, user_obj.get_uid())


def hook_dailyverse():
    debug.log_hook(HOOK_DAILYVERSE)

    user_utils.for_each_user(resolve_dailyverse)
 