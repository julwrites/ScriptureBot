from tms.tms_class import TMSPack


LOA_KEY = "LOA"
def keys():
    return [ 
        LOA_KEY 
    ]

LOA = [
    ("Assurance of Salvation",          "1 John 5 : 11 - 12"),
    ("Assurance of Answered Prayer",    "John 16 : 24"),
    ("Assurance of Victory",            "1 Corinthians 10 : 13"),
    ("Assurance of Forgiveness",        "1 John 1 : 9"),
    ("Assurance of Guidance",           "Proverbs 3 : 5 - 6")
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
        LOA_KEY :       ["BWC", "LOA", "Beginning with Christ", "Lessons of Assurance"],
    }

def top():
    return LOA_KEY


LOA_PACK = TMSPack(keys(), data(), names(), aliases(), top())
def pack():
    return LOA_PACK