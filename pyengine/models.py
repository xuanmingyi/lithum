from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker
from sqlalchemy import create_engine
from sqlalchemy import Column, String, Integer, DateTime

from config import DATABASE_DSN
from datetime import datetime


engine = create_engine(DATABASE_DSN, pool_pre_ping=True)
Base = declarative_base()
SessionFactory = sessionmaker(engine)


class JiandanImage(Base):
    __tablename__ = "jiandan"

    id = Column(Integer, primary_key=True)
    url = Column(String(256))
    status = Column(String(32))
    date = Column(DateTime, default=datetime.now())


class LVV2Image(Base):
    __tablename__ = "lvv2_image"

    id = Column(Integer, primary_key=True)
    thread_id = Column(Integer)
    url = Column(String(256))
    status = Column(String(32))
    date = Column(DateTime, default=datetime.now())


class LVV2Thread(Base):
    __tablename__ = "lvv2_thread"

    id = Column(Integer, primary_key=True)
    title = Column(String(255))
    url = Column(String(255))
    status = Column(String(32))
    date = Column(DateTime, default=datetime.now())
    tag = Column(String(32))
