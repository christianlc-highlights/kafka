CRD := https://github.com/banzaicloud/koperator/releases/download/v0.16.0/kafka-operator.crds.yaml

.DEFAULT_GOAL := @goal
.ONESHELL:

## recipe
@goal: cleandist dist install

dist:
	mkdir $@
	cp -rf manifest $@
	cat manifest/kafka.yaml | tee $@/manifest.yaml

install:
	#kubectl apply -f dist/manifest.yaml -n kafka
	kubectl create --validate=false -f dist/manifest/crd.yaml ||:
	helm repo add banzaicloud-stable https://kubernetes-charts.banzaicloud.com
	helm template -f dist/manifest/values.yaml \
		kafka-operator \
			--create-namespace \
			--namespace=kafka \
		banzaicloud-stable/kafka-operator \
	| tee dist/manifest/generated.yaml
	kubectl apply \
		-f dist/manifest/generated.yaml \
		-n kafka

cleandist:
	rm -rvf dist

clean:
	kubectl delete -f dist/manifest.yaml ||:

	helm delete kafka-operator ||:
	kubectl delete -f dist/manifest/crd.yaml ||:
	kubectl delete -f dist/manifest/generated.yaml

