from configparser import ConfigParser
import os


class Config:
    __data__ = None

    @classmethod
    def Load(cls):
        current_path = os.path.abspath(os.path.dirname(__file__))

        cls.__data__ = {}

        conf = ConfigParser()
        conf.read(os.path.join(current_path, 'alembic.ini'), encoding='utf-8')
        cls.__data__['default.dsn'] = conf.get('alembic', 'sqlalchemy.url')

        for filename in ['config.ini', 'local_config.ini']:
            conf = ConfigParser()
            conf.read(os.path.join(current_path, filename), encoding='utf-8')
            for section in conf.sections():
                for item in conf.items(section):
                    cls.__data__['{}.{}'.format(section, item[0])] = item[1].strip()

    @classmethod
    def Get(cls, name):
        if not cls.__data__:
            cls.Load()
        return cls.__data__.get(name)
