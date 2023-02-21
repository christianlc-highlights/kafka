
export NAME := $(shell basename "$$PWD" )

.DEFAULT_GOAL := @goal
.ONESHELL:

## recipe
@goal: distclean dist install

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
	go build -o dist/build main.go

namespace:
	kubectl create namespace $(NAME) \
		--dry-run=client \
		-oyaml \
	| kubectl apply -f-
	kubectl config set-context \
		--current \
		--namespace $(NAME)

install: namespace build
	helm repo add bitnami https://charts.bitnami.com/bitnami
	helm template -f dist/manifest/values.yaml \
		kafka\
			--create-namespace \
			--namespace=kafka \
		bitnami/kafka \
	| tee dist/manifest/generated.yaml
	kubectl apply \
		-f dist/manifest/generated.yaml \
		-n kafka

distclean:
	rm -rvf dist

clean:
	kubectl delete -f dist/manifest/generated.yaml
