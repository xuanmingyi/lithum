DATABASE_HOST='127.0.0.1'
DATABASE_PORT=3306
DATABASE_USER='root'
DATABASE_PASSWORD='root'
DATABASE_NAME='sys'

try:
    from local_config import *
except:
    pass