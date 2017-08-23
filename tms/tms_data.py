
BWC_KEY = "BWC"
BWC = [
    ("Assurance of Salvation",          "1 John 5 : 11 - 12"),
    ("Assurance of Answered Prayer",    "John 16 : 24"),
    ("Assurance of Victory",            "1 Corinthians 10 : 13"),
    ("Assurance of Forgiveness",        "1 John 1 : 9"),
    ("Assurance of Guidance",           "Proverbs 3 : 5 - 6")
]

A_KEY = "A"
A_PACK = [
    ("Christ the Center",               "2 Corinthians 5 : 17"),
    ("Christ the Center",               "Galatians 2 : 20"),
    ("Obedience to Christ",             "Romans 12 : 1"),
    ("Obedience to Christ",             "John 14 : 21"),
    ("The Word",                        "2 Timothy 3 : 16"),
    ("The Word",                        "Joshua 1 : 8"),
    ("Prayer",                          "John 15 : 7"),
    ("Prayer",                          "Philippians 4 : 6 - 7"),
    ("Fellowship",                      "Matthew 18 : 20"),
    ("Fellowship",                      "Hebrews 10 : 24 - 25"),
    ("Witnessing",                      "Matthew 4 : 19"),
    ("Witnessing",                      "Romans 1 : 16")
]

B_KEY = "B"
B_PACK = [
    ("All have Sinned",                 "Romans 3 : 23"),
    ("All have Sinned",                 "Isaiah 53 : 6"),
    ("Sin's Penalty",                   "Romans 6 : 23"),
    ("Sin's Penalty",                   "Hebrews 9 : 27"),
    ("Christ paid the Penalty",         "Romans 5 : 8"),
    ("Christ paid the Penalty",         "1 Peter 3 : 18"),
    ("Salvation not by Works",          "Ephesians 2 : 8 - 9"),
    ("Salvation not by Works",          "Titus 3 : 5"),
    ("Must receive Christ",             "John 1 : 12"),
    ("Must receive Christ",             "Revelation 3 : 20"),
    ("Assurance of Salvation",          "1 John 5 : 13"),
    ("Assurance of Salvation",          "John 5 : 24")
]

C_KEY = "C"
C_PACK = [
    ("His Spirit",                      "1 Corinthians 3 : 16"),
    ("His Spirit",                      "1 Corinthians 2 : 12"),
    ("His Strength",                    "Isaiah 41 : 10"),
    ("His Strength",                    "Philippians 4 : 13"),
    ("His Faithfulness",                "Lamentations 3 : 22 - 23"),
    ("His Faithfulness",                "Numbers 23 : 19"),
    ("His Peace",                       "Isaiah 26 : 3"),
    ("His Peace",                       "1 Peter 5 : 7"),
    ("His Provision",                   "Romans 8 : 32"),
    ("His Provision",                   "Philippians 4 : 19"),
    ("His Help in Temptation",          "Hebrews 2 : 18"),
    ("His Help in Temptation",          "Psalms 119 : 9, 11")
]

D_KEY = "D"
D_PACK = [
    ("Put Christ First",                "Matthew 6 : 33"),
    ("Put Christ First",                "Luke 9 : 23"),
    ("Separate from the World",         "1 John 2 : 15 - 16"),
    ("Separate from the World",         "Romans 12 : 2"),
    ("Be Steadfast",                    "1 Corinthians 15 : 58"),
    ("Be Steadfast",                    "Hebrews 12 : 3"),
    ("Serve Others",                    "Mark 10 : 45"),
    ("Serve Others",                    "2 Corinthians 4 : 5"),
    ("Give Generously",                 "Proverbs 3 : 9 - 10"),
    ("Give Generously",                 "2 Corinthians 9 : 6 - 7"),
    ("Develop World Vision",            "Acts 1 : 8"),
    ("Develop World Vision",            "Matthew 28 : 19 - 20")
]

E_KEY = "E"
E_PACK = [
    ("Love",                            "John 13 : 34 - 35"),
    ("Love",                            "1 John 3 : 18"),
    ("Humility",                        "Philippians 2 : 3 - 4"),
    ("Humility",                        "1 Peter 5 : 5 - 6"),
    ("Purity",                          "Ephesians 5 : 3"),
    ("Purity",                          "1 Peter 2 : 11"),
    ("Honesty",                         "Leviticus 19 : 11"),
    ("Honesty",                         "Acts 24 : 16"),
    ("Faith",                           "Hebrews 11 : 6"),
    ("Faith",                           "Romans 4 : 20 - 21"),
    ("Good Works",                      "Galatians 6 : 9 - 10"),
    ("Good Works",                      "Matthew 5 : 16")
]

TMS_DATA = {
    BWC_KEY :   BWC,
    A_KEY :     A_PACK,
    B_KEY :     B_PACK,
    C_KEY :     C_PACK,
    D_KEY :     D_PACK,
    E_KEY :     E_PACK
}

TMS_KEYS = [
    BWC_KEY,
    A_KEY,
    B_KEY,
    C_KEY,
    D_KEY,
    E_KEY
]

TMS_ALIAS = {
    BWC_KEY :   ["BWC", "LOA", "Beginning with Christ", "Lessons of Assurance"],
    A_KEY :     ["A", "A Pack", "TMS A"],
    B_KEY :     ["B", "B Pack", "TMS B"],
    C_KEY :     ["C", "C Pack", "TMS C"],
    D_KEY :     ["D", "D Pack", "TMS D"]
}

TMS_NAME = {
    BWC_KEY :   "Beginning with Christ/Lessons of Assurance",
    A_KEY :     "A",
    B_KEY :     "B",
    C_KEY :     "C",
    D_KEY :     "D",
    E_KEY :     "E"
}

def get_keys():
    return TMS_KEYS

def get_data():
    return TMS_DATA

def get_aliases():
    return TMS_ALIAS

def get_names():
    return TMS_NAMES

def get_top():
    return BWC_KEY