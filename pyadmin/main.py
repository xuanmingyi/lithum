import tornado.ioloop
import tornado.web
import os.path
import json

from database import Database
from base import load_models

import json, datetime
from tornado import escape
DT_HANDLER = lambda obj: obj.isoformat() if isinstance(obj, datetime.datetime) or isinstance(obj, datetime.date) else None
def json_encode(value):
    return json.dumps(value, default=DT_HANDLER).replace("</", "<\/")

escape.json_encode = json_encode

template_path = os.path.join(os.path.dirname(__file__), 'templates')
static_path = os.path.join(os.path.dirname(__file__), 'statics')


class IndexHandler(tornado.web.RequestHandler):
    def get(self):
        models = load_models()
        if len(models) < 1:
            self.finish("404 Not Found")
        url = '/manage/{}/index'.format(models[0].name)
        print(url)
        self.redirect(url)


class ManageHandler(tornado.web.RequestHandler):
    def initialize(self, db):
        self.db = db

    def get(self, name):
        models = load_models()
        current_model = None
        for _model in models:
            if _model.name == name:
                current_model = _model

        kwargs = {
            'title': '',
            'models': models,
            'current_model': current_model
        }

        self.render('index.html', **kwargs)


class TableDataHandler(tornado.web.RequestHandler):
    def initialize(self, db):
        self.db = db

    def get(self, name):
        page = self.get_argument('page', 1)
        limit = self.get_argument('limit', 10)

        cursor = self.db.cursor()
        cursor.execute('SELECT COUNT(id) FROM {}'.format(name))
        count = cursor.fetchone()

        rows = cursor.execute('SELECT * FROM {} ORDER BY id DESC LIMIT {} OFFSET {}'.format(name, limit, (int(page) - 1) * (int(limit))))

        data = []
        for row in cursor.fetchall():
            data.append(row)

        self.write(tornado.escape.json_encode({'code': 0, 'msg': '', 'count': count['COUNT(id)'], 'data': data}))


class FormHandler(tornado.web.RequestHandler):

    def initialize(self, db):
        self.db = db

    def get(self, name):
        kwargs = {}
        id = self.get_argument('id')
        action_name = self.get_argument('action')

        current_object = self.db.get_by_id(name, id)

        # 获取model
        models = load_models()
        current_model = None
        for _model in models:
            if _model.name == name:
                current_model = _model

        current_action = None

        for action in current_model.table.actions:
            if action.name == action_name:
                current_action = action
        else:
            for action in current_model.table.row_actions:
                if action.name == action_name:
                    current_action = action

        kwargs = {
            'current_model': current_model,
            'current_action': current_action,
            'current_object': current_object
        }
        self.render('form.html', **kwargs)


def make_app():
    initialize_options = {
        'db': Database(),
    }
    options = {
        'template_path': template_path,
        'static_path': static_path,
        'debug': True
    }
    return tornado.web.Application([
        (r'/', IndexHandler),
        (r'/manage/(?P<name>.+)/index', ManageHandler, dict(initialize_options)),
        (r'/manage/(?P<name>.+)/data', TableDataHandler, dict(initialize_options)),
        (r'/manage/(?P<name>.+)/form', FormHandler, dict(initialize_options))
    ], **options)


if __name__ == '__main__':
    app = make_app()
    app.listen(8881)
    tornado.ioloop.IOLoop.current().start()
