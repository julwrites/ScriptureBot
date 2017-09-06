
from tms.tms_classes import TMSPack

from tms import tms_data_loa
from tms import tms_data_60
from tms import tms_data_lifeissues

TMS = TMSPack(top=tms_data_loa.top())    \
    .add(tms_data_loa.pack())            \
    .add(tms_data_60.pack())             \
    .add(tms_data_lifeissues.pack())

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