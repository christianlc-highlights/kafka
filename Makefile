export NAME := $(shell basename "$$PWD" )
export ORG := christianelsee
export SHA := $(shell git rev-parse --short HEAD)
export TS  := $(shell date +%s)


.DEFAULT_GOAL := @goal
.ONESHELL:
.POSIX:

## recipe
@goal: distclean dist build

dist:
	mkdir $@
	go mod init github.com/christianlc-highlights/kafka ||:
	go mod tidy

	cat manifest/cli.yaml \
		| envsubst \
		| tee $@/manifest.yaml
	helm repo add bitnami https://charts.bitnami.com/bitnami
	helm repo update
	helm template -f helm/values.yaml \
		$(NAME) \
			--create-namespace \
			--namespace=$(NAME) \
		bitnami/$(NAME) \
	| tee -a $@/manifest.yaml

build: dist
	docker build \
		-t local/$(NAME) \
		-t docker.io/$(ORG)/$(NAME) \
		-t docker.io/$(ORG)/$(NAME):$(SHA) \
		.

namespace:
	kubectl create namespace $(NAME) \
		--dry-run=client \
		-oyaml \
	| kubectl apply -f-
	kubectl config set-context \
		--current \
		--namespace $(NAME)

install: build namespace
	<secrets/docker.io.token.gpg gpg -d \
		| xargs -- \
			docker login \
				-u $(ORG) \
				-p

	# if push fails, it may be due to docker image
	# with current sha, not having been built yet
	docker push docker.io/$(ORG)/$(NAME):$(SHA)
	docker push docker.io/$(ORG)/$(NAME):latest
	kubectl apply \
		-f dist/manifest.yaml \
		-n $(NAME)

distclean:
	rm -rvf dist

clean:
	kubectl delete -f dist/manifest.yaml

## sanity targets
lint:
	goimports -l .
	golint ./...
	go vet ./... ||:

local-build:
	go build -o dist/build main.go
