
# Local modules
from common import telegram_utils

from tms.tms_data import *

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
    select_pack = TMS.get(pack)

    if select_pack is not None:
        return select_pack

    return None

def get_verse_by_pack(pack, pos):
    select_pack = get_pack(pack)

    if select_pack is not None:
        select_verse = select_pack[pos - 1]

        if select_verse is not None:
            return Verse(select_verse[1], select_verse[0], pack, pos)

    return None

def get_verse_by_title(title, pos):
    verses = get_verses_by_title(title)

    if len(verses) > pos:
        return verses[pos - 1]

    return None

def get_verses_by_title(title):
    verses = []

    for pack_key in TMS.keys():
        pack = TMS.get(pack_key)
        size = len(pack)
        for i in range(0, size):
            if title == pack[i][0]:
                select_verse = pack[i]
                verse = Verse(select_verse[1], select_verse[0], pack_key, i + 1)
                verses.append(verse)

    return verses

def get_start_verse():
    start_key = 'BWC'
    select_pack = TMS.get(start_key)
    select_verse = select_pack[0]
    return Verse(select_verse[1], select_verse[0], start_key, 1)

def format_verse(verse, text):
    verse_prep = []

    verse_prep.append(verse.get_pack() + ' ' + str(verse.get_position()))
    verse_prep.append(text)
    verse_prep.append(telegram_utils.bold(verse.reference))

    return telegram_utils.join(verse_prep, '\n\n')

