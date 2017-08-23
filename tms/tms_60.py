
from tms.tms_class import TMSPack, TMSVerse

TMS_A_KEY = "TMS_A"
TMS_B_KEY = "TMS_B"
TMS_C_KEY = "TMS_C"
TMS_D_KEY = "TMS_D"
TMS_E_KEY = "TMS_E"

def keys():
  return [ 
    TMS_A_KEY,
    TMS_B_KEY,
    TMS_C_KEY,
    TMS_D_KEY,
    TMS_E_KEY
  ]

A_PACK = [
  TMSVerse(TMS_A_KEY,    "Christ the Center",     1,      "2 Corinthians 5 : 17"),    
  TMSVerse(TMS_A_KEY,    "Christ the Center",     2,      "Galatians 2 : 20"),    
  TMSVerse(TMS_A_KEY,    "Obedience to Christ",   3,      "Romans 12 : 1"),    
  TMSVerse(TMS_A_KEY,    "Obedience to Christ",   4,      "John 14 : 21"),    
  TMSVerse(TMS_A_KEY,    "The Word",              5,      "2 Timothy 3 : 16"),    
  TMSVerse(TMS_A_KEY,    "The Word",              6,      "Joshua 1 : 8"),    
  TMSVerse(TMS_A_KEY,    "Prayer",                7,      "John 15 : 7"),    
  TMSVerse(TMS_A_KEY,    "Prayer",                8,      "Philippians 4 : 6 - 7"),    
  TMSVerse(TMS_A_KEY,    "Fellowship",            9,      "Matthew 18 : 20"),    
  TMSVerse(TMS_A_KEY,    "Fellowship",            10,     "Hebrews 10 : 24 - 25"),    
  TMSVerse(TMS_A_KEY,    "Witnessing",            11,     "Matthew 4 : 19"),    
  TMSVerse(TMS_A_KEY,    "Witnessing",            12,     "Romans 1 : 16")
]

B_PACK = [
  TMSVerse(TMS_B_KEY,    "All have Sinned",          1,     "Romans 3 : 23"),
  TMSVerse(TMS_B_KEY,    "All have Sinned",          2,     "Isaiah 53 : 6"),
  TMSVerse(TMS_B_KEY,    "Sin's Penalty",            3,     "Romans 6 : 23"),
  TMSVerse(TMS_B_KEY,    "Sin's Penalty",            4,     "Hebrews 9 : 27"),
  TMSVerse(TMS_B_KEY,    "Christ paid the Penalty",  5,     "Romans 5 : 8"),
  TMSVerse(TMS_B_KEY,    "Christ paid the Penalty",  6,     "1 Peter 3 : 18"),
  TMSVerse(TMS_B_KEY,    "Salvation not by Works",   7,     "Ephesians 2 : 8 - 9"),
  TMSVerse(TMS_B_KEY,    "Salvation not by Works",   8,     "Titus 3 : 5"),
  TMSVerse(TMS_B_KEY,    "Must receive Christ",      9,     "John 1 : 12"),
  TMSVerse(TMS_B_KEY,    "Must receive Christ",      10,    "Revelation 3 : 20"),
  TMSVerse(TMS_B_KEY,    "Assurance of Salvation",   11,    "1 John 5 : 13"),
  TMSVerse(TMS_B_KEY,    "Assurance of Salvation",   12,    "John 5 : 24")
]

C_PACK = [
  TMSVerse(TMS_C_KEY,    "His Spirit",               1,     "1 Corinthians 3 : 16"),
  TMSVerse(TMS_C_KEY,    "His Spirit",               2,     "1 Corinthians 2 : 12"),
  TMSVerse(TMS_C_KEY,    "His Strength",             3,     "Isaiah 41 : 10"),
  TMSVerse(TMS_C_KEY,    "His Strength",             4,     "Philippians 4 : 13"),
  TMSVerse(TMS_C_KEY,    "His Faithfulness",         5,     "Lamentations 3 : 22 - 23"),
  TMSVerse(TMS_C_KEY,    "His Faithfulness",         6,     "Numbers 23 : 19"),
  TMSVerse(TMS_C_KEY,    "His Peace",                7,     "Isaiah 26 : 3"),
  TMSVerse(TMS_C_KEY,    "His Peace",                8,     "1 Peter 5 : 7"),
  TMSVerse(TMS_C_KEY,    "His Provision",            9,     "Romans 8 : 32"),
  TMSVerse(TMS_C_KEY,    "His Provision",            10,    "Philippians 4 : 19"),
  TMSVerse(TMS_C_KEY,    "His Help in Temptation",   11,    "Hebrews 2 : 18"),
  TMSVerse(TMS_C_KEY,    "His Help in Temptation",   12,    "Psalms 119 : 9, 11")
]

