
# Python std modules
import webapp2

from bible import bible_hooks

APP_HOOKS_URL = "/hooks"

app = webapp2.WSGIApplication([
    # (url being accessed, class to call)
    (APP_HOOKS_URL + bible_hooks.get().identifier(), bible_hooks.get().execute),
], debug=True)
 