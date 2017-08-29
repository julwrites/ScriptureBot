
# Local modules
from bible import bgw_utils

def get_passage_raw(ref, version='NIV'):
    return bgw_utils.get_passage_raw(ref, version)

def get_passage(ref, version='NIV'):
    return bgw_utils.get_passage(ref, version)

def get_reference(query):
    return bgw_utils.get_reference(query)