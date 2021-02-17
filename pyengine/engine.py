import threading
import time

from config import *

pipelines = []


class Pipeline(threading.Thread):
    def __init__(self, *args, **kwargs):
        threading.Thread.__init__(self, name=kwargs.get('name'))
        self.name = kwargs.get('name')
        self.path = kwargs.get('path')

        print(self.name, self.path)

    def run(self):
        module = __import__(self.path)
        print(dir(module.jiandan ))
        self.obj = self._class()
        while True:
            print(self.obj.get_config())
            print(self.obj.run())
            print(self.name)


def load_pipelines():
    for pipeline in PIPELINES:
        pipelines.append(Pipeline(**pipeline))


def start_pipelines():
    for pipeline in pipelines:
        pipeline.start()


def join_pipelines():
    for pipeline in pipelines:
        pipeline.join()


def main():
    load_pipelines()
    start_pipelines()
    join_pipelines()


if __name__ == '__main__':
    main()
