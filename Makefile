PACKAGES = $(shell glide novendor)

deps:
	dep ensure

vet:
	go vet $(PACKAGES)

dev:
	DATABASE_NAME=childre_qu \
	DATABASE_IP=104.236.47.92:3306 \
	DATABASE_DEBUG=true \
	DATABASE_USER=root \
	DATABASE_PASSWORD=koe7POut \
	DEBUG=true \
	PORT=1323 \
	gin -p 8114 -a 1323 -x node_modules
