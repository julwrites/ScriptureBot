
from tms.tms_classes import TMSPack

from tms import tms_loa
from tms import tms_60
from tms import tms_lifeissues

TMS = TMSPack(top=tms_loa.top())    \
    .add(tms_loa.pack())            \
    .add(tms_60.pack())             \
    .add(tms_lifeissues.pack())

def get_keys():
    return TMS.get_keys()

def get_data():
    return TMS.get_data()

def get_names():
    return TMS.get_names()

def get_aliases():
    return TMS.get_aliases()

def get_top():
    return TMS.get_top()