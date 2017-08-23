
from tms import tms_loa
from tms import tms_60
from tms import tms_lifeissues

def merge(lhs, rhs):
    lhs.extend(rhs)
    return lhs

TMS_DATA = dict(\
    merge(tms_loa.data(),
    merge(tms_60.data(),
    tms_lifeissues.data()
    ))
)

TMS_KEYS = \
    merge(tms_loa.keys(),
    merge(tms_60.keys(),
    tms_lifeissues.keys()
    ))

TMS_ALIAS = dict(\
    merge(tms_loa.aliases(),
    merge(tms_60.aliases(),
    tms_lifeissues.aliases()
    ))
)

TMS_NAMES = dict(\
    merge(tms_loa.names(),
    merge(tms_60.names(),
    tms_lifeissues.names()
    ))
)


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