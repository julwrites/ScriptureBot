# coding=utf-8

# Local modules
from common import debug, text_utils
from common.action import action_classes
from common.telegram import telegram_utils

from bible import bible_utils

PASSAGE_PROMPT = "Give me a Bible reference"
SEARCH_PROMPT = "Give me a word or phrase to search"
STRONGS_PROMPT = "Give me a Strongs Lexicon reference"

BADQUERY = [
    "Sorry {}, I can't find it", "I don't know what that is, {}",
    "Please forgive my ignorance, {}", "What is that, {}?"
]


class BiblePassageAction(action_classes.Action):
    def identifier(self):
        return "/passage"

    def name(self):
        return "Bible Passage"

    def description(self):
        return "Search for a passage of Scripture"

    def match(self, msg):
        return msg is not None

    def resolve(self, userObj, msg):
        query = telegram_utils.strip_command(msg, self.identifier())

        if text_utils.is_valid(query):
            passage = bible_utils.get_passage(query, userObj.get_version())

            if passage is not None:
                telegram_utils.send_msg(user=userObj.get_uid(), text=passage)
                userObj.set_state(None)
            elif self.waiting(userObj):
                telegram_utils.send_msg(
                    user=userObj.get_uid(),
                    text=userObj.get_reply_string(BADQUERY))
            else:
                return False
        else:
            telegram_utils.send_msg(
                user=userObj.get_uid(), text=PASSAGE_PROMPT)

            userObj.set_state(self.identifier())

        return True


class BibleSearchAction(action_classes.Action):
    def identifier(self):
        return "/search"

    def name(self):
        return "Bible Search"

    def description(self):
        return "Search for a word or phrase in the Bible"

    def resolve(self, userObj, msg):
        query = telegram_utils.strip_command(msg, self.identifier())

        if text_utils.is_valid(query):
            passage = bible_utils.get_search(query, userObj.get_version())

            if passage is not None:
                telegram_utils.send_msg(user=userObj.get_uid(), text=passage)
                userObj.set_state(None)
            elif self.waiting(userObj):
                telegram_utils.send_msg(
                    user=userObj.get_uid(),
                    text=userObj.get_reply_string(BADQUERY))
            else:
                return False
        else:
            telegram_utils.send_msg(user=userObj.get_uid(), text=SEARCH_PROMPT)

            userObj.set_state(self.identifier())

        return True


class BiblePassageStrongsAction(action_classes.Action):
    def identifier(self):
        return "/strongspassage"

    def name(self):
        return "Bible Passage with Strongs"

    def description(self):
        return "Retrieve a passage with Strongs references"

    def resolve(self, userObj, msg):
        query = telegram_utils.strip_command(msg, self.identifier())

        if text_utils.is_valid(query):
            passage = bible_utils.get_passage_strongs(query,
                                                      userObj.get_version())

            if passage is not None:
                telegram_utils.send_msg(user=userObj.get_uid(), text=passage)
                userObj.set_state(None)
            elif self.waiting(userObj):
                telegram_utils.send_msg(
                    user=userObj.get_uid(),
                    text=userObj.get_reply_string(BADQUERY))
            else:
                return False
        else:
            telegram_utils.send_msg(
                user=userObj.get_uid(), text=PASSAGE_PROMPT)

            userObj.set_state(self.identifier())

        return True


class BiblePassageStrongsAction(action_classes.Action):
    def identifier(self):
        return "/strongs"

    def name(self):
        return "Strongs Lexicon entry"

    def description(self):
        return "Retrieve the Strongs Lexicon entry"

    def resolve(self, userObj, msg):
        query = telegram_utils.strip_command(msg, self.identifier())

        if text_utils.is_valid(query):
            link = bible_utils.get_strongs_entry(query)

            if link is not None:
                telegram_utils.send_msg(user=userObj.get_uid(), text=passage)
                userObj.set_state(None)
            elif self.waiting(userObj):
                telegram_utils.send_msg(
                    user=userObj.get_uid(),
                    text=userObj.get_reply_string(BADQUERY))
            else:
                return False
        else:
            telegram_utils.send_msg(
                user=userObj.get_uid(), text=STRONGS_PROMPT)

            userObj.set_state(self.identifier())

        return True


def get():
    return [
        BiblePassageAction(),
        BibleSearchAction(), BiblePassageStrongsAction,
        BibleStrongsAction()
    ]
