CRD := https://github.com/banzaicloud/koperator/releases/download/v0.16.0/kafka-operator.crds.yaml

.DEFAULT_GOAL := @goal
.ONESHELL:

## recipe
@goal: cleandist dist install

dist:
	mkdir $@
	cp -rf manifest $@

install:
	kubectl create --validate=false -f dist/manifest/crd.yaml ||:

	helm repo add banzaicloud-stable https://kubernetes-charts.banzaicloud.com
	helm install -f dist/manifest/values.yaml \
		kafka-operator \
			--create-namespace \
			--namespace=kafka \
	banzaicloud-stable/kafka-operator

cleandist:
	rm -rvf dist

clean:
	helm delete kafka-operator
	kubectl delete -f dist/manifest/crd.yaml ||:
