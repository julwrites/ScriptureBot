
# Python std modules
import webapp2
import json

# Local modules
from common import debug

import components

APP_HOOKS_URL = "/hooks"
APP_DAILY_HOOKS_URL = APP_HOOKS_URL + "/daily"

class HookHandler(webapp2.RequestHandler):
    def get(self):
        self.post()

    def post(self):
        data = json.loads(self.request.body)
        debug.log(data)

        hooks = components.hooks()

        for hook in hooks:
            hook.dispatch(data):



app = webapp2.WSGIApplication([
    # (url being accessed, class to call)
    (APP_DAILY_HOOKS_URL, HookHandler),
], debug=True)
 