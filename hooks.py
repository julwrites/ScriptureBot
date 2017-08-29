
# Python std modules
import webapp2

from tms import tms_hooks
from bible import bible_hooks
from devo import devo_hooks

APP_HOOKS_URL = "/hooks"

app = webapp2.WSGIApplication([
    # (url being accessed, class to call)
    (APP_HOOKS_URL + tms_hooks.HOOK_DAILYTMS, tms_hooks.hooks),
    (APP_HOOKS_URL + bible_hooks.HOOK_DAILYVERSE, bible_hooks.hooks),
    (APP_HOOKS_URL + devo_hooks.HOOK_DAILYDEVO, devo_hooks.hooks)
], debug=True)
 