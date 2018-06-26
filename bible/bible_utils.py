# coding=utf-8

# Local modules
from bible import bgw_utils, blb_utils
from common.telegram import telegram_utils


def get_passage_raw(ref, version="NIV"):
    return bgw_utils.get_passage_raw(ref, version)


def get_passage(ref, version="NIV"):
    return bgw_utils.get_passage(ref, version)


def get_reference(query):
    return bgw_utils.get_reference(query)


def get_link(query):
    return bgw_utils.get_link(query)


def get_versions():
    return bgw_utils.get_versions()


def get_search(query, version="NASB"):
    return blb_utils.get_search(query, version)