SHELL 			  = bash -e -o pipefail
KIND_CLUSTER_NAME ?= kind
DOCKER_USER       ?=
DOCKER_PASSWORD   ?=

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
	@cd models && for model in *; do echo -e "Generating $$model:\n"; docker run -v $$(pwd)/$$model:/config-model onosproject/model-compiler:latest; echo -e "\n\n"; done

docker-build: models # @HELP Build Docker containers for all the models
	@cd models && for model in *; do echo -e "Buildind container for $$model:\n"; pushd $$model; make image; popd; echo -e "\n\n"; done

docker-push: # @HELP Publish Docker containers for all the models
	@cd models && for model in *; do pushd $$model; make publish; popd; done

kind-load: # @HELP Load Docker containers for all the models in a kind cluster (use: KIND_CLUSTER_NAME to customize the cluster name)
	@cd models && for model in *; do pushd $$model; make kind; popd; done

yang-lint: # @HELP Lint the yang models
yang-lint: pyang-tool
	@cd models && for model in *; do echo -e "Linting YANG files for: $$model"; pyang --lint $$model/yang/*.yang; done

pyang-tool: # @HELP Install the Pyang tool if needed
	pyang --version || python -m pip install pyang==2.5.2

clean:	# @HELP Removes the generated code
	@cd models && for model in *; do pushd $$model; rm -f Dockerfile Makefile *.tree; popd; done;

test: yang-lint models # @HELP Make sure the generated code has been committed
	@bash test/generated.sh

docker-login:
ifdef DOCKER_USER
ifdef DOCKER_PASSWORD
	echo ${DOCKER_PASSWORD} | docker login -u ${DOCKER_USER} --password-stdin
else
	@echo "DOCKER_USER is specified but DOCKER_PASSWORD is missing"
	@exit 1
endif
endif

jenkins-publish: docker-build docker-login docker-push# @HELP Target used by Jenkins to publish docker images