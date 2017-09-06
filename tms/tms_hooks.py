

# Local modules
from common import debug
from common.telegram import telegram_utils
from common.user import user_utils

from tms import tms_utils
from bible import bible_utils


HOOK_DAILYTMS = '/dailytms'
SUBSCRIPTION_DAILYTMS = '/*dailytms*/'

def hooks(data):
    debug.log('Running TMS hooks')

    return (    \
    hook_dailytms()   \
    )

def resolve_dailytms(userObj):
    if userObj is not None:
        if userObj.has_subscription(SUBSCRIPTION_DAILYTMS):
            verse = tms_utils.get_random_verse()
            passage = bible_utils.get_passage_raw(verse.reference, userObj.get_version())
            verseMsg = tms_utils.format_verse(verse, passage)

            debug.log("Sending verse: " + verseMsg)
            
            telegram_utils.send_msg(verseMsg, userObj.get_uid())

def hook_dailytms():
    debug.log_hook(HOOK_DAILYTMS)

    user_utils.for_each_user(resolve_dailytms)
 