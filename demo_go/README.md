# Golang and EFK

## How to:

Notice: You need to change the ```fluentd-address``` if you are not on MacOS:

	$ make start
	# check whether everyhing runs
	$ docker ps

Available endpoints:

- [http://127.0.0.1:8080/hello](http://127.0.0.1:8080/hello) - service
- [http://127.0.0.1:5601](http://127.0.0.1:5601) - Kibana
- [http://127.0.0.1:9200](http://127.0.0.1:9200) - ElasticSearch

## TODO:

- extend the golang example
- add [giantswarm/curator](https://github.com/giantswarm/curator) to the demo

## Reference:

- https://docs.fluentd.org/v0.12/articles/docker-logging-efk-compose
- https://www.fluentd.org/guides/recipes/docker-logging
- https://banzaicloud.com/blog/runtime-logging/
- https://godoc.org/golang.org/x/xerrors
- https://kuther.net/blog/prometheus-run-queries-against-elasticsearch-and-turn-it-metrics-and-alerts 
