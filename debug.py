
# Local modules
import logging
import chrono

DEBUG_MODE = True

def debug():
    return DEBUG_MODE

def toggle():
    DEBUG_MODE = not DEBUG_MODE

def log(msg):
    if not debug():
        return

    logging.debug(msg)

def datetime():
    if not debug():
        return

    logging.debug(str(chrono.now()))