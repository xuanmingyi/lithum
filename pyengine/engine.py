from pipelines import PIPELINES
from config import Config


class Runner(object):
    pipelines = []

    def __init__(self):
        pipeline_names = list(map(lambda x: x.strip(), Config.Get('default.pipelines').split(',')))

        for _class in PIPELINES:
            c = _class()
            if c.name in pipeline_names:
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
