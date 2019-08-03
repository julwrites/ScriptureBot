# coding=utf-8

# Python std modules
import flask
from flask import Flask, request
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
    logging.warn("Warning start hor")
    data = request.get_json(force=True)
    debug_utils.log(data)
    logging.warn("I warn you again ah")

    if data.get("message"):
        logging.warn("Warn you about message")
        msg = data.get("message")

        # Read the user to echo back
        userId = user_utils.get_uid(msg.get("from").get("id"))
        userObj = user_utils.get_user(userId)

        logging.warn("User is here, be warned")
        if action_utils.execute(actions.get(), userObj, msg):
            return ''

        logging.warn("Warn you of failure")
        telegram_utils.send_msg(
            user=msg.get("from").get("id"), text="Hello, I am bot")

    logging.warn("Warn you ah, exit is here")

    return ''


if __name__ == '__main__':
    app.run(debug=True)