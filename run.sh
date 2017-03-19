#!/usr/bin/env bash

set -o xtrace

ROOT=${ROOT:-/out}
GOPATHADD=":/go"
GOPATH=$ROOT$GOPATHADD
echo $GOPATH

GENGO=${GENGO:-/go/src/github.com/pwittrock/apiserver-helloworld/}
GO2IDL=$GENGO/vendor/k8s.io/kubernetes/cmd/libs/go2idl


export REPO="$(find $ROOT -name vendor -type d | head -n 1)"
export REPO="${REPO%/vendor}"
export REPO="${REPO#$ROOT/}"
export REPO="${REPO#src/}"
echo "Repo $REPO"

cd $ROOT/src/$REPO

apis="$REPO/apis/..."
input="$apis,k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/apimachinery/pkg/api/resource/,k8s.io/apimachinery/pkg/version/,k8s.io/apimachinery/pkg/runtime/,k8s.io/apimachinery/pkg//util/intstr/"
peers="k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/apimachinery/pkg/conversion,k8s.io/apimachinery/pkg/runtime"

echo "Generating wiring"
/usr/local/go/bin/go run $GENGO/vendor/k8s.io/apiserver-builder/cmd/genwiring/main.go --input-dirs "$REPO/apis/..."

echo "Generating conversions"
/usr/local/go/bin/go run $GO2IDL/conversion-gen/main.go -i "$REPO/apis/..."  --extra-peer-dirs=$peers -o $ROOT/src  -O zz_generated.conversion --go-header-file $GENGO/boilerplate.go.txt

echo "Generating deepcopy"
/usr/local/go/bin/go run $GO2IDL/deepcopy-gen/main.go -i "$REPO/apis/..." -o $ROOT/src/ -O zz_generated.deepcopy --go-header-file $GENGO/boilerplate.go.txt

echo "Generating openapi"
/usr/local/go/bin/go run $GO2IDL/openapi-gen/main.go  -i $input -o $ROOT/src --output-package $REPO/pkg/openapi --go-header-file $GENGO/boilerplate.go.txt
