
from tms.tms_class import TMSPack, TMSVerse

LIFE_ISSUES_ANGER_KEY = "Anger"
LIFE_ISSUES_SIN_KEY = "Sin"
LIFE_ISSUES_DEPRESSION_KEY = "Depression"
LIFE_ISSUES_GUILT_KEY = "Guilt"
LIFE_ISSUES_GODSWILL_KEY = "God's Will"
LIFE_ISSUES_LOVE_KEY = "Love"
LIFE_ISSUES_MONEY_KEY = "Money"
LIFE_ISSUES_PERFECTIONISM_KEY = "Perfectionism"
LIFE_ISSUES_SELFIMAGE_KEY = "Self-Image"
LIFE_ISSUES_SEX_KEY = "Sex"
LIFE_ISSUES_STRESS_KEY = "Stress"
LIFE_ISSUES_SUFFERING_KEY = "Suffering"

def keys():
    return [ 
        LIFE_ISSUES_ANGER_KEY,
        LIFE_ISSUES_SIN_KEY,
        LIFE_ISSUES_DEPRESSION_KEY,
        LIFE_ISSUES_GUILT_KEY,
        LIFE_ISSUES_GODSWILL_KEY,
        LIFE_ISSUES_LOVE_KEY,
        LIFE_ISSUES_MONEY_KEY,
        LIFE_ISSUES_PERFECTIONISM_KEY,
        LIFE_ISSUES_SELFIMAGE_KEY,
        LIFE_ISSUES_SEX_KEY,
        LIFE_ISSUES_STRESS_KEY,
        LIFE_ISSUES_SUFFERING_KEY
    ]

ANGER_PACK = [
    TMSVerse(LIFE_ISSUES_ANGER_KEY,		"Anger",        1,     "Proverbs 15 : 1"),
    TMSVerse(LIFE_ISSUES_ANGER_KEY,		"Anger",        2,     "Proverbs 29 : 11"),
    TMSVerse(LIFE_ISSUES_ANGER_KEY,		"Anger",        3,     "Romans 12 : 19"),
    TMSVerse(LIFE_ISSUES_ANGER_KEY,		"Anger",        4,     "Ephesians 4 : 26 - 27"),
    TMSVerse(LIFE_ISSUES_ANGER_KEY,		"Anger",        5,     "Colossians 3 : 8 - 10"),
    TMSVerse(LIFE_ISSUES_ANGER_KEY,		"Anger",        6,     "James 1 : 19 - 20")
]

SIN_PACK = [
    TMSVerse(LIFE_ISSUES_SIN_KEY,		"Sin",          1,     "Romans 6 : 11 - 13"),
    TMSVerse(LIFE_ISSUES_SIN_KEY,		"Sin",          2,     "1 Corinthians 10 : 13"),
    TMSVerse(LIFE_ISSUES_SIN_KEY,		"Sin",          3,     "Galatians 6 : 1 - 2"),
    TMSVerse(LIFE_ISSUES_SIN_KEY,		"Sin",          4,     "Ephesians 6 : 10 - 12"),
    TMSVerse(LIFE_ISSUES_SIN_KEY,		"Sin",          5,     "James 4 : 7 - 8"),
    TMSVerse(LIFE_ISSUES_SIN_KEY,		"Sin",          6,     "1 John 1 : 8"),
    TMSVerse(LIFE_ISSUES_SIN_KEY,		"Sin",          7,     "1 John 1 : 9")
]

DEPRESSION_PACK = [
    TMSVerse(LIFE_ISSUES_DEPRESSION_KEY,	"Depression",      1,   "Isaiah 43 : 1 - 3"),
    TMSVerse(LIFE_ISSUES_DEPRESSION_KEY,	"Depression",      2,   "2 Corinthians 4 : 7 - 10"),
    TMSVerse(LIFE_ISSUES_DEPRESSION_KEY,	"Depression",      3,   "Psalm 42 : 5"),
    TMSVerse(LIFE_ISSUES_DEPRESSION_KEY,	"Depression",      4,   "Psalm 34 : 17 - 18"),
    TMSVerse(LIFE_ISSUES_DEPRESSION_KEY,	"Depression",      5,   "Lamentations 3 : 19 - 23"),
    TMSVerse(LIFE_ISSUES_DEPRESSION_KEY,	"Depression",      6,   "2 Corinthians 1 : 8 - 9")
]

