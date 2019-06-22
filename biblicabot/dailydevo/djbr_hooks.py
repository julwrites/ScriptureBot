# coding=utf-8

# Local modules
from common.utils import debug_utils
from common.telegram import telegram_utils
from common.action import hook_classes

from dailydevo import djbr_utils
from user import user_actions

PROMPT = "Here are today's Bible Reading passages!\n{}\nTap on any one to get the passage!"

REFLECTION = [
    "Take this time to reflect over the last few days. How has God been speaking to you, {}?",
    "As you read this, slow down and take a few deep breaths. Meditate on the living Word that God has been giving you, {}.",
    "{}! Stop right now and ask the Holy Spirit to remind you of what God has been speaking to your heart this week",
    "He hears you, {}. The question now is whether you will listen to Him. Spend some time alone with Him today.",
    "Jesus invites all to come to Him, and He promises to give rest for your soul. Yes, you, {}",
    "We know He speaks, but do you listen? Open your heart to Him today, {}, and believe He will speak",
    "{}, have faith that He will speak as you meditate; to come to Him you must have faith.",
]


class DJBRDailyHook(hook_classes.Hook):
    def identifier(self):
        return "/djbr"

    def name(self):
        return "Discipleship Journal Bible Reading Plan"

    def description(self):
        return "Discipleship Journal 1-Year Bible Reading Plan"

    def resolve(self, userObj):
        debug_utils.log("Resolving DJBR hook")

        refs = djbr_utils.get_djbr()

        if refs is not None:
            refString = "\n".join(refs)

            if refs[0].find("Reflection") != -1:
                prompt = userObj.get_reply_string(REFLECTION)
                refs = [user_actions.UserDoneAction().name()]
            else:
                prompt = PROMPT.format(refString)
                refs.append(user_actions.UserDoneAction().name())

            options = [
                telegram_utils.make_reply_button(text=ref) for ref in refs
            ]

            telegram_utils.send_reply(
                user=userObj.get_uid(),
                text=prompt,
                reply=telegram_utils.make_reply_keyboard(
                    buttons=options, width=1))

            userObj.set_state(self.identifier())


def get():
    return [DJBRDailyHook()]
