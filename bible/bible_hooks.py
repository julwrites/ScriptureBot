
# Local modules
from tms import tms_utils
from common import debug
from common import telegram
from common import telegram_utils
from bible import bible_utils

from user.bibleuser_utils import *

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
            passage = bible_utils.get_passage_raw(verse.reference, user.get_version())
            verse_msg = tms_utils.format_verse(verse, passage)

            debug.log("Sending verse: " + verse_msg)
            
            telegram.send_msg(verse_msg, user.get_uid())


def hook_dailyverse():
    debug.log_hook(HOOK_DAILYVERSE)

    telegram_utils.foreach_user(resolve_dailyverse)
 