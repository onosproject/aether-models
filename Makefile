# SPDX-FileCopyrightText: 2023-present Intel Corporation
# SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

SHELL 			  		= bash -e -o pipefail
KIND_CLUSTER_NAME 		?= kind
DOCKER_USER       		?=
DOCKER_PASSWORD   		?=
MODEL_COMPILER_VERSION  ?= v0.11.6

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

models: clean # @HELP Generate Golang code for all the models
	@for model in models/*; do \
		echo -e "Generating $$model:\n"; \
		docker run -v $$(pwd)/$$model:/config-model onosproject/model-compiler:${MODEL_COMPILER_VERSION}; \
		for yangfile in $$model/yang/*; do \
			sed '1 i\// SPDX-License-Identifier: Apache-2.0\n' $$(pwd)/$$yangfile > $$(pwd)/$$yangfile.temp ; \
			sed '1 i\// SPDX-FileCopyrightText: 2022-present Intel Corporation\n// SPDX-FileCopyrightText: 2021 Open Networking Foundation\n//' $$(pwd)/$$yangfile.temp > $$(pwd)/$$yangfile ; \
			rm $$(pwd)/$$yangfile.temp; \
		done; \
		echo -e "\n\n"; \
	done


models-openapi: # @HELP generates the openapi specs for the models
	@for model in models/*; do echo -e "Building OpenApi Specs for $$model:\n"; make -C $$model openapi; echo -e "\n\n"; done

docker-build: models models-openapi # @HELP Build Docker containers for all the models
	@for model in models/*; do echo -e "Building container for $$model:\n"; make -C $$model image; echo -e "\n\n"; done

docker-push: # @HELP Publish Docker containers for all the models
	@for model in models/*; do make -C $$model publish; done

kind-load: # @HELP Load Docker containers for all the models in a kind cluster (use: KIND_CLUSTER_NAME to customize the cluster name)
	@for model in models/*; do make -C $$model kind; done

clean:	# @HELP Removes the generated code
	@for model in models/*; do pushd $$model; rm -f Dockerfile Makefile *.tree; popd; done;

check-models-tag: # @HELP check that the model tags are valid
	@for model in models/*; do make -C $$model check-tag; done;

test: check-models-tag models models-openapi # @HELP Make sure the generated code has been committed
	# @bash test/generated.sh # TODO uncomment after AETHER-3550 is solved
	@for model in models/*; do make -C $$model test; done;

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
	cd .. && go install github.com/jstemmer/go-junit-report@latest && go install github.com/t-yuki/gocover-cobertura@latest

build-tools: # @HELP install the ONOS build tools if needed
	@if [ ! -d "../build-tools" ]; then cd .. && git clone https://github.com/onosproject/build-tools.git; fi

jenkins-publish: build-tools jenkins-tools docker-build docker-login docker-push# @HELP Target used by Jenkins to publish docker images
