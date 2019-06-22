# coding=utf-8


class BLBPassage():
    def __init__(self, ref, ver, txt, lex):
        self.reference = ref
        self.version = ver
        self.text = txt
        self.strongs = lex

    def get_reference(self):
        return self.reference

    def get_version(self):
        return self.version

    def get_text(self):
        return self.text

    def get_strongs(self):
        return self.strongs
