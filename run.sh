#!/usr/bin/env bash

ROOT=${ROOT:-/out}
GOROOT=${GOROOT:-/go}
GOPATH=$ROOT:$GOROOT

export REPO="$(find . -name vendor -type d | head -n 1)"
export REPO="${REPO%/vendor}"
export REPO="${REPO#./}"

generate=""
for i in $(grep "// +genapi=true" -R .); do
    generate=$i+","
done

echo "gen " $generate

go run /go/src/github.com/pwittrock/apiserver-helloworld/vendor/k8s.io/apiserver-builder/cmd/genwiring/main.go --input-dirs $REPO/apis/...
go run /go/src/github.com/pwittrock/apiserver-helloworld/vendor/k8s.io/kubernetes/cmd/libs/go2idl/conversion-gen/main.go -i $REPO/apis/...  --extra-peer-dirs="k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/apimachinery/pkg/conversion,k8s.io/apimachinery/pkg/runtime" -o $ROOT/$REPO/src  -O zz_generated.conversion --go-header-file /go/src/github.com/pwittrock/apiserver-helloworld/boilerplate.go.txt
go run /go/src/github.com/pwittrock/apiserver-helloworld/vendor/k8s.io/kubernetes/cmd/libs/go2idl/deepcopy-gen/main.go -i $REPO/apis/... -o $ROOT/$REPO/src/ -O zz_generated.deepcopy --go-header-file /go/src/github.com/pwittrock/apiserver-helloworld/boilerplate.go.txt
go run /go/src/github.com/pwittrock/apiserver-helloworld/vendor/k8s.io/kubernetes/cmd/libs/go2idl/openapi-gen/main.go  -i "$REPO/apis...,k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/apimachinery/pkg/api/resource/,k8s.io/apimachinery/pkg/version/,k8s.io/apimachinery/pkg/runtime/,k8s.io/apimachinery/pkg//util/intstr/" -o $ROOT/$REPO/src --output-package "$REPO/pkg/openapi" --go-header-file /go/src/github.com/pwittrock/apiserver-helloworld/boilerplate.go.txt
go run /go/src/github.com/pwittrock/apiserver-helloworld/vendor/k8s.io/kubernetes/cmd/libs/go2idl/openapi-gen/main.go  -i "$REPO/apis...,k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/apimachinery/pkg/api/resource/,k8s.io/apimachinery/pkg/version/,k8s.io/apimachinery/pkg/runtime/,k8s.io/apimachinery/pkg//util/intstr/" -o $ROOT/$REPO/src --output-package "$REPO/pkg/openapi" --go-header-file /go/src/github.com/pwittrock/apiserver-helloworld/boilerplate.go.txt
