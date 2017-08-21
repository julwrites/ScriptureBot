
# Python std modules
import webapp2

# Local modules
import debug
import database
import modules.telegram
import modules.telegram_utils
import biblegateway

import tms_hooks

from constants import APP_HOOKS_URL

app = webapp2.WSGIApplication([
    # (url being accessed, class to call)
    (APP_HOOKS_URL + tms_hooks.HOOK_DAILYTMS, tms_hooks.hooks)
], debug=True)
 