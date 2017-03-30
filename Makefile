all: build

clean: cleanbin cleangenerated cleandocs

cleanbin:
	rm -f main

cleangenerated:
	bash -c "find ./apis/ -name zz_generated.api.*.go | xargs rm -f"

generate: cleangenerated
	go generate main.go
	go run vendor/k8s.io/kubernetes/cmd/libs/go2idl/conversion-gen/main.go -i github.com/pwittrock/apiserver-helloworld/apis/...  --extra-peer-dirs="k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/apimachinery/pkg/conversion,k8s.io/apimachinery/pkg/runtime" -o ~/apiserver-helloworld/src/  -O zz_generated.conversion --go-header-file boilerplate.go.txt
	go run vendor/k8s.io/kubernetes/cmd/libs/go2idl/deepcopy-gen/main.go -i github.com/pwittrock/apiserver-helloworld/apis/... -o ~/apiserver-helloworld/src/ -O zz_generated.deepcopy --go-header-file boilerplate.go.txt
	go run vendor/k8s.io/kubernetes/cmd/libs/go2idl/openapi-gen/main.go  -i "github.com/pwittrock/apiserver-helloworld/apis/...,k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/apimachinery/pkg/api/resource/,k8s.io/apimachinery/pkg/version/,k8s.io/apimachinery/pkg/runtime/,k8s.io/apimachinery/pkg//util/intstr/" --output-package "github.com/pwittrock/apiserver-helloworld/pkg/openapi" --go-header-file boilerplate.go.txt

build: cleanbin generate
	go build main.go

cleandocs:
	rm -rf $(shell pwd)/docs/build
	rm -rf $(shell pwd)/docs/includes
	rm -rf $(shell pwd)/docs/manifest.json
	rm -rf $(shell pwd)/docs/includes/_generated_*

docs: cleandocs
	go run vendor/github.com/kubernetes-incubator/reference-docs/main.go --doc-type open-api --allow-errors --use-tags --config-dir docs --gen-open-api-dir vendor/github.com/kubernetes-incubator/reference-docs/gen_open_api
	docker run -v $(shell pwd)/docs/includes:/source -v $(shell pwd)/docs/build:/build -v $(shell pwd)/docs/:/manifest pwittrock/brodocs

run: build
	./main -v 10 --authentication-kubeconfig ~/.kube/auth_config --authorization-kubeconfig ~/.kube/auth_config --client-ca-file /var/run/kubernetes/apiserver.crt  --requestheader-client-ca-file /var/run/kubernetes/apiserver.crt --requestheader-username-headers=X-Remote-User --requestheader-group-headers=X-Remote-Group --requestheader-extra-headers-prefix=X-Remote-Extra- --etcd-servers=http://localhost:2379 --secure-port=9443 --tls-ca-file  /var/run/kubernetes/apiserver.crt

# Alias for setting up docker command
# export GOPATH=/Users/pwittroc/test
# export REPO_NAME=github.com/pwittrock/test
# alias kc="docker run -i -t -v  $GOPATH:/out pwittrock/kubec"
# kc init --repo-name=github.com/pwittrock/test


# docker run -i -t -v  /Users/pwittroc/test:/out --entrypoint bash pwittrock/kubec
# docker run -i -t -v  /Users/pwittroc/test/src/:/out pwittrock/kubec
# docker run -i -t -v  /Users/pwittroc/test/:/out pwittrock/kubec init --repo-name=$REPO_NAME
# docker run -i -t -v  /Users/pwittroc/test/:/out pwittrock/kubec add-types --repo-name=$REPO_NAME --repo-package github.com/pwittrock/apiserver-helloworld --types hyrule/v1/HyruleCastle,mushroomkingdom/v2/PeachesCastle
# docker run -i -t -v  /Users/pwittroc/test/:/out pwittrock/kubec generate --repo-name=$REPO_NAME
