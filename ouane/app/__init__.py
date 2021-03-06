from flask import Flask
from flask import g
from flask.ext.sqlalchemy import SQLAlchemy
from config import IPSERVER, PORTSERVER # maxime t'es trop con
import datetime
import redis

red = redis.StrictRedis()

app = Flask(__name__)
app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///ouane.db'
app.config.from_pyfile('config.py')
app.config['DEBUG'] = True

db = SQLAlchemy(app)

from dbUtils import *

db.drop_all()
db.create_all()

def event_stream():
    pubsub = red.pubsub()
    pubsub.subscribe('ouane')
    # TODO: handle client disconnection.
    for message in pubsub.listen():
        # print message
        yield 'data: %s\n\n' % message['data']

from api import Api
a = Api(IPSERVER, PORTSERVER)
#a = Api(sys.argv[1], int(sys.argv[2]))
a.start()
#a.createColumns(1,"1",1)
from app import views
