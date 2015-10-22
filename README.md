* Simple Flask + Redis app that shows and increments the the number of page 
visits each time the web page is accessed.

* To run:
~~~
$ docker run -d --name=redis -v $(pwd)/redis:/redis -p 6379 centos_redis redis-server
$ docker run -d -p 5000:5000 --link redis:redis --name flask centos_flask
~~~
* Access [http://localhost:5000](http://localhost:5000) and refresh.

* Nulecule specification is mentioned in the `nulecule` directory. But it's WIP.
