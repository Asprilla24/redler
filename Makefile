export SERVICE=redler
export REVISION_ID?=unknown
export BUILD_DATE?=unknown
export GIT_HASH?=unknown

build:
	docker build --build-arg SERVICE=$(SERVICE) \
		--build-arg REVISION_ID=$(REVISION_ID) \
		--build-arg BUILD_DATE=$(BUILD_DATE) \
		--build-arg GIT_HASH=$(GIT_HASH) \
		--tag="$(SERVICE):$(REVISION_ID)" --tag="$(SERVICE):latest" -f build/Dockerfile .

run:
	docker-compose -f build/docker-compose.yaml up

.PHONY: build run