from secret import BOT_ID

APP_MAIN_URL = "/"
APP_HOOKS_URL = "/hooks"
APP_BOT_URL = "/" + BOT_ID

BOT_STATE_IDLE = 'Idle'

TELEGRAM_URL = 'https://api.telegram.org/bot' + BOT_ID
TELEGRAM_URL_SEND = TELEGRAM_URL + '/sendMessage'
TELEGRAM_MAX_LENGTH = 4096

JSON_HEADER = {'Content-Type': 'application/json;charset=utf-8'}

URL_TIMEOUT = 30