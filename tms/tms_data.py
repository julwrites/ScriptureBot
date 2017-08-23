
from tms import tms_loa
from tms import tms_60
from tms import tms_lifeissues

TMS_DATA = \
    tms_loa.data() +        \
    tms_60.data() +         \
    tms_lifeissues.data()

TMS_KEYS = \
    tms_loa.keys() +        \
    tms_60.keys() +         \
    tms_lifeissues.keys()

TMS_ALIAS = \
    tms_loa.aliases() +        \
    tms_60.aliases() +         \
    tms_lifeissues.aliases()

TMS_NAMES = \
    tms_loa.names() +        \
    tms_60.names() +         \
    tms_lifeissues.names()


def get_keys():
    return TMS_KEYS

def get_data():
    return TMS_DATA

def get_aliases():
    return TMS_ALIAS

def get_names():
    return TMS_NAMES

def get_top():
    return tms_loa.top()