GUILT_PACK = [
    TMSVerse(LIFE_ISSUES_GUILT_KEY,		"Guilt",        1,     "Psalm 32 : 1 - 2"),
    TMSVerse(LIFE_ISSUES_GUILT_KEY,		"Guilt",        2,     "Psalm 51 : 9 - 10"),
    TMSVerse(LIFE_ISSUES_GUILT_KEY,		"Guilt",        3,     "Proverbs 28 : 13"),
    TMSVerse(LIFE_ISSUES_GUILT_KEY,		"Guilt",        4,     "Romans 8 : 1 - 2"),
    TMSVerse(LIFE_ISSUES_GUILT_KEY,		"Guilt",        5,     "2 Corinthians 7 : 10"),
    TMSVerse(LIFE_ISSUES_GUILT_KEY,		"Guilt",        6,     "James 5 : 16")
]

GODSWILL_PACK = [
    TMSVerse(LIFE_ISSUES_GODSWILL_KEY,		"God's Will",       1,  "Proverbs 3 : 5 - 6"),
    TMSVerse(LIFE_ISSUES_GODSWILL_KEY,		"God's Will",       2,  "Proverbs 3 : 7"),
    TMSVerse(LIFE_ISSUES_GODSWILL_KEY,		"God's Will",       3,  "Proverbs 16 : 9"),
    TMSVerse(LIFE_ISSUES_GODSWILL_KEY,		"God's Will",       4,  "Isaiah 30 : 21"),
    TMSVerse(LIFE_ISSUES_GODSWILL_KEY,		"God's Will",       5,  "Jeremiah 29 : 11 - 13"),
    TMSVerse(LIFE_ISSUES_GODSWILL_KEY,		"God's Will",       6,  "Romans 12 : 1 - 2"),
    TMSVerse(LIFE_ISSUES_GODSWILL_KEY,		"God's Will",       7,  "1 John 5 : 14 - 15")
]

LOVE_PACK = [
    TMSVerse(LIFE_ISSUES_LOVE_KEY,		"Love",         1,     "Matthew 22 : 37 - 40"),
    TMSVerse(LIFE_ISSUES_LOVE_KEY,		"Love",         2,     "John 13 : 34 - 35"),
    TMSVerse(LIFE_ISSUES_LOVE_KEY,		"Love",         3,     "Romans 8 : 38 - 39"),
    TMSVerse(LIFE_ISSUES_LOVE_KEY,		"Love",         4,     "1 Corinthians 13 : 1 - 3"),
    TMSVerse(LIFE_ISSUES_LOVE_KEY,		"Love",         5,     "1 Corinthians 13 : 4 - 8"),
    TMSVerse(LIFE_ISSUES_LOVE_KEY,		"Love",         6,     "1 John 4 : 20")
]

MONEY_PACK = [
    TMSVerse(LIFE_ISSUES_MONEY_KEY,		"Money",        1,     "Deuteronomy 8 : 17 - 18"),
    TMSVerse(LIFE_ISSUES_MONEY_KEY,		"Money",        2,     "Proverbs 3 : 9 - 10"),
    TMSVerse(LIFE_ISSUES_MONEY_KEY,		"Money",        3,     "Matthew 6 : 19 - 21"),
    TMSVerse(LIFE_ISSUES_MONEY_KEY,		"Money",        4,     "Matthew 6 : 24"),
    TMSVerse(LIFE_ISSUES_MONEY_KEY,		"Money",        5,     "Philippians 4 : 11 - 13"),
    TMSVerse(LIFE_ISSUES_MONEY_KEY,		"Money",        6,     "1 Timothy 6 : 9 - 10")
]

