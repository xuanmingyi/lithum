from pipelines.jiandan_ooxx import JiandanOOXX
from pipelines.jiandan_ooxx_downloader import JiandanOOXXDownloader
from pipelines.lvv2_thread import LVV2ThreadPipeline
from pipelines.lvv2_image import LVV2ImagePipeline
from pipelines.lvv2_image_downloader import LVV2ImageDownloaderPipeline

#PIPELINES = [JiandanOOXX, JiandanOOXXDownloader, LVV2Thread]
#PIPELINES = [LVV2ThreadPipeline, LVV2ImagePipeline, LVV2ImageDownloaderPipeline]

PIPELINES = [LVV2ImageDownloaderPipeline,]