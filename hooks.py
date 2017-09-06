
# Python std modules
import webapp2
import json

# Local modules
from common import debug
from bible import bible_hooks

APP_HOOKS_URL = "/hooks"

class HookHandler(webapp2.RequestHandler):
    def get(self):
        self.post()

    def post(self):
        data = json.loads(self.request.body)
        debug.log(data)

        actions = bible_hooks.get()

        for action in actions:
            if action.execute(data):
                return



app = webapp2.WSGIApplication([
    # (url being accessed, class to call)
    (APP_HOOKS_URL, HookHandler),
], debug=True)
 