PERFECTIONISM_PACK = [
    TMSVerse(LIFE_ISSUES_PERFECTIONISM_KEY,		"Perfectionism",    1,  "Psalm 127 : 1 - 2"),
    TMSVerse(LIFE_ISSUES_PERFECTIONISM_KEY,		"Perfectionism",    2,  "Ecclesiastes 2 : 10 - 11"),
    TMSVerse(LIFE_ISSUES_PERFECTIONISM_KEY,		"Perfectionism",    3,  "Luke 10 : 40 - 42"),
    TMSVerse(LIFE_ISSUES_PERFECTIONISM_KEY,		"Perfectionism",    4,  "2 Corinthians 12 : 9"),
    TMSVerse(LIFE_ISSUES_PERFECTIONISM_KEY,		"Perfectionism",    5,  "Galatians 3 : 3"),
    TMSVerse(LIFE_ISSUES_PERFECTIONISM_KEY,		"Perfectionism",    6,  "Ephesians 2 : 8 - 9")
]

SELFIMAGE_PACK = [
    TMSVerse(LIFE_ISSUES_SELFIMAGE_KEY,		"Self-Image",       1,  "1 Samuel 16 : 7"),
    TMSVerse(LIFE_ISSUES_SELFIMAGE_KEY,		"Self-Image",       2,  "Psalm 139 : 13 - 14"),
    TMSVerse(LIFE_ISSUES_SELFIMAGE_KEY,		"Self-Image",       3,  "Jeremiah 9 : 23 - 24"),
    TMSVerse(LIFE_ISSUES_SELFIMAGE_KEY,		"Self-Image",       4,  "Matthew 10 : 29 - 31"),
    TMSVerse(LIFE_ISSUES_SELFIMAGE_KEY,		"Self-Image",       5,  "Philippians 2 : 3 - 11"),
    TMSVerse(LIFE_ISSUES_SELFIMAGE_KEY,		"Self-Image",       6,  "1 Peter 3 : 3 - 4")
]

SEX_PACK = [
    TMSVerse(LIFE_ISSUES_SEX_KEY,		"Sex",         1,      "Matthew 5 : 27 - 28"),
    TMSVerse(LIFE_ISSUES_SEX_KEY,		"Sex",         2,      "Romans 13 : 13 - 14"),
    TMSVerse(LIFE_ISSUES_SEX_KEY,		"Sex",         3,      "1 Corinthians 6 : 18 - 20"),
    TMSVerse(LIFE_ISSUES_SEX_KEY,		"Sex",         4,      "Ephesians 5 : 3"),
    TMSVerse(LIFE_ISSUES_SEX_KEY,		"Sex",         5,      "1 Thessalonians 4 : 3 - 5"),
    TMSVerse(LIFE_ISSUES_SEX_KEY,		"Sex",         6,      "Hebrews 13 : 4")
]

STRESS_PACK = [
    TMSVerse(LIFE_ISSUES_STRESS_KEY,		"Stress",           1,  "Psalm 73 : 26"),
    TMSVerse(LIFE_ISSUES_STRESS_KEY,		"Stress",           2,  "Psalm 118 : 5 - 6"),
    TMSVerse(LIFE_ISSUES_STRESS_KEY,		"Stress",           3,  "Matthew 11 : 28 - 30"),
    TMSVerse(LIFE_ISSUES_STRESS_KEY,		"Stress",           4,  "2 Corinthians 4 : 16 - 18"),
    TMSVerse(LIFE_ISSUES_STRESS_KEY,		"Stress",           5,  "Philippians 4 : 6 - 7"),
    TMSVerse(LIFE_ISSUES_STRESS_KEY,		"Stress",           6,  "1 Peter 5 : 5 - 7")
]

SUFFERING_PACK = [
    TMSVerse(LIFE_ISSUES_SUFFERING_KEY,		"Suffering",        1,  "Romans 5 : 2 - 5"),
    TMSVerse(LIFE_ISSUES_SUFFERING_KEY,		"Suffering",        2,  "2 Corinthians 1 : 3 - 4"),
    TMSVerse(LIFE_ISSUES_SUFFERING_KEY,		"Suffering",        3,  "James 1 : 2 - 4"),
    TMSVerse(LIFE_ISSUES_SUFFERING_KEY,		"Suffering",        4,  "James 1 : 12"),
    TMSVerse(LIFE_ISSUES_SUFFERING_KEY,		"Suffering",        5,  "1 Peter 1 : 6 - 7"),
    TMSVerse(LIFE_ISSUES_SUFFERING_KEY,		"Suffering",        6,  "1 Peter 4 : 12 - 13")
]


