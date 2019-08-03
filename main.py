# coding=utf-8

# Python std modules
import flask
from flask import Flask, escape, request
import json
import logging

# Local modules
from common.utils import debug_utils, text_utils
from common.telegram import telegram_utils
from common.action import action_utils
from user import user_utils

import actions

from secret import BOT_ID
APP_BOT_URL = "/" + BOT_ID

app = Flask(__name__)


@app.route(APP_BOT_URL, methods=["POST"])
def webhook():
    if flask.request.headers.get('content-type') == 'application/json':
        data = flask.request.get_data().decode('utf-8')
    else:
        flask.abort(400)

    data = request.get_json()
    debug_utils.log(data)

    if data.get("message"):
        msg = data.get("message")

        # Read the user to echo back
        userId = user_utils.get_uid(msg.get("from").get("id"))
        userObj = user_utils.get_user(userId)

        if action_utils.execute(actions.get(), userObj, msg):
            return ''

        telegram_utils.send_msg(
            user=msg.get("from").get("id"), text="Hello, I am bot")

    return ''


if __name__ == '__main__':
    app.run(debug=True)