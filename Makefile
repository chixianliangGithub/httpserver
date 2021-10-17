IMAGE = http_server
VERSION = v0.0.1
CONTAINER_NAME = http-server

build:
	docker build -t $(IMAGE):$(VERSION) .

run:
	docker run --name=$(CONTAINER_NAME) -d -p 8080:8080 $(IMAGE):$(VERSION)

stop:
	docker stop $(CONTAINER_NAME)

remove:
	docker stop $(CONTAINER_NAME) && docker rm $(CONTAINER_NAME)
