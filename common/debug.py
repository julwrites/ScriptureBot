
# coding=utf-8

# Local modules
import logging

from common import chrono

DEBUG_MODE = True

def debug():
    return DEBUG_MODE

def toggle():
    DEBUG_MODE = not DEBUG_MODE

def log(msg):
    if not debug():
        return
    logging.debug(msg)

def log_cmd(cmd):
    if not debug():
        return
    logging.debug('Command: ' + cmd)

def log_state(state):
    if not debug():
        return
    logging.debug('State: ' + state)

def log_action(action):
    if not debug():
        return
    logging.debug('Action: ' + action)

def log_hook(hook):
    if not debug():
        return
    logging.debug('Hook: ' + hook)

def datetime():
    if not debug():
        return

    logging.debug(str(chrono.now()))