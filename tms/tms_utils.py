
# Python modules
import random

# Local modules
from common import debug, text_utils
from common.telegram import telegram_utils

from tms import tms_data


# Gettor functions: These return some kind of data with no fuzziness
def get_pack(packKey):
    if text_utils.is_valid(packKey):
        selectPack = tms_data.get_data().get(packKey)

        if selectPack is not None:
            return selectPack
    return None

def get_aliases(packKey):
    if text_utils.is_valid(packKey):
        selectAliases = tms_data.get_aliases().get(packKey)

        if selectAliases is not None:
            return selectAliases
    return None

def get_names(packKey):
    if text_utils.is_valid(packKey):
        selectNames = tms_data.get_names().get(packKey)

        if selectNames is not None:
            return selectNames
    return None

def get_all_pack_keys():
    return tms_data.get_keys()

def get_verse_by_pack_pos(packKey, pos):
    if text_utils.is_valid(packKey) and pos is not None:
        selectPack = get_pack(packKey)

        if selectPack is not None:
            selectVerse = selectPack[pos - 1]

            if selectVerse is not None:
                return selectVerse
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

        for packKey in get_all_pack_keys():
            selectPack = get_pack(packKey)
            size = len(selectPack)

            for i in range(0, size):
                selectVerse = selectPack[i]
                if text_utils.text_compare(title, selectVerse.get_title()):
                    verses.append(selectVerse)
        
        return verses
    return None

def get_start_verse():
    startKey = tms_data.get_top()
    selectPack = get_pack(startKey)
    selectVerse = selectPack[0]
    return selectVerse

def get_random_verse():
    packKeys = get_all_pack_keys()
    numPacks = len(packKeys)

    if numPacks > 0:
        choose = random.randint(0, numPacks - 1)
        selectPack = get_pack(packKeys[choose])
        numVerses = len(selectPack)

        if numVerses > 0:
            choose = random.randint(0, numVerses - 1)
            return selectPack[choose]


# Querying functions: These do a lookup based on some text search
def query_pack_by_alias(query):
    if text_utils.is_valid(query):
        for packKey in get_all_pack_keys():
            aliases = get_aliases(packKey)

            for alias in aliases:

                if text_utils.text_compare(query, alias):
                    return packKey
    return None

def query_verse_by_pack_pos(query):
    if text_utils.is_valid(query):
        debug.log("Attempting to get by position: " + query)
        queryNum = text_utils.strip_alpha(query)
        queryText = text_utils.strip_numbers(query)

        if text_utils.is_valid(queryText) and text_utils.is_valid(queryNum):
            packKey = query_pack_by_alias(queryText)

            if packKey is not None:
                selectPack = get_pack(packKey)

                if selectPack is not None:
                    size = len(selectPack)
                    pos = int(queryNum)

                    if size >= pos:
                        return selectPack[pos - 1]
    return None

def query_verse_by_reference(query):
    if text_utils.is_valid(query):
        debug.log("Attempting to get by reference " + query)

        for packKey in get_all_pack_keys():
            selectPack = get_pack(packKey)
            size = len(selectPack)

            for i in range(0, size):
                selectVerse = selectPack[i]

                if text_utils.text_compare(query, selectVerse.get_reference()):
                    return selectVerse
    return None

def query_verse_by_topic(query):
    if text_utils.is_valid(query):
        debug.log("Attempting to get by topic " + query)
        query = text_utils.strip_numbers(query)
        shortlist = []

        for packKey in get_all_pack_keys():
            pack = get_pack(packKey)
            addPack = False

            debug.log('Check alias for ' + packKey)
            for alias in get_aliases(packKey):

                if text_utils.fuzzy_compare(query, alias):
                    shortlist.extend(pack)
                    addPack = True
                    break

            if not addPack:
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
        versePrep = []

        versePrep.append(get_names(verse.get_pack()) + ' ' + str(verse.get_position()))
        versePrep.append(telegram_utils.bold(verse.get_title()))
        versePrep.append(telegram_utils.bold(verse.reference) + ' ' \
                        + telegram_utils.bracket(passage.get_version()))
        verseprep.append(passage.get_text())
        versePrep.append(telegram_utils.bold(verse.reference))

        return telegram_utils.join(versePrep, '\n\n')
    return None
