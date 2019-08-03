# coding=utf-8

# Python std modules
import webapp2
import json

# Local modules
from common.utils import debug_utils

import modules

APP_HOOKS_URL = "/hooks"
APP_DAILY_HOOKS_URL = APP_HOOKS_URL + "/daily"

app = Flask(__name__)


@app.route(APP_DAILY_HOOKS_URL, methods=["GET", "POST"])
def main():
    hooks = modules.get_hooks()

    for hook in hooks:
        hook.dispatch()

    return ''


if __name__ == '__main__':
    app.run(debug=True)