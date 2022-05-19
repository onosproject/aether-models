# SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

SHELL 			  		= bash -e -o pipefail
KIND_CLUSTER_NAME 		?= kind
DOCKER_USER       		?=
DOCKER_PASSWORD   		?=
MODEL_COMPILER_VERSION  ?= v0.9.28

.PHONY: models

help: # @HELP Print the command options
	@echo
	@echo -e "\033[0;31m    Aether Models \033[0m"
	@echo
	@grep -E '^.*: .* *# *@HELP' $(MAKEFILE_LIST) \
    | sort \
    | awk ' \
        BEGIN {FS = ": .* *# *@HELP"}; \
        {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}; \
	'

models: clean# @HELP Generate Golang code for all the models
	@cd models && for model in *; do echo -e "Generating $$model:\n"; docker run -v $$(pwd)/$$model:/config-model onosproject/model-compiler:${MODEL_COMPILER_VERSION}; echo -e "\n\n"; done

models-openapi: # @HELP generates the openapi specs for the models
	@cd models && for model in *; do echo -e "Building OpenApi Specs for $$model:\n"; pushd $$model; make openapi; popd; echo -e "\n\n"; done

docker-build: models models-openapi # @HELP Build Docker containers for all the models
	@cd models && for model in *; do echo -e "Building container for $$model:\n"; pushd $$model; make image; popd; echo -e "\n\n"; done

docker-push: # @HELP Publish Docker containers for all the models
	@cd models && for model in *; do pushd $$model; make publish; popd; done

kind-load: # @HELP Load Docker containers for all the models in a kind cluster (use: KIND_CLUSTER_NAME to customize the cluster name)
	@cd models && for model in *; do pushd $$model; make kind; popd; done

yang-lint: # @HELP Lint the yang models
yang-lint: pyang-tool
	@cd models && for model in *; do echo -e "Linting YANG files for: $$model"; pyang --lint --ignore-error=XPATH_FUNCTION $$model/yang/*.yang; done

pyang-tool: # @HELP Install the Pyang tool if needed
	pyang --version || python -m pip install pyang==2.5.2

clean:	# @HELP Removes the generated code
	@cd models && for model in *; do pushd $$model; rm -f Dockerfile Makefile *.tree; popd; done;

test: yang-lint models models-openapi# @HELP Make sure the generated code has been committed
	@bash test/generated.sh
	cd models/aether-2.0.x && make test
	cd models/aether-2.1.x && make test
	cd models/aether-4.x && make test

docker-login:
ifdef DOCKER_USER
ifdef DOCKER_PASSWORD
	echo ${DOCKER_PASSWORD} | docker login -u ${DOCKER_USER} --password-stdin
else
	@echo "DOCKER_USER is specified but DOCKER_PASSWORD is missing"
	@exit 1
endif
endif

jenkins-tools: # @HELP installs tooling needed for Jenkins
	cd .. && go get -u github.com/jstemmer/go-junit-report && go get github.com/t-yuki/gocover-cobertura

build-tools: # @HELP install the ONOS build tools if needed
	@if [ ! -d "../build-tools" ]; then cd .. && git clone https://github.com/onosproject/build-tools.git; fi

jenkins-publish: build-tools jenkins-tools docker-build docker-login docker-push# @HELP Target used by Jenkins to publish docker images
