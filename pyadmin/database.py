# https://blog.csdn.net/woshiaotian/article/details/51442876
import pymysql
from config import *


class Singleton(type):
    _instances = {}

    def __call__(cls, *args, **kwargs):
        if cls not in cls._instances:
            cls._instances[cls] = super(Singleton, cls).__call__(*args, **kwargs)
        return cls._instances[cls]


class Database(metaclass=Singleton):
    def __init__(self, *args, **kwargs):
        self.db = pymysql.connect(
            host=kwargs.get('host', DATABASE_HOST),
            user=kwargs.get('user', DATABASE_USER),
            password=kwargs.get('password', DATABASE_PASSWORD),
            database=kwargs.get('name', DATABASE_NAME),
            charset='utf8mb4')

    def cursor(self):
        return self.db.cursor(pymysql.cursors.DictCursor)


class Redis(metaclass=Singleton):
    def __init__(self, *args, **kwargs):
        print(kwargs)