
# Python modules
import datetime

def now():
    date = (datetime.datetime.utcnow() + datetime.timedelta(hours=8)).date()
    time = datetime.datetime(
        date.year, 
        date.month, 
        date.day) - datetime.timedelta(hours=8)
    return time
