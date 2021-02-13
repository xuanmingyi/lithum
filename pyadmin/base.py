class Action(object):
    class Field(object):
        def __init__(self, *args, **kwargs):
            self.name = kwargs.get('name')
            self.disabled = kwargs.get('disabled', False)

        @classmethod
        def create(cls, *args, **kwargs):
            return cls(*args, **kwargs)


    def __init__(self, *args, **kwargs):
        self.name = kwargs.get('name')
        self.display = kwargs.get('display')
        self._class = kwargs.get('class')

    @classmethod
    def create(cls, *args, **kwargs):
        return cls(*args, **kwargs)


class Table(object):
    class Column(object):

        def __init__(self, *args, **kwargs):
            self.name = kwargs.get('name')
            self.display = kwargs.get('display')
            self.width = kwargs.get('width', 0)
            self.align = kwargs.get('align', '')
            self.templet = kwargs.get('templet', '')

        @classmethod
        def create(cls, *args, **kwargs):
            return cls(*args, **kwargs)

    def __init__(self, *args, **kwargs):
        self.columns = []
        for column in kwargs.get('columns', []):
            self.columns.append(self.Column.create(**column))
        self.actions = []
        for action in kwargs.get('actions'):
            self.actions.append(Action.create(**action))
        self.row_actions = []
        for action in kwargs.get('row_actions'):
            self.row_actions.append(Action.create(**action))


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
        self.table = Table.create(**kwargs.get('table', {}))

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
                {'name': 'id', 'display': 'ID', 'width': 80},
                {'name': 'url', 'display': '地址'},
                {'name': 'status', 'display': '状态', 'width': 80, 'align': 'center',
                 'templet': """ function(res) { if (res.status == 'download') { return '<i class="layui-icon layui-icon-ok-circle"></i>' }} """},
                {'name': 'date', 'display': '日期', 'width': 180}
            ],
            'actions': [
                {'name': 'create', 'display': '创建', 'type': 'dialog',
                 'dialog': {'title': '创建实例',
                            'fields': [{'name': 'id'}, {'name': 'url'}, {'name': 'status'}, {'name': 'date'}]
                            }
                },
                {'name': 'batch_delete', 'display': '批量删除', 'type': 'dialog', 'class': 'layui-btn-danger',
                 'dialog': {'title': '创建实例',
                            'fields': [{'name': 'id'}, {'name': 'url'}, {'name': 'status'}, {'name': 'date'}]
                            }
                }
            ],
            'row_actions': [
                {'name': 'view', 'display': '查看', 'type': 'dialog', 'class': 'layui-btn-primary',
                 'dialog': {'title': '创建实例',
                            'fields': [{'name': 'id', 'disabled': True},
                                       {'name': 'url', 'disabled': True},
                                       {'name': 'status', 'disabled': True},
                                       {'name': 'date', 'disabled': True}]
                            }
                 },
                {'name': 'edit', 'display': '编辑', 'type': 'dialog',
                 'dialog': {'title': '编辑实例',
                            'fields': [{'name': 'id', 'disabled': True},
                                       {'name': 'url'},
                                       {'name': 'status'},
                                       {'name': 'date'}]
                            }
                 },
                {'name': 'delete', 'display': '删除', 'type': 'dialog', 'class': 'layui-btn-danger',
                 'dialog': {'title': '删除实例',
                            'fields': [{'name': 'id', 'disabled': True},
                                       {'name': 'url', 'disabled': True},
                                       {'name': 'status', 'disabled': True},
                                       {'name': 'date', 'disabled': True}]
                            }
                 }
            ]
        }
    }


    models = [
        Model.create(**jiandan_ooxx_model)
    ]
    return models