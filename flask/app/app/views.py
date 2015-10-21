from app import app
import redis

r = redis.StrictRedis(host="redis", port="6379", db=0)

@app.route('/')
@app.route('/index')
def index():
    if r.get("count") == None:
	r.set("count", 0)
    c = int(r.get("count"))
    c += 1
    r.set("count", c)
    return "This page has been visited " + r.get("count") + " times!"
