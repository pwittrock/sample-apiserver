all: build

clean:
	bash -c "find ./apis/ -name zz_generated.api.*.go | xargs rm -f"
	rm -f main

generate: clean
	go generate main.go
	go run vendor/k8s.io/kubernetes/cmd/libs/go2idl/conversion-gen/main.go -i k8s.io/sample-apiserver/apis/... -o ~/sample-apiserver/src/  -O zz_generated.conversion -v 1
	go run vendor/k8s.io/kubernetes/cmd/libs/go2idl/deepcopy-gen/main.go -i k8s.io/sample-apiserver/apis/... -o ~/sample-apiserver/src/ -O zz_generated.deepcopy
	go run vendor/k8s.io/kubernetes/cmd/libs/go2idl/openapi-gen/main.go  -i "k8s.io/sample-apiserver/apis/...,k8s.io/kubernetes/pkg/api/v1,k8s.io/kubernetes/pkg/apis/meta/v1" --output-package "k8s.io/sample-apiserver/pkg/openapi"

build: generate
	go build main.go

run: build
	./main -v 10 --authentication-kubeconfig ~/config --authorization-kubeconfig ~/config --client-ca-file /var/run/kubernetes/apiserver.crt  --requestheader-client-ca-file /var/run/kubernetes/apiserver.crt --requestheader-username-headers=X-Remote-User --requestheader-group-headers=X-Remote-Group --requestheader-extra-headers-prefix=X-Remote-Extra- --etcd-servers=http://localhost:2379 --secure-port=9443 --tls-ca-file  /var/run/kubernetes/apiserver.crt
