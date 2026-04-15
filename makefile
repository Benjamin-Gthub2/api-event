UNAME_S := $(shell uname -s)

PROJECT_PATH = ./
REGISTRY_URL = "localhost:32000"
DB_HOST = "192.168.71.200"
NAMESPACE = "smartone-local"

#ifeq ($(UNAME_S),Darwin)
#    REGISTRY_URL := "192.168.64.2:32000"
#endif

ifdef REGISTRY_URL_DEV
    REGISTRY_URL := $(REGISTRY_URL_DEV)
endif

ifdef REGISTRY_URL_PROD
    REGISTRY_URL := $(REGISTRY_URL_PROD)
endif

deploy-micro:
	cd "$(DIR_MICRO)" && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app . && \
	docker buildx build --platform linux/amd64 -t "$(REGISTRY_URL)/$(IMAGE):$(VERSION)" -f "./Dockerfile" . && \
    docker push "$(REGISTRY_URL)/$(IMAGE):$(VERSION)" && \
	rm -rf app

deploy-logistics-main:
	@export DIR_MICRO="." && export IMAGE="logistics.smartone" && $(MAKE) deploy-micro

deploy-micro-arm:
	cd "$(DIR_MICRO)" && CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o app . && \
	docker build -t "$(REGISTRY_URL)/$(IMAGE):$(VERSION)" -f "./Dockerfile" . && \
    docker push "$(REGISTRY_URL)/$(IMAGE):$(VERSION)" && \
	rm -rf app

deploy-logistics-main-arm:
	@export DIR_MICRO="." && export IMAGE="arm.logistics.smartone" && $(MAKE) deploy-micro-arm

deploy-logistics-main-local:
	@export DIR_MICRO="." && export IMAGE="logistics.smartone" && export VERSION="v1.0.0" && $(MAKE) deploy-micro && \
	export REGISTRY_URL="$(REGISTRY_URL)" && export NAMESPACE="$(NAMESPACE)" && export LABEL="logistics-smartone" && \
	envsubst < cicd/k8s/local/deployment-api.yaml | kubectl --kubeconfig="$(HOME)/.kube/config-local" apply -f - && \
	envsubst < cicd/k8s/local/service-api.yaml | kubectl --kubeconfig="$(HOME)/.kube/config-local" apply -f - && \
	kubectl --kubeconfig="$(HOME)/.kube/config-local" delete pods -l app="$(LABEL)" -n "$(NAMESPACE)"
