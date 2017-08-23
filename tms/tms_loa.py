from tms.tms_class import TMSPack, TMSVerse


LOA_NUM = 0

LOA_KEY = "LOA"
def keys():
    return [ 
        LOA_KEY 
    ]

LOA = [
    TMSVerse(LOA_KEY,		"Assurance of Salvation",		    1,		"1 John 5 : 11 - 12"),
    TMSVerse(LOA_KEY,		"Assurance of Answered Prayer",		2,		"John 16 : 24"),
    TMSVerse(LOA_KEY,		"Assurance of Victory",		        3,		"1 Corinthians 10 : 13"),
    TMSVerse(LOA_KEY,		"Assurance of Forgiveness",		    4,		"1 John 1 : 9"),
    TMSVerse(LOA_KEY,		"Assurance of Guidance",		    5,		"Proverbs 3 : 5 - 6")
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
        LOA_KEY :       ["BWC", "LOA", "Beginning with Christ", "Lessons of Assurance",
                        "New Christian", "Assurances", "Assurance", "Assure",
                        "Beginning", "Begin"],
    }

def top():
    return LOA_KEY


LOA_PACK = TMSPack(keys(), data(), names(), aliases(), top())
def pack():
    return LOA_PACK