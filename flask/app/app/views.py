from app import app
from urlparse import urlparse
import os
import redis

redis_port = os.environ['REDIS_PORT']

r = redis.StrictRedis(host=urlparse(redis_port).hostname,
                      port=urlparse(redis_port).port,
                      db=0)


@app.route('/')
@app.route('/index')
def index():
    if r.get("count") == None:
        r.set("count", 0)
    c = int(r.get("count"))
    c += 1
    r.set("count", c)
    return "<div align=center>" +\
        "<img src='https://pbs.twimg.com/profile_images/" +\
        "458352291767013376/K9nN_rhH_400x400.png'>" +\
        "<h1>This page has been visited " + r.get("count") + " times!</h1>" +\
        "<br>" +\
        "</div>"
