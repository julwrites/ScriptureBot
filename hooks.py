
# Python std modules
import webapp2

# Local modules
from common import debug
from common import database
from common import telegram
from common import telegram_utils
from bgw import bgw

from tms import tms_hooks

from common.constants import APP_HOOKS_URL

app = webapp2.WSGIApplication([
    # (url being accessed, class to call)
    (APP_HOOKS_URL + tms_hooks.HOOK_DAILYTMS, tms_hooks.hooks)
], debug=True)
 