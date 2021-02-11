import tornado.ioloop
import tornado.web

import os.path


template_path = os.path.join(os.path.dirname(__file__), 'templates')
static_path = os.path.join(os.path.dirname(__file__), 'statics')


class IndexHandler(tornado.web.RequestHandler):
    def get(self):
        self.write('hello, world')


class ManageHandler(tornado.web.RequestHandler):
    def get(self):
        self.render('index.html')


def make_app():
    options = {
        'template_path': template_path,
        'static_path': static_path
    }
    return tornado.web.Application([
        (r'/', ManageHandler)
    ], **options)


if __name__ == '__main__':
    app = make_app()
    app.listen(8888)
    tornado.ioloop.IOLoop.current().start()
