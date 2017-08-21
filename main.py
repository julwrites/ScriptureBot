# Python std modules
import webapp2

from common.constants import APP_MAIN_URL

class MainPage(webapp2.RequestHandler):
    def get(self):
        self.response.headers['Content-Type'] = 'text/plain'
        self.response.write("Hi,I'm julwrites")

app = webapp2.WSGIApplication([
    # (url being accessed, class to call)
    (APP_MAIN_URL, MainPage),
], debug=True)
