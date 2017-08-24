from tms.tms_class import TMSPack, TMSVerse


LOA_NUM = 0

LOA_KEY = "LOA"
def keys():
    return [ 
        LOA_KEY 
    ]

LOA = [
    TMSVerse(LOA_KEY,        "Assurance of Salvation",              1,        "1 John 5 : 11 - 12",
                            ['Assurance', 'Salvation', 'Saved', 'Eternal Life']),
    TMSVerse(LOA_KEY,        "Assurance of Answered Prayer",        2,        "John 16 : 24",
                            ['Assurance', 'Prayer', 'Pray', 'Praying', 'Jesus']),
    TMSVerse(LOA_KEY,        "Assurance of Victory",                3,        "1 Corinthians 10 : 13",
                            ['Assurance', 'Victory', 'Temptation', 'Tempted']),
    TMSVerse(LOA_KEY,        "Assurance of Forgiveness",            4,        "1 John 1 : 9",
                            ['Assurance', 'Sin', 'Confess', 'Forgive', 'Forgiveness']),
    TMSVerse(LOA_KEY,        "Assurance of Guidance",               5,        "Proverbs 3 : 5 - 6",
                            ['Assurance', 'Guidance', 'Guiding', 'Path'])
]
def data():
    return {
        LOA_KEY :       LOA,
    }

def names():
    return {
        LOA_KEY :       "Beginning with Christ/Lessons of Assurance",
    }

def aliases():
    return {
        LOA_KEY :       ["BWC", "LOA", "Beginning with Christ", "Lessons of Assurance"]
    }

def top():
    return LOA_KEY


LOA_PACK = TMSPack(keys(), data(), names(), aliases(), top())
def pack():
    return LOA_PACK