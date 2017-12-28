# coding=utf-8

# Local modules
from bible import bgw_utils
from common.telegram import telegram_utils


def fetch_passage_html(ref, version="NIV"):
    return bgw_utils.fetch_bgw(ref, version)


def get_passage_raw(ref, version="NIV"):
    return bgw_utils.get_passage_raw(ref, version)


def get_passage(ref, version="NIV", ref_md=telegram_utils.bold, text_md=None):
    return bgw_utils.get_passage(ref, version, ref_md, text_md)


def get_reference(query):
    return bgw_utils.get_reference(query)


def get_link(query):
    return bgw_utils.get_link(query)


def get_versions():
    return bgw_utils.get_versions()