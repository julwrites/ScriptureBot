# coding=utf-8

# Local modules
import logging

from common import chrono
from common.utils import text_utils

DEBUG_MODE = True
VERBOSE_MODE = True


def debug():
    return DEBUG_MODE


def verbose():
    return VERBOSE_MODE


def toggle():
    DEBUG_MODE = not DEBUG_MODE


def log(msg, args=[]):
    if debug():
        logging.getLogger().setLevel(logging.DEBUG)
    else:
        logging.getLogger().setLevel(logging.INFO)

    if not verbose():
        return

    if len(args) > 0:
        msg = msg.format(*[text_utils.to_utf8(arg) for arg in args])
    logging.getLogger().debug(msg)


def err(e):
    if not debug():
        return
    log("Error: {}", [e])


def log_cmd(cmd):
    if not debug():
        return
    log("Command: {}", [cmd])


def log_state(state):
    if not debug():
        return
    log("State: {}", [state])


def log_action(action):
    if not debug():
        return
    log("Action: {}", [action])


def log_hook(hook):
    if not debug():
        return
    log("Hook: {}", [hook])


def datetime():
    if not debug():
        return

    log("Time: {}", [chrono.now()])