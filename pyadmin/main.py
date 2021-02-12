import tornado.ioloop
import tornado.web
import os.path

from database import Database

template_path = os.path.join(os.path.dirname(__file__), 'templates')
static_path = os.path.join(os.path.dirname(__file__), 'statics')


class IndexHandler(tornado.web.RequestHandler):
    def get(self):
        self.write('hello, world')


class ManageHandler(tornado.web.RequestHandler):
    def initialize(self, db):
        self.db = db

    def get(self):
        kwargs = {
            'title': 'ssss'
        }
        self.render('index.html', **kwargs)


def make_app():
    initialize_options = {
        'db': Database(),
    }
    options = {
        'template_path': template_path,
        'static_path': static_path
    }
    return tornado.web.Application([
        (r'/', ManageHandler, dict(initialize_options))
    ], **options)


if __name__ == '__main__':
    app = make_app()
    app.listen(8881)
    tornado.ioloop.IOLoop.current().start()
