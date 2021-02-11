class Table(object):
    class Column(object):

        def __int__(self, *args, **kwargs):
            self.name = kwargs.get('name')
            self.dispaly = kwargs.get('display')
            self.width = kwargs.get('width', 60)
            self.align = kwargs.get('align', 'center')

        @classmethod
        def create(cls, *args, **kwargs):
            return cls(*args, **kwargs)

    def __init__(self, *args, **kwargs):
        self.columns = kwargs.get('columns', [])

    @classmethod
    def create(cls, *args, **kwargs):
        return cls(*args, **kwargs)


class Attribute(object):
    class Option(object):
        def __init__(self, *args, **kwargs):
            self.display = kwargs.get('display')
            self.value = kwargs.get('value')

        @classmethod
        def create(cls, *args, **kwargs):
            return cls(*args, **kwargs)

    def __init__(self, *args, **kwargs):
        self.name = kwargs.get('name')
        self.display = kwargs.get('display')
        self.tag = kwargs.get('tag')
        if self.tag == 'select':
            self.options = []
            for option in kwargs.get('options'):
                self.options.append(self.Option.create(**option))


    @classmethod
    def create(cls, *args, **kwargs):
        return cls(*args, **kwargs)


class Model(object):
    def __init__(self, *args, **kwargs):
        self.name = kwargs.get('name')
        self.title = kwargs.get('title')
        self.attributes = []
        for attribute in kwargs.get('attributes', []):
            self.attributes.append(Attribute.create(**attribute))

    @classmethod
    def create(cls, *args, **kwargs):
        return cls(*args, **kwargs)

def load_models():
    jiandan_ooxx_model = {
        'name': 'jiandan_ooxx',
        'title': '煎蛋图片',
        'attributes': [
            {'name': 'id', 'display': 'ID', 'tag': 'input'},
            {'name': 'url', 'display': '地址', 'tag': 'input'},
            {'name': 'status', 'display': '状态', 'tag': 'select',
                'options': [{'display': '新抓取', 'value': 'new'}, {'display': '已下载', 'value': 'download'}]},
            {'name': 'date', 'display': '日期', 'tag': 'input'},
        ],
        'table': {
            'columns': [
                {'name': 'id', 'display': 'ID'},
                {'name': 'url', 'display': '地址'},
                {'name': 'status', 'display': '状态'},
                {'name': 'date', 'display': '日期'}
            ],
            'actions': [],
            'row_actions': []
        }
    }


    models = [
        Model.create(**jiandan_ooxx_model)
    ]
    return models


if __name__ == '__main__':
    for model in load_models():
        print(model.name)
        print(model.attributes)