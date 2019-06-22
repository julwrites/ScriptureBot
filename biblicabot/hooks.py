# coding=utf-8

# Python std modules
import webapp2
import json

# Local modules
from common.utils import debug_utils

import modules

APP_HOOKS_URL = "/hooks"
APP_DAILY_HOOKS_URL = APP_HOOKS_URL + "/daily"


class HookHandler(webapp2.RequestHandler):
    def get(self):
        self.post()

    def post(self):
        hooks = modules.get_hooks()

        for hook in hooks:
            hook.dispatch()


app = webapp2.WSGIApplication(
    [
        # (url being accessed, class to call)
        (APP_DAILY_HOOKS_URL, HookHandler),
    ],
    debug=True)
