
# Python modules
import random

# Local modules
from common import debug, text_utils
from common.telegram import telegram_utils

from tms import tms_data


# Gettor functions: These return some kind of data with no fuzziness
def get_pack(pack_key):
    if text_utils.is_valid(pack_key):
        select_pack = tms_data.get_data().get(pack_key)

        if select_pack is not None:
            return select_pack
    return None

def get_aliases(pack_key):
    if text_utils.is_valid(pack_key):
        select_aliases = tms_data.get_aliases().get(pack_key)

        if select_aliases is not None:
            return select_aliases
    return None

def get_names(pack_key):
    if text_utils.is_valid(pack_key):
        select_names = tms_data.get_names().get(pack_key)

        if select_names is not None:
            return select_names
    return None

def get_all_pack_keys():
    return tms_data.get_keys()

def get_verse_by_pack_pos(pack_key, pos):
    if text_utils.is_valid(pack_key) and pos is not None:
        select_pack = get_pack(pack_key)

        if select_pack is not None:
            select_verse = select_pack[pos - 1]

            if select_verse is not None:
                return select_verse
    return None

def get_verse_by_title(title, pos):
    if text_utils.is_valid(title) and pos is not None:
        verses = get_verses_by_title(title)

        if len(verses) > pos:
            return verses[pos - 1]
    return None

def get_verses_by_title(title):
    if text_utils.is_valid(title):
        verses = []

        for pack_key in get_all_pack_keys():
            select_pack = get_pack(pack_key)
            size = len(select_pack)

            for i in range(0, size):
                select_verse = select_pack[i]
                if text_utils.text_compare(title, select_verse.get_title()):
                    verses.append(select_verse)
        
        return verses
    return None

def get_start_verse():
    start_key = tms_data.get_top()
    select_pack = get_pack(start_key)
    select_verse = select_pack[0]
    return select_verse

def get_random_verse():
    pack_keys = get_all_pack_keys()
    num_packs = len(pack_keys)

    if num_packs > 0:
        choose = random.randint(0, num_packs - 1)
        select_pack = get_pack(pack_keys[choose])
        num_verses = len(select_pack)

        if num_verses > 0:
            choose = random.randint(0, num_verses - 1)
            return select_pack[choose]


# Querying functions: These do a lookup based on some text search
def query_pack_by_alias(query):
    if text_utils.is_valid(query):
        for pack_key in get_all_pack_keys():
            aliases = get_aliases(pack_key)

            for alias in aliases:

                if text_utils.text_compare(query, alias):
                    return pack_key
    return None

def query_verse_by_pack_pos(query):
    if text_utils.is_valid(query):
        debug.log("Attempting to get by position: " + query)
        query_num = text_utils.strip_alpha(query)
        query_text = text_utils.strip_numbers(query)

        if text_utils.is_valid(query_text) and text_utils.is_valid(query_num):
            pack_key = query_pack_by_alias(query_text)

            if pack_key is not None:
                select_pack = get_pack(pack_key)

                if select_pack is not None:
                    size = len(select_pack)
                    pos = int(query_num)

                    if size >= pos:
                        return select_pack[pos - 1]
    return None

def query_verse_by_reference(query):
    if text_utils.is_valid(query):
        debug.log("Attempting to get by reference " + query)

        for pack_key in get_all_pack_keys():
            select_pack = get_pack(pack_key)
            size = len(select_pack)

            for i in range(0, size):
                select_verse = select_pack[i]

                if text_utils.text_compare(query, select_verse.get_reference()):
                    return select_verse
    return None

def query_verse_by_topic(query):
    if text_utils.is_valid(query):
        debug.log("Attempting to get by topic " + query)
        query = text_utils.strip_numbers(query)
        shortlist = []

        for pack_key in get_all_pack_keys():
            pack = get_pack(pack_key)
            add_pack = False

            debug.log('Check alias for ' + pack_key)
            for alias in get_aliases(pack_key):

                if text_utils.fuzzy_compare(query, alias):
                    shortlist.extend(pack)
                    add_pack = True
                    break

            if not add_pack:
                debug.log('Check verses for related topics')
                for verse in pack:

                    if text_utils.fuzzy_compare(query, verse.get_title()):
                        shortlist.append(verse)
                    else:
                        for topic in verse.get_topics():
                            if text_utils.fuzzy_compare(query, topic):
                                shortlist.append(verse)

        debug.log('Found these related queries: ' + str(shortlist))
        num = len(shortlist)
        if num > 0:
            choose = random.randint(0, num - 1)
            return shortlist[choose]


# Formatting functions: These just do text manipulation and combination
def format_verse(verse, passage):
    if verse is not None and passage is not None:
        verse_prep = []

        verse_prep.append(get_names(verse.get_pack()) + ' ' + str(verse.get_position()))
        verse_prep.append(telegram_utils.bold(verse.get_title()))
        verse_prep.append(telegram_utils.bold(verse.reference) + ' ' \
                        + telegram_utils.bracket(passage.get_version()))
        verse_prep.append(passage.get_text())
        verse_prep.append(telegram_utils.bold(verse.reference))

        return telegram_utils.join(verse_prep, '\n\n')
    return None
