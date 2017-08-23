
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
    if pack is not None:
        select_pack = tms_data.get_data().get(pack)

        if select_pack is not None:
            return select_pack
    return None

def get_alias(pack):
    if pack is not None:
        select_alias = tms_data.get_aliases().get(pack)

        if select_alias is not None:
            return select_alias
    return None

def get_name(pack):
    if pack is not None:
        select_name = tms_data.get_names().get(pack)

        if select_name is not None:
            return select_name
    return None

def get_all_pack_keys():
    return tms_data.get_keys()

def query_pack_by_alias(query):
    if query is not None:
        stripped_query = text_utils.strip_numbers(query)
        for pack_key in get_all_pack_keys():
            aliases = get_alias(pack_key)
            for alias in aliases:
                if text_utils.fuzzy_compare(stripped_query, alias):
                    return pack_key
    return None

def query_verse_by_pack_pos(query):
    if query is not None:
        pack_key = query_pack_by_alias(query)
        if pack_key is not None:
            pack_name = get_name(pack_key)
            select_pack = get_pack(pack_key)
            if select_pack is not None:
                size = len(select_pack)
                pos = int(text_utils.strip_alpha(query))
                if size >= pos:
                    select_verse = select_pack[pos - 1]
                    return Verse(select_verse[1], select_verse[0], pack_name, pos)
    return None

def query_verse_by_reference(ref):
    if ref is not None:
        for pack_key in get_all_pack_keys():
            select_pack = get_pack(pack_key)
            pack_name = get_name(pack_key)
            size = len(select_pack)
            for i in range(0, size):
                select_verse = select_pack[i]
                if text_utils.fuzzy_compare(select_verse[1], ref):
                    return Verse(select_verse[1], select_verse[0], pack_name, i + 1)
    return None
   
def get_verse_by_pack(pack, pos):
    if pack is not None and pack is not None:
        pack_key = query_pack_by_alias(pack)
        if pack_key is not None:
            pack_name = get_name(pack_key)
            select_pack = get_pack(pack_key)

            if select_pack is not None:
                select_verse = select_pack[pos - 1]

                if select_verse is not None:
                    return Verse(select_verse[1], select_verse[0], pack_name, pos)
    return None

def get_verse_by_title(title, pos):
    if title is not None and pos is not None:
        verses = get_verses_by_title(title)

        if len(verses) > pos:
            return verses[pos - 1]
    return None

def get_verses_by_title(title):
    if title is not None:
        verses = []
        for pack_key in get_all_pack_keys():
            pack_name = get_name(pack_key)
            select_pack = get_pack(pack_key)
            size = len(select_pack)
            for i in range(0, size):
                select_verse = select_pack[i]
                if text_utils.fuzzy_compare(title, select_verse[0]):
                    verses.append(Verse(select_verse[1], select_verse[0], pack_name, i + 1))
        
        return verses
    return None

def get_start_verse():
    start_key = tms_data.get_top()
    pack_name = get_name(start_key)
    select_pack = get_pack(start_key)
    select_verse = select_pack[0]
    return Verse(select_verse[1], select_verse[0], pack_name, 1)

def format_verse(verse, passage):
    if verse is not None and passage is not None:
        verse_prep = []

        verse_prep.append(verse.get_pack() + ' ' + str(verse.get_position()))
        verse_prep.append(telegram_utils.bold(verse.reference) + ' ' \
                        + telegram_utils.bracket(passage.get_version()))
        verse_prep.append(passage.get_text())
        verse_prep.append(telegram_utils.bold(verse.reference))

        return telegram_utils.join(verse_prep, '\n\n')
    return None
