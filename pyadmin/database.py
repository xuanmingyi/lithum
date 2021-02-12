# https://blog.csdn.net/woshiaotian/article/details/51442876
import os

class Singleton(type):
    _instances = {}

    def __call__(cls, *args, **kwargs):
        if cls not in cls._instances:
            cls._instances[cls] = super(Singleton, cls).__call__(*args, **kwargs)
        return cls._instances[cls]


class Database(metaclass=Singleton):
    def __init__(self, *args, **kwargs):
        print(kwargs.get('dsn'))


class Redis(metaclass=Singleton):
    def __init__(self, *args, **kwargs):
        print(kwargs)