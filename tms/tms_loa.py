
LOA_KEY = "LOA"
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

def keys():
    return [ 
        LOA_KEY 
    ]

def aliases():
    return {
        LOA_KEY :       ["BWC", "LOA", "Beginning with Christ", "Lessons of Assurance"],
    }

def names():
    return {
        LOA_KEY :       "Beginning with Christ/Lessons of Assurance",
    }

def top():
    return LOA_KEY