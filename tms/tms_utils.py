
# Local modules
from common import debug
from common import telegram_utils
from common import text_utils

from tms import tms_data

class Verse():
    def __init__(self, ref, title, pack, pos):
        self.reference = ref
        self.title = title
        self.pack = pack
        self.position = pos
    
    def get_reference(self):
        return self.reference

    def get_title(self):
        return self.title

    def get_pack(self):
        return self.pack

    def get_position(self):
        return self.position

def get_pack(pack):
    select_pack = tms_data.get_tms().get(pack)

    if select_pack is not None:
        return select_pack

    return None

def get_all_pack_keys():
    return tms_data.get_tms().keys()

def query_verse_by_pack_pos(query):
    for pack_key in get_all_pack_keys():
        pack = get_pack(pack_key)
        size = len(pack)
        for i in range(0, size):
            select_verse = pack[i]
            try_packpos = pack_key + str(i + 1)
            if text_utils.fuzzy_compare(query, try_packpos):
                return Verse(select_verse[1], select_verse[0], pack_key, i + 1)
    return None

def query_verse_by_reference(ref):
    for pack_key in get_all_pack_keys():
        pack = get_pack(pack_key)
        size = len(pack)
        for i in range(0, size):
            select_verse = pack[i]
            if text_utils.fuzzy_compare(select_verse[1], ref):
                return Verse(select_verse[1], select_verse[0], pack_key, i + 1)
    return None
   
def get_verse_by_pack(pack, pos):
    select_pack = get_pack(pack)

    if select_pack is not None:
        select_verse = select_pack[pos - 1]

        if select_verse is not None:
            return Verse(select_verse[1], select_verse[0], pack, pos)

def get_verse_by_title(title, pos):
    verses = get_verses_by_title(title)

    if len(verses) > pos:
        return verses[pos - 1]

    return None
def get_verses_by_title(title):
    verses = []

    for pack_key in get_all_pack_keys():
        pack = get_pack(pack_key)
        size = len(pack)
        for i in range(0, size):
            select_verse = pack[i]
            if text_utils.fuzzy_compare(title, select_verse[0]):
                verses.append(Verse(select_verse[1], select_verse[0], pack_key, i + 1))
    
    return verses

def get_start_verse():
    start_key = 'BWC'
    select_pack = get_pack(start_key)
    select_verse = select_pack[0]
    return Verse(select_verse[1], select_verse[0], start_key, 1)

def format_verse(verse, passage):
    verse_prep = []

    verse_prep.append(verse.get_pack() + ' ' + str(verse.get_position()))
    verse_prep.append(telegram_utils.bold(verse.reference) + ' ' \
                    + telegram_utils.bracket(passage.get_version()))
    verse_prep.append(passage.get_text())
    verse_prep.append(telegram_utils.bold(verse.reference))

    return telegram_utils.join(verse_prep, '\n\n')

