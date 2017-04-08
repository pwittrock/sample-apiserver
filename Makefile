# Copyright 201 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# NOTE: Project must live under GOPATH/src/github.com/pwittrock/apiserver-helloworld
# for the code to compile
REPO=github.com/pwittrock/apiserver-helloworld
SRC=$(REPO)/pkg/apis/...
GENERIC_API="k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/apimachinery/pkg/api/resource/,k8s.io/apimachinery/pkg/version/,k8s.io/apimachinery/pkg/runtime/,k8s.io/apimachinery/pkg//util/intstr/"
OUT=~/apiserver-helloworld/src/
CLIENT_PKG=$(REPO)/pkg/client
CLIENT_PATH=$(CLIENT_PKG)/clientset_generated
CLIENT=clientset
INTERNAL_CLIENT=internalclientset
LISTERS_PKG="$(CLIENT_PKG)/listers_generated"
INFORMERS_PKG="$(CLIENT_PKG)/informers_generated"

all: build

clean: cleanbin cleangenerated cleandocs

cleanbin:
	rm -f main

cleangenerated:
	bash -c "find ./pkg/apis/ -name zz_generated.api.*.go | xargs rm -f"
	bash -c "find ./pkg/apis/ -name zz_generated.deepcopy.go | xargs rm -f"
	bash -c "find ./pkg/apis/ -name zz_generated.conversion.go | xargs rm -f"
	rm -rf pkg/client/clientset_generated/

generate: cleangenerated
#generate:
	go run vendor/k8s.io/apiserver-builder/cmd/genwiring/main.go -i $(REPO)/pkg/apis/...
	go run vendor/k8s.io/kubernetes/cmd/libs/go2idl/conversion-gen/main.go -i "$(SRC)"  --extra-peer-dirs="k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/apimachinery/pkg/conversion,k8s.io/apimachinery/pkg/runtime" -o $(OUT)   -O zz_generated.conversion --go-header-file boilerplate.go.txt
	go run vendor/k8s.io/kubernetes/cmd/libs/go2idl/deepcopy-gen/main.go -i "$(SRC)" -o $(OUT) -O zz_generated.deepcopy --go-header-file boilerplate.go.txt
	go run vendor/k8s.io/kubernetes/cmd/libs/go2idl/openapi-gen/main.go  -i "$(SRC),$(GENERIC_API)" --output-package "$(REPO)/pkg/openapi" --go-header-file boilerplate.go.txt
	go run vendor/k8s.io/kubernetes/cmd/libs/go2idl/client-gen/main.go --clientset-path=$(CLIENT_PATH)  --clientset-name=$(CLIENT) -o $(OUT) --go-header-file boilerplate.go.txt
	go run vendor/k8s.io/kubernetes/cmd/libs/go2idl/client-gen/main.go --input-base "$(REPO)/pkg/apis/" --input "mushroomkingdom/v2,hyrulekingdom/v3" --clientset-path=$(CLIENT_PATH)  --clientset-name=$(CLIENT) -o $(OUT) --go-header-file boilerplate.go.txt
	go run vendor/k8s.io/kubernetes/cmd/libs/go2idl/client-gen/main.go --input-base "$(REPO)/pkg/apis/" --input "mushroomkingdom/,hyrulekingdom/" --clientset-path=$(CLIENT_PATH) --clientset-name=$(INTERNAL_CLIENT)  -o $(OUT) --go-header-file boilerplate.go.txt
	go run vendor/k8s.io/kubernetes/cmd/libs/go2idl/lister-gen/main.go -i $(SRC) --output-package=$(LISTERS_PKG) -o $(OUT) --go-header-file boilerplate.go.txt
	go run vendor/k8s.io/kubernetes/cmd/libs/go2idl/informer-gen/main.go -i $(SRC) -o $(OUT) --go-header-file boilerplate.go.txt --internal-clientset-package="$(CLIENT_PATH)/$(INTERNAL_CLIENT)" --versioned-clientset-package="$(CLIENT_PATH)/$(CLIENT)" --listers-package=$(LISTERS_PKG) --output-package=$(INFORMERS_PKG)

build: cleanbin generate
	go build main.go

cleandocs:
	rm -rf $(shell pwd)/docs/build
	rm -rf $(shell pwd)/docs/includes
	rm -rf $(shell pwd)/docs/manifest.json
	rm -rf $(shell pwd)/docs/includes/_generated_*

docs: cleandocs build
	./main --delegated-auth=false --etcd-servers=http://localhost:2379 --secure-port=9443 --print-openapi > ./docs/openapi-spec/swagger.json
	go run vendor/github.com/kubernetes-incubator/reference-docs/main.go --doc-type open-api --allow-errors --use-tags --config-dir docs --gen-open-api-dir vendor/github.com/kubernetes-incubator/reference-docs/gen_open_api
	docker run -v $(shell pwd)/docs/includes:/source -v $(shell pwd)/docs/build:/build -v $(shell pwd)/docs/:/manifest pwittrock/brodocs

run: build
	./main -v 10 --authentication-kubeconfig ~/.kube/auth_config --authorization-kubeconfig ~/.kube/auth_config --client-ca-file /var/run/kubernetes/apiserver.crt  --requestheader-client-ca-file /var/run/kubernetes/apiserver.crt --requestheader-username-headers=X-Remote-User --requestheader-group-headers=X-Remote-Group --requestheader-extra-headers-prefix=X-Remote-Extra- --etcd-servers=http://localhost:2379 --secure-port=9443 --tls-ca-file  /var/run/kubernetes/apiserver.crt --print-bearer-token
