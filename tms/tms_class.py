
class TMSVerse():
    def __init__(self, pack, title, pos, ref):
        self.pack = pack
        self.title = title
        self.position = pos
        self.reference = ref
    
    def get_reference(self):
        return self.reference

    def get_title(self):
        return self.title

    def get_pack(self):
        return self.pack

    def get_position(self):
        return self.position


class TMSPack():
    def __init__(self, keys=[], data={}, names={}, aliases={}, top=None):
        self.keys = keys
        self.data = data
        self.names = names
        self.aliases = aliases
        self.top = top

    def get_keys(self):
        return self.keys

    def get_data(self):
        return self.data

    def get_names(self):
        return self.names

    def get_aliases(self):
        return self.aliases

    def get_top(self):
        return self.top

    def add(self, pack):
        self.keys.extend(pack.keys)
        self.data.update(pack.data)
        self.names.update(pack.data)
        self.aliases.update(pack.data)