def data():
    return {
        LIFE_ISSUES_ANGER_KEY :             ANGER_PACK,
        LIFE_ISSUES_SIN_KEY :               SIN_PACK,
        LIFE_ISSUES_DEPRESSION_KEY :        DEPRESSION_PACK,
        LIFE_ISSUES_GUILT_KEY :             GUILT_PACK,
        LIFE_ISSUES_GODSWILL_KEY :          GODSWILL_PACK,
        LIFE_ISSUES_LOVE_KEY :              LOVE_PACK,
        LIFE_ISSUES_MONEY_KEY :             MONEY_PACK,
        LIFE_ISSUES_PERFECTIONISM_KEY :     PERFECTIONISM_PACK,
        LIFE_ISSUES_SELFIMAGE_KEY :         SELFIMAGE_PACK,
        LIFE_ISSUES_SEX_KEY :               SEX_PACK,
        LIFE_ISSUES_STRESS_KEY :            STRESS_PACK,
        LIFE_ISSUES_SUFFERING_KEY :         SUFFERING_PACK
    }

def names():
    return {
        LIFE_ISSUES_ANGER_KEY :             "Life Issues: Anger",
        LIFE_ISSUES_SIN_KEY :               "Life Issues: Sin",
        LIFE_ISSUES_DEPRESSION_KEY :        "Life Issues: Depression",
        LIFE_ISSUES_GUILT_KEY :             "Life Issues: Guilt",
        LIFE_ISSUES_GODSWILL_KEY :          "Life Issues: God's Will",
        LIFE_ISSUES_LOVE_KEY :              "Life Issues: Love",
        LIFE_ISSUES_MONEY_KEY :             "Life Issues: Money",
        LIFE_ISSUES_PERFECTIONISM_KEY :     "Life Issues: Perfectionism",
        LIFE_ISSUES_SELFIMAGE_KEY :         "Life Issues: Self-Image",
        LIFE_ISSUES_SEX_KEY :               "Life Issues: Sex",
        LIFE_ISSUES_STRESS_KEY :            "Life Issues: Stress",
        LIFE_ISSUES_SUFFERING_KEY :         "Life Issues: Suffering"
    }

def aliases():
    return {
        LIFE_ISSUES_ANGER_KEY :             ["Anger", "Rage", "Wrath", "Frustration"],
        LIFE_ISSUES_SIN_KEY :               ["Sin", "Wrongdoing", "Wrong"],
        LIFE_ISSUES_DEPRESSION_KEY :        ["Depression", "Sadness"],
        LIFE_ISSUES_GUILT_KEY :             ["Guilt"],
        LIFE_ISSUES_GODSWILL_KEY :          ["God's Will", "Will of God", "Will"],
        LIFE_ISSUES_LOVE_KEY :              ["Love", "Storge", "Phileo", "Agape", "Eros"],
        LIFE_ISSUES_MONEY_KEY :             ["Money", "Rich", "Wealth"],
        LIFE_ISSUES_PERFECTIONISM_KEY :     ["Perfectionism", "Perfect"],
        LIFE_ISSUES_SELFIMAGE_KEY :         ["Self-Image", "Self Image"],
        LIFE_ISSUES_SEX_KEY :               ["Sex", "Purity"],
        LIFE_ISSUES_STRESS_KEY :            ["Stress", "Peace"],
        LIFE_ISSUES_SUFFERING_KEY :         ["Suffering", "Pain", "Enduring Pain"]
    }

def top():
    return LIFE_ISSUES_ANGER_KEY

LIFE_ISSUES_PACK = TMSPack(keys(), data(), names(), aliases(), top())
def pack():
    return LIFE_ISSUES_PACK