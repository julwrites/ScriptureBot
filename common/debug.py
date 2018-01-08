# coding=utf-8

# Local modules
import logging

from common import chrono

DEBUG_MODE = True
VERBOSE_MODE = True


def debug():
    return DEBUG_MODE


def verbose():
    return VERBOSE_MODE


def toggle():
    DEBUG_MODE = not DEBUG_MODE


def log(msg):
    if not debug():
        return
    if not verbose():
        return
    logging.debug(msg)


def err(e):
    if not debug():
        return
    logging.debug("Error: " + text_utils.stringify(e))


def log_cmd(cmd):
    if not debug():
        return
    logging.debug("Command: " + cmd)


def log_state(state):
    if not debug():
        return
    logging.debug("State: " + state)


def log_action(action):
    if not debug():
        return
    logging.debug("Action: " + action)


def log_hook(hook):
    if not debug():
        return
    logging.debug("Hook: " + hook)


def datetime():
    if not debug():
        return

    logging.debug(text_utils.stringify(chrono.now()))