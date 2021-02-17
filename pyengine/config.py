DATABASE_DRIVER="mysql"
DATABASE_DSN="root:password@(host:port)/dbname"

PIPELINES=[
    {"name": "jiandan_ooxx", "path": "pipelines.jiandan_ooxx"},
    {"name": "jidan_ooxx_downloader", "path": "pipelines.jiandan_ooxx_downloader"}
]

try:
    from local_config import *
except:
    pass