D_PACK = [
  TMSVerse(TMS_D_KEY,    "Put Christ First",         1,     "Matthew 6 : 33"),
  TMSVerse(TMS_D_KEY,    "Put Christ First",         2,     "Luke 9 : 23"),
  TMSVerse(TMS_D_KEY,    "Separate from the World",  3,     "1 John 2 : 15 - 16"),
  TMSVerse(TMS_D_KEY,    "Separate from the World",  4,     "Romans 12 : 2"),
  TMSVerse(TMS_D_KEY,    "Be Steadfast",             5,     "1 Corinthians 15 : 58"),
  TMSVerse(TMS_D_KEY,    "Be Steadfast",             6,     "Hebrews 12 : 3"),
  TMSVerse(TMS_D_KEY,    "Serve Others",             7,     "Mark 10 : 45"),
  TMSVerse(TMS_D_KEY,    "Serve Others",             8,     "2 Corinthians 4 : 5"),
  TMSVerse(TMS_D_KEY,    "Give Generously",          9,     "Proverbs 3 : 9 - 10"),
  TMSVerse(TMS_D_KEY,    "Give Generously",          10,    "2 Corinthians 9 : 6 - 7"),
  TMSVerse(TMS_D_KEY,    "Develop World Vision",     11,    "Acts 1 : 8"),
  TMSVerse(TMS_D_KEY,    "Develop World Vision",     12,    "Matthew 28 : 19 - 20")
]

E_PACK = [
  TMSVerse(TMS_E_KEY,    "Love",                     1,     "John 13 : 34 - 35"),
  TMSVerse(TMS_E_KEY,    "Love",                     2,     "1 John 3 : 18"),
  TMSVerse(TMS_E_KEY,    "Humility",                 3,     "Philippians 2 : 3 - 4"),
  TMSVerse(TMS_E_KEY,    "Humility",                 4,     "1 Peter 5 : 5 - 6"),
  TMSVerse(TMS_E_KEY,    "Purity",                   5,     "Ephesians 5 : 3"),
  TMSVerse(TMS_E_KEY,    "Purity",                   6,     "1 Peter 2 : 11"),
  TMSVerse(TMS_E_KEY,    "Honesty",                  7,     "Leviticus 19 : 11"),
  TMSVerse(TMS_E_KEY,    "Honesty",                  8,     "Acts 24 : 16"),
  TMSVerse(TMS_E_KEY,    "Faith",                    9,     "Hebrews 11 : 6"),
  TMSVerse(TMS_E_KEY,    "Faith",                    10,    "Romans 4 : 20 - 21"),
  TMSVerse(TMS_E_KEY,    "Good Works",               11,    "Galatians 6 : 9 - 10"),
  TMSVerse(TMS_E_KEY,    "Good Works",               12,    "Matthew 5 : 16")
]

def data():
  return {
    TMS_A_KEY :     A_PACK,
    TMS_B_KEY :     B_PACK,
    TMS_C_KEY :     C_PACK,
    TMS_D_KEY :     D_PACK,
    TMS_E_KEY :     E_PACK
  }

def names():
  return {
    TMS_A_KEY :     "A: Living the New Life",
    TMS_B_KEY :     "B: Proclaiming Christ",
    TMS_C_KEY :     "C: Reliance on God's Resources",
    TMS_D_KEY :     "D: Being Christ's Disciple",
    TMS_E_KEY :     "E: Growth in Christlikeness"
  }

def aliases():
  return {
    TMS_A_KEY :     ["A", "A Pack", "TMS A", 
                    "Living the New Life", "New Life in Christ", "New Life", 
                    "Wheel", "Christian Living"],
    TMS_B_KEY :     ["B", "B Pack", "TMS B", 
                    "Proclaiming Christ", "Evangelism", 
                    "Bridge", "Bridge Illustration", "Evanglise", "Sharing Gospel",
                    "Gospel", "Gospel Presentation", "Present Gospel"],
    TMS_C_KEY :     ["C", "C Pack", "TMS C", 
                    "Reliance on God's Resources", "Reliance", "Empowering"],
    TMS_D_KEY :     ["D", "D Pack", "TMS D", 
                    "Being Christ's Disciple", "Discipleship", "Disciple",
                    "Discipling"],
    TMS_E_KEY :     ["E", "E Pack", "TMS E", 
                    "Growth in Christlikeness", "Christlikeness", "Christlike",
                    "Christ-like", "Christ Likeness"]
  }

def top():
  return TMS_A_KEY

TMS_PACK = TMSPack(keys(), data(), names(), aliases(), top())
def pack():
  return TMS_PACK