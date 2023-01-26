CRD := https://github.com/banzaicloud/koperator/releases/download/v0.16.0/kafka-operator.crds.yaml

.DEFAULT_GOAL := @goal
.ONESHELL:

## recipe
@goal: cleandist dist install

dist:
	mkdir $@
	cp -rf manifest $@
	go mod init github.com/christianlc-highlights/kafka ||:
	go mod tidy

build: dist
	go build -o dist/build main.go

install:
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

cleandist:
	rm -rvf dist

clean:
	kubectl delete -f dist/manifest/generated.yaml

