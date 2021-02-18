from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker
from sqlalchemy import create_engine
from config import DATABASE_DSN

engine = create_engine(DATABASE_DSN)
Base = declarative_base()
Session = sessionmaker(engine)