

class TMSPack():
    def __init__(self):
        self.data = {}
        self.keys = []
        self.names = {}
        self.aliases = {}
        self.top = None

    def get_data(self):
        return self.data

    def get_keys(self):
        return self.keys

    def get_names(self):
        return self.names

    def get_aliases(self):
        return self.aliases

    def get_top(self):
        return self.top

    def join(self, pack):
        
