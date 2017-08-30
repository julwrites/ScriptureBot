
# Local modules
from common import debug, telegram
from common import user

import tms
import bible


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
            verse = tms.utils.get_random_verse()
            passage = bible.utils.get_passage_raw(verse.reference, user_obj.get_version())
            verse_msg = tms.utils.format_verse(verse, passage)

            debug.log("Sending verse: " + verse_msg)
            
            telegram.utils.send_msg(verse_msg, user_obj.get_uid())


def hook_dailyverse():
    debug.log_hook(HOOK_DAILYVERSE)

    user.utils.for_each_user(resolve_dailyverse)
 