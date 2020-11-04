SHELL := /bin/bash

# Variables definitions
# -----------------------------------------------------------------------------
IMAGE_NAME := arthurhenrique/go-ocr

ifeq ($(ENV_VAR_SAMPLE),)
ENV_VAR_SAMPLE := "Environment variable not exported!"
endif

# Target section and Global definitions
# -----------------------------------------------------------------------------
CGO_ENABLED=1


.PHONY: clean docker

docker/build:
	docker build -t $(IMAGE_NAME) -f Dockerfile .

docker/shell:
	docker run  -it -v "$(shell pwd)/:/go/src/github.com/arthurhenrique/go-ocr" -e CGO_ENABLED $(IMAGE_NAME) sh


install:
	go mod vendor

clean:
	@find . -name '*.txt' -exec rm -rf {} \;
	@find . -name 'Thumbs.db' -exec rm -rf {} \;
	@find . -name '*~' -exec rm -rf {} \;
	rm -rf .cache
	rm -rf build