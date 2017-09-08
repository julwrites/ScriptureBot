
# coding=utf-8

class BGWPassage():
    def __init__(self, ref, ver, txt):
        self.reference = ref
        self.version = ver
        self.text = txt

    def get_reference(self):
        return self.reference

    def get_version(self):
        return self.version

    def get_text(self):
        return self.text

