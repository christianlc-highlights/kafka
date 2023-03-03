
export NAME := $(shell basename "$$PWD" )
export ORG := christianelsee
sha := $(shell git rev-parse --short HEAD)

.DEFAULT_GOAL := @goal
.ONESHELL:

## recipe
@goal: distclean dist build

dist:
	mkdir $@
	cp -rf manifest $@
	go mod init github.com/christianlc-highlights/kafka ||:
	go mod tidy

lint:
	goimports -l .
	golint ./...
	go vet ./... ||:

build: dist
	docker build \
		-t local/$(NAME) \
		-t docker.io/$(ORG)/$(NAME) \
		-t docker.io/$(ORG)/$(NAME):$(sha) \
		.

namespace:
	kubectl create namespace $(NAME) \
		--dry-run=client \
		-oyaml \
	| kubectl apply -f-
	kubectl config set-context \
		--current \
		--namespace $(NAME)

install: namespace
	<secrets/docker.io.token.gpg gpg -d \
		| xargs -- \
			docker login \
				-u $(ORG) \
				-p
	docker push docker.io/$(ORG)/$(NAME):$(sha)

	helm repo add bitnami https://charts.bitnami.com/bitnami
	helm repo update
	helm template -f dist/manifest/values.yaml \
		kafka \
			--create-namespace \
			--namespace=$(NAME) \
		bitnami/$(NAME) \
	| tee dist/manifest/generated.yaml
	kubectl apply \
		-f dist/manifest/generated.yaml \
		-n $(NAME)

distclean:
	rm -rvf dist

clean:
	kubectl delete -f dist/manifest/generated.yaml
