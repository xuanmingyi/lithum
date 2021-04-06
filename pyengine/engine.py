import logging
import sys

from pipelines import PIPELINES

LOG_FORMAT = "%(asctime)s - %(levelname)s - %(message)s"
logging.basicConfig(stream=sys.stdout, level=logging.DEBUG, format=LOG_FORMAT)


class Runner(object):
    pipelines = []

    def __init__(self):
        for _class in PIPELINES:
            self.pipelines.append(_class())

    def start(self):
        for pipeline in self.pipelines:
            pipeline.start()

    def join(self):
        for pipeline in self.pipelines:
            pipeline.join()


def main():
    runner = Runner()
    runner.start()
    runner.join()


if __name__ == '__main__':
    main()
