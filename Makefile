.PHONY: deps vet test dev build clean buildFront

PACKAGES = $(shell glide novendor)
DOCKER_REPO_URL = jack08300/childrenlab_admin

deps:
	dep ensure

vet:
	go vet $(PACKAGES)

dev:
	DATABASE_NAME=childre_qu \
	DATABASE_IP=165.227.250.24:3306 \
	DATABASE_DEBUG=true \
	DATABASE_USER=root \
	DATABASE_PASSWORD=koe7POut \
	DEBUG=true \
	PORT=1323 \
	gin -p 8114 -a 1323 -x node_modules

build: clean buildFront
	GOOS=linux go build -o ./build/main *.go

clean:
	rm -f view/build/*
	rm -rf build/*
	find . -name '*.test' -delete

push-image: build build-image
	docker tag childrenlab_admin $(DOCKER_REPO_URL):latest
	docker push $(DOCKER_REPO_URL):latest

buildFront:
	yarn build

build-image:
	docker build --rm -t childrenlab_admin:latest .