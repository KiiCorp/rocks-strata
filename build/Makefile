.DEFAULT_GOAL := help

image_name      := rocks-strata
dockerfile_rpm  := Dockerfile
rpm_target_path := /rpm

ifdef CIRCLE_BUILD_NUM
else
	CIRCLE_BUILD_NUM=local
endif

.PHONY: build_rpm_image
build_rpm_image: ## Build a docker image in which there is a rocks-strata rpm package
	docker build -t $(image_name) .. -f $(dockerfile_rpm) --build-arg CIRCLE_BUILD_NUM=$(CIRCLE_BUILD_NUM) --build-arg RPM_TARGET_PATH=$(rpm_target_path)

.PHONY: cp
cp: ## Copy rpm package
	$(eval CONTAINER_ID := $(shell docker run -d $(image_name)))
	docker cp $(CONTAINER_ID):$(rpm_target_path) .
	docker rm -f $(CONTAINER_ID)

.PHONY: build_local
build_local: ## Build rocks-strata rpm package
	make build_rpm_image cp

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
	| